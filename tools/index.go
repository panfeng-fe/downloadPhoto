package tools

import (
	"fmt"
	"os"
	"strings"

	"github.com/levigross/grequests"
	"github.com/panfeng-fe/pain-utils/perr"
)

func Get(href string) (*grequests.Response, error) {
	fmt.Println("正在访问：" + href)
	res, err := grequests.Get(href, nil)
	return res, err
}

func IgnoreErr[T any](res T, err error) T {
	if err != nil {
		fmt.Println(err)
	}
	return res
}

func IgnoreSecond[T any](res T, tmp bool) T {
	return res
}

func GetExePath(mode string) (dir string) {
	if mode == "procedure" {
		// 程序启动
		dir = strings.Replace(perr.PanicErrDouble(os.Executable()), "/main", "", 1)
	} else {
		// 源码启动
		dir = perr.PanicErrDouble(os.Getwd())
	}
	return
}
