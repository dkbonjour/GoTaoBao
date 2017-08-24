/*
   Created by jinhan on 17-8-24.
   Tip:
   Update:
*/
package src

import (
	"fmt"
	"github.com/go_tool/util"
	"path/filepath"
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
		filename := filepath.Join(util.CurDir(), "..", "data", "search.html")
		util.MakeDirByFile(filename)
		e := util.SaveToFile(filename, data)
		fmt.Printf("%#v\n", e)
	}
}

func TestParseSeach(t *testing.T) {
	filename := filepath.Join(util.CurDir(), "..", "data", "search1.html")
	data, err := util.ReadfromFile(filename)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		xx := ParseSeach(data)
		fmt.Println(string(xx))
	}
}
