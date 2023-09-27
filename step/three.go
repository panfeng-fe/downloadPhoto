package step

import (
	"bufio"
	"downloadPhoto/tools"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/levigross/grequests"
	"github.com/panfeng-fe/pain-utils/file"
	"github.com/panfeng-fe/pain-utils/perr"
)

var threeWg sync.WaitGroup

func Three(curPath string) {

	data := perr.PanicErrDouble(ioutil.ReadFile(curPath + "/study/record.txt"))
	photoInfos := strings.Split(string(data), "\n")
	g := tools.Workers(15)

	for _, v := range photoInfos {
		if len(v) != 0 {
			photoInfo := strings.Split(v, "：")
			threeWg.Add(1)
			path := curPath + "/study/photo/" + photoInfo[0]
			if !file.PathExists(path) {
				os.Mkdir(path, 0777)
				os.Chmod(path, 0777)
			}
			g.Run(getPhoto(path+"/"+photoInfo[2]+".jpg", photoInfo[1]))
		}
	}

	threeWg.Wait()
}

func getPhoto(name string, href string) func() {
	return func() {
		res := tools.IgnoreErr(tools.Get(href))
		go storagePhoto(res, name)
	}
}

func storagePhoto(res *grequests.Response, name string) {
	reader := bufio.NewReaderSize(res, 100*1024)

	file := tools.IgnoreErr(os.Create(name))
	defer file.Close()

	writer := bufio.NewWriter(file)
	io.Copy(writer, reader)

	threeWg.Done()
}
