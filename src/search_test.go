/*
   Created by jinhan on 17-8-24.
   Tip:
   Update:
*/
package src

import (
	"fmt"
	"github.com/hunterhug/GoSpider/util"
	"path/filepath"
	"strings"
	"testing"
)

func TestSearchPrepare(t *testing.T) {
	keyword := "Mac 苹果"
	page := 1
	types := 1
	url := SearchPrepare(keyword, page, types)
	fmt.Println(url)
}

func TestSearchPrepareTmall(t *testing.T) {
	keyword := "Mac 苹果"
	page := 1
	types := 2
	url := SearchPrepareTmall(keyword, page, types)
	fmt.Println(url)
}

func TestSearch(t *testing.T) {
	keyword := "Mac 苹果"
	page := 1
	types := 1
	url := SearchPrepareTmall(keyword, page, types)
	fmt.Println(url)
	data, err := Search(url)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		filename := filepath.Join(util.CurDir(), "..", "原始数据", "search.html")
		util.MakeDirByFile(filename)
		e := util.SaveToFile(filename, data)
		fmt.Printf("%#v\n", e)
	}
}

func TestParseSeach(t *testing.T) {
	file := filepath.Join(util.CurDir(), "..", "原始数据", "mac")
	filejson := filepath.Join(util.CurDir(), "..", "原始数据", "macjson")
	util.MakeDir(filejson)
	files, err := util.ListDirOnlyName(file, "html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, filename := range files {
		data, err := util.ReadfromFile(file + "/" + filename)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			xx, _ := util.JsonBack(ParseSeach(data))
			fmt.Println(string(xx))
			nowjson := filejson + "/" + strings.Replace(filename, "html", "json", -1)
			fmt.Println(nowjson)
			util.SaveToFile(nowjson, xx)
		}
	}
}
