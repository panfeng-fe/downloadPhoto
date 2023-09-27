package step

import (
	"downloadPhoto/tools"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/panfeng-fe/pain-utils/perr"
)

var twoWg sync.WaitGroup

func Two(curPath string) {
	data := perr.PanicErrDouble(ioutil.ReadFile(curPath + "/study/pages.txt"))
	photoInfos := strings.Split(string(data), "\n")
	g := tools.Workers(15)
	photoChan := make(chan string, 200)
	go twoSave(photoChan, curPath)
	for _, v := range photoInfos {
		if len(v) != 0 {
			photoInfo := strings.Split(v, "：")
			twoWg.Add(1)
			g.Run(getAllPhotoHref(photoInfo[0], photoInfo[1], photoChan))
		}
	}
	twoWg.Wait()
}

func getAllPhotoHref(title string, href string, hrefChan chan string) func() {
	return func() {
		allPhotoHref := ""
		tools.IgnoreErr[*goquery.Document](goquery.NewDocumentFromReader(tools.IgnoreErr(tools.Get(href)))).Find("#content #primary main p img").Each(func(i int, s *goquery.Selection) {
			allPhotoHref += title + "：" + tools.IgnoreSecond(s.Attr("src")) + "：" + fmt.Sprint(i+1) + "\n"
		})
		hrefChan <- allPhotoHref
		twoWg.Done()
	}
}

func twoSave(hrefChan chan string, curPath string) {
	file := tools.IgnoreErr(os.OpenFile(curPath+"/study/record.txt", os.O_APPEND|os.O_WRONLY, 0644))
	defer file.Close()
	for {
		select {
		case hrefList := <-hrefChan:
			file.WriteString(hrefList)
		}
	}
}
