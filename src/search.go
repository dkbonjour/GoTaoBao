/*
   Created by jinhan on 17-8-24.
   Tip:
   Update:
*/
package src

import (
	"encoding/json"
	"fmt"
	"github.com/hunterhug/GoSpider/spider"
	"github.com/hunterhug/GoSpider/util"
	"github.com/hunterhug/GoSpider/util/open"
	"github.com/hunterhug/GoSpider/util/xid"
	"path/filepath"
	"regexp"
	"strings"
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
		1: "综合排序(MayBe千人千面)",
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
	return url + "&filter_tianmao=tmall&tab=mall"
}

func Search(url string) ([]byte, error) {
	爬虫.SetUrl(url)
	return 爬虫.Get()
}

type Mods struct {
	ModData Items `json:"mods"`
	//PageName string `json:"pageName"`
}
type Items struct {
	Items ItemList `json:"itemlist"`
}
type ItemList struct {
	Data ItemData `json:"data"`
}

// core
type ItemData struct {
	Auctions []ItemObject `json:"auctions"`
}

// 我的商品(不区分广告，某些商品做了广告会被置顶！)
type ItemObject struct {
	IsTmallObject IsTmall `json:"shopcard"`      // 是否天猫
	CommentCount  string  `json:"comment_count"` // 评论数
	Nid           string  `json:"nid"`           // 商品ID
	//CommentUrl   string `json:"comment_url"`
	//DetailUrl    string `json:"detail_url"`
	ItemLoc  string `json:"item_loc"`  // 发货地
	Nick     string `json:"nick"`      //  店铺名字
	PicUrl   string `json:"pic_url"`   // 商品图片
	RawTitle string `json:"raw_title"` // 商品标题
	//ShopLink     string `json:"shopLink"`   // 店铺URL
	UserId    string `json:"user_id"`    // 卖家ID
	ViewFee   string `json:"view_fee"`   // 小费？
	ViewPrice string `json:"view_price"` // 价格
	ViewSales string `json:"view_sales"` // 付款人数

}

type IsTmall struct {
	Yes bool `json:"isTmall"`
}

func ParseSearchPrepare(data []byte) []byte {
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
	return []byte("")
}

// 解析到结构体
func ParseSearch(data []byte) Mods {
	items := Mods{}
	err := json.Unmarshal(data, &items)
	if err != nil {
		fmt.Println(err.Error())
	}
	return items
}

func SearchMain() {
	for {
		csv := []ItemObject{}

		fmt.Println(`
	-------------------------------
	欢迎使用强大的搜索框小工具
	你只需安装提示进行即可！
	联系QQ：459527502
	----------------------------------
	`)
		keyword := util.Input("请输入关键字(请使用+代替空格！):", "")
		keyword = strings.Replace(keyword, " ", "+", -1)
		types := 请问搜索如何排序()
		tmall := false
		if strings.Contains(strings.ToLower(util.Input("是否只搜索天猫商品(Y/y),默认N", "n")), "y") {
			tmall = true
		}

		pagestemp := util.Input("你要抓几页(1-100):", "1")
		pages, err := util.SI(pagestemp)
		if err != nil {
			fmt.Println("输入页数有问题")
			break
		}
		if pages > 100 || pages < 1 {
			fmt.Printf("你选择的页数有问题：%d\n", pages)
			break
		}
		url := ""
		for page := 1; page <= pages; page++ {
			if tmall {
				url = SearchPrepareTmall(keyword, page, types)
			} else {
				url = SearchPrepare(keyword, page, types)
			}
			fmt.Println("搜索:" + url)
			data, err := Search(url)
			if err != nil {
				fmt.Printf("抓取第%d页 失败：%s\n", page, err.Error())
			} else {
				fmt.Printf("抓取第%d页\n", page)
				/*	filename := filepath.Join(".", "原始数据", util.ValidFileName(keyword), "search"+util.IS(page)+".html")
					util.MakeDirByFile(filename)
					e := util.SaveToFile(filename, data)
					if e != nil {
						fmt.Printf("保存数据在:%s 失败:%s\n", filename, e.Error())
						continue
					}
					fmt.Printf("保存数据在:%s 成功\n", filename)*/
				xx := ParseSearchPrepare(data)
				if string(xx) == "" {
					fmt.Println("这页数据为空...")
					continue
				}
				a := ParseSearch(xx)
				if len(a.ModData.Items.Data.Auctions) > 0 {
					for _, v := range a.ModData.Items.Data.Auctions {
						csv = append(csv, v)
						//fmt.Printf("%#v\n", v)
					}
				}
			}
		}

		if len(csv) == 0 {
			fmt.Println("啥都没抓到")
			continue
		}
		/**************************/
		id := xid.New().String()
		fileonly := util.TodayString(5) + "*" + id
		rootdir := filepath.Join(".", "搜索结果", util.ValidFileName(keyword))
		util.MakeDir(rootdir)
		tempdata := "排序,商品标题,店铺名,发货地址,评论数,是否天猫,小费,价格,销量,用户ID,店铺URL,商品ID,商品详情URL,商品评论URL图片地址\n"

		for k, v := range csv {
			tempdata = tempdata + fmt.Sprintf("%v,%s,%s,%s,", k+1, CD(v.RawTitle), v.Nick, v.ItemLoc)
			tempdata = tempdata + fmt.Sprintf("%s,%v,%s,%s,%s,", v.CommentCount, v.IsTmallObject.Yes, v.ViewFee, v.ViewPrice, v.ViewSales)
			s1 := "http://store.taobao.com/shop/view_shop.htm?user_number_id=" + v.UserId
			s2 := "http://detail.tmall.com/item.htm?id=" + v.Nid
			s3 := s2 + "&on_comment=1"
			tempdata = tempdata + fmt.Sprintf("%s,%s,%s,%s,%s,%s\n", v.UserId, s1, v.Nid, s2, s3, "http:"+v.PicUrl)
		}

		filekeep := rootdir + "/" + fileonly + ".csv"
		util.SaveToFile(filekeep, []byte(tempdata))
		fmt.Println("保存成功，请打开:" + filekeep)
		if strings.Contains(strings.ToLower(util.Input("是否打开文件(Y/y)", "n")), "y") {
			open.Start(filekeep)
		}
		/*************************/
		if cancle() == "y" {
			break
		}
	}

}

func CD(a string) string {
	return TripAll(strings.Replace(a, ",", "*", -1))
}
