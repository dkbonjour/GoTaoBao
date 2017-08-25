/*
   Created by jinhan on 17-8-24.
   Tip:
   Update:
*/
package src

import (
	"fmt"
	"github.com/hunterhug/GoSpider/spider"
	"github.com/hunterhug/GoSpider/util"
	//"regexp"
	"os"
	"path/filepath"
	"regexp"
)

// 每页11列 44个商品 // 不用 ajax方式
type SearchQuery struct {
	KsTS        string `json:"-"`           // 1503560947454_856
	Ajax        bool   `json:"ajax"`        // 不要改  true
	Bcoffset    int    `json:"bcoffset"`    // 不要改 4
	Callback    string `json:"callback"`    // jsonp857
	DataKey     string `json:"data-key"`    // s
	Ntoffset    int    `json:"ntoffset"`    // 0
	P4ppushleft string `json:"p4ppushleft"` // 1,48

	// 重要
	Page      int    `json:"-"`          // 第5页
	DataValue string `json:"data-value"` // 2156 ---> （50-1）*44=2156
	KeyWord   string `json:"q"`          // 搜索关键字
}

var (
	爬虫   *spider.Spider
	搜索链接 = "https://s.taobao.com/search?q=%s&s=%d&sort=%s"
	搜索排序 = map[int]string{
		1: "综合排序",
		2: "人气从高到低",
		3: "销量从高到低",
		4: "信用从高到低",
		5: "价格 低-高",
		6: "价格 高-低",
		7: "总价 低-高",
		8: "总价 高-低",
	}
	OrderMap = map[int]string{
		1: "default",     // 综合排序
		2: "renqi-desc",  // 人气从高到低
		3: "sale-desc",   // 销量从高到低
		4: "credit-desc", // 信用从高到低
		5: "price-asc",   // 价格 低-高
		6: "price-desc",  // 价格 高-低
		7: "total-asc",   //总价 低-高
		8: "total-desc",  //总价 高-低
	}
)

func init() {
	爬虫, _ = spider.New(nil)
	爬虫.SetUa(spider.RandomUa())
	爬虫.SetHeaderParm("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	爬虫.SetHeaderParm("Accept-Encoding", "gzip, deflate, br")
	爬虫.SetHeaderParm("Accept-Language", "en-US,en;q=0.5")
}

func 请问搜索如何排序() int {
	fmt.Println("我想问你,想如何排序：")
	fmt.Println("----------------")
	for k := 1; k <= len(OrderMap); k++ {
		fmt.Printf("%-20s 请选择:%d\n", 搜索排序[k], k)
	}
	fmt.Println("----------------")
	choice := util.Input("请选择：", "1")
	fmt.Println("选择完毕:" + choice)
	if i, e := util.SI(choice); e != nil {
		fmt.Println("请认真选择！")
		return 请问搜索如何排序()
	} else {
		return i
	}
}

// 搜索全部类型商品
func SearchPrepare(keyword string, page int, order int) string {
	orderstring, ok := OrderMap[order]
	if !ok {
		orderstring = "default"
		fmt.Println("排序条件出错，采用默认")
	}
	url := fmt.Sprintf(搜索链接, util.UrlE(keyword), (page-1)*44, orderstring)
	return url
}

// 只搜索天猫
func SearchPrepareTmall(keyword string, page int, order int) string {
	url := SearchPrepare(keyword, page, order)
	return url + "&filter_tianmao=tmall"
}

func Search(url string) ([]byte, error) {
	爬虫.SetUrl(url)
	return 爬虫.Get()
}

func ParseSeach(data []byte) []byte {
	parsereg := "g_page_config = ({.*})"
	r, err := regexp.Compile(parsereg)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		bb := r.FindAllSubmatch(data, -1)
		if len(bb) > 0 {
			return bb[0][1]
		}
	}
	return []byte("empty")
}

func SearchMain() {
	fmt.Println(`
	-------------------------------
	欢迎使用强大的搜索框小工具
	你只需安装提示进行即可！
	联系QQ：459527502
	----------------------------------
	`)
	keyword := util.Input("请输入关键字:", "")
	types := 请问搜索如何排序()

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
		url := SearchPrepare(keyword, page, types)
		data, err := Search(url)
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
