package main

import (
	. "downloadPhoto/step"
	"downloadPhoto/tools"
	"fmt"
	"os"

	"github.com/panfeng-fe/pain-utils/file"
)

func main() {
	var (
		url     string
		mode    string
		tagUrl  string
		tagMode string
	)
	fmt.Print("是否使用默认地址（常规使用默认地址），请输入 yes or no \n")
	fmt.Scanf("%s", &tagUrl)
	if tagUrl == "yes" {
		url = "https://b.taotu.in"
	} else {
		fmt.Print("请输入自定义地址 \n")
		fmt.Scanf("%s", &url)
	}
	fmt.Print("是否程序启动（常规使用程序启动），请输入 yes or no \n")
	fmt.Scanf("%s", &tagMode)
	if tagMode == "yes" {
		mode = "procedure"
	} else {
		mode = "code"
	}
	curPath := tools.GetExePath(mode)
	fmt.Println(curPath, url)
	cinit(mode, curPath)
	One(url, curPath)
	Two(curPath)
	Three(curPath)

}

func cinit(mode string, curPath string) {
	if !file.PathExists(curPath + "/study/photo") {
		os.MkdirAll(curPath+"/study/photo", 0777)
	}
	if !file.PathExists(curPath + "/study/pages.txt") {
		os.Create(curPath + "/study/pages.txt")
	}
	if !file.PathExists(curPath + "/study/record.txt") {
		os.Create(curPath + "/study/record.txt")
	}
}
