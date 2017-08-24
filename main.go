/*
   Created by jinhan on 17-8-24.
   Tip:
   Update:
*/
package main

import (
	"fmt"
	"github.com/hunterhug/GoSpider/util"
	"github.com/hunterhug/GoTaoBao/src"
	"os"
	"path/filepath"
)

func main() {
	keyword := util.Input("请输入关键字:", "")
	types := src.Ask()

	pagestemp := util.Input("你要抓几页(1-100):", "1")
	pages, err := util.SI(pagestemp)
	if err != nil {
		fmt.Println("输入页数有问题")
		os.Exit(1)
	}
	if pages > 100 || pages < 1 {
		fmt.Printf("你选择的页数有问题：%d\n", pages)
		pages = 1
	}
	for page := 1; page <= pages; page++ {
		url := src.SearchPrepare(keyword, page, types)
		data, err := src.Search(url)
		if err != nil {
			fmt.Printf("抓取第%d页 失败：%s\n", page, err.Error())
		} else {
			fmt.Printf("抓取第%d页\n", page)
			filename := filepath.Join(util.CurDir(), "data", "search"+util.IS(page)+".html")
			util.MakeDirByFile(filename)
			e := util.SaveToFile(filename, data)
			fmt.Printf("%#v", e)
		}
	}
}
