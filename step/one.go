package step

import (
	"os"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/panfeng-fe/pain-utils/perr"

	"downloadPhoto/tools"
)

var oneWg sync.WaitGroup

func getIndexPage(href string, hrefChan chan string) func() {
	return func() {
		hrefList := ""
		tools.IgnoreErr[*goquery.Document](goquery.NewDocumentFromReader(tools.IgnoreErr(tools.Get(href)))).Find("#content article header h2 a").Each(func(i int, s *goquery.Selection) {
			hrefList += strings.ReplaceAll(s.Text(), "/", "&") + "：" + tools.IgnoreSecond(s.Attr("href")) + "\n"
		})
		hrefChan <- hrefList
		oneWg.Done()
	}
}

func oneSave(hrefChan chan string, curPath string) {
	file := tools.IgnoreErr(os.OpenFile(curPath+"/study/pages.txt", os.O_APPEND|os.O_WRONLY, 0644))
	defer file.Close()
	for {
		select {
		case hrefList := <-hrefChan:
			file.WriteString(hrefList)
		}
	}
}

func One(homePage string, curPath string) {
	g := tools.Workers(15)
	hrefChan := make(chan string, 100)

	go oneSave(hrefChan, curPath)
	perr.PanicErrDouble[*goquery.Document](goquery.NewDocumentFromReader(perr.PanicErrDouble(tools.Get(homePage)))).Find("option ").Each(func(i int, s *goquery.Selection) {
		oneWg.Add(1)
		g.Run(getIndexPage(tools.IgnoreSecond(s.Attr("value")), hrefChan))
	})

	oneWg.Wait()
	close(hrefChan)
}
