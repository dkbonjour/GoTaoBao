/*
   Created by jinhan on 17-8-25.
   Tip:
   Update:
*/
package src

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/hunterhug/GoSpider/query"
	"github.com/hunterhug/GoSpider/util"
	"path/filepath"
	"regexp"
	"strings"
)

// 详情页主图
func DownloadPicMain() {
	for {
		fmt.Println(`
	-------------------------------
	欢迎使用强大的图片下载小工具
	你只需按照提示进行即可！
	联系QQ：459527502
	----------------------------------
	`)
		fmt.Println("请输入网址链接*保存目录")
		fmt.Println("如：https://item.taobao.com/item.htm?id=40066362090*taobao")
		fmt.Println("------------以上详情页会保存在“图片/taobao”文件夹下--------------")
		url := util.Input("请输入：", "")
		downlod(TripAll(url))
		if cancle() == "y" {
			break
		}
	}
}

func downlod(urlmany string) {
	temp := strings.Split(urlmany, "*")
	url := temp[0]
	//filename := util.TodayString(3)
	filename := "默认保存"
	if len(temp) >= 2 && temp[1] != "" {
		filename = util.ValidFileName(temp[1])
	}
	dir := filepath.Join(".", "图片", filename)
	util.MakeDir(dir)
	爬虫.Url = url
	urlhost := strings.Split(url, "//")
	if len(urlhost) != 2 {
		fmt.Println("网站错误：开头必须为http://或https://")
		return
	}
	content, err := 爬虫.Get()
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		docm, err := query.QueryBytes(content)
		if err != nil {
			fmt.Println(err.Error())
			return
		} else {
			docm.Find("img").Each(func(num int, node *goquery.Selection) {
				img, e := node.Attr("src")
				if e == false {
					img, e = node.Attr("data-src")
				}
				if e && img != "" {
					if strings.Contains(img, ".gif") {
						return
					}
					fmt.Println("原始文件：" + img)
					temp := img
					if strings.Contains(url, "taobao.com") || strings.Contains(url, "tmall.com") {
						r, _ := regexp.Compile(`([\d]{1,4}x[\d]{1,4})`)
						imgdudu := r.FindStringSubmatch(img)
						sizes := "720*720"
						if len(imgdudu) == 2 {
							sizes = imgdudu[1]
						}
						temp = strings.Replace(img, sizes, "720x720", -1)
					}
					filename := util.Md5(temp)
					if util.FileExist(dir + "/" + filename + ".jpg") {
						fmt.Println("文件存在：" + dir + "/" + filename)
					} else {
						fmt.Println("下载:" + temp)
						爬虫.Url = temp
						if strings.HasPrefix(temp, "//") {
							爬虫.Url = "http:" + temp
						}
						imgsrc, e := 爬虫.Get()
						if e != nil {
							fmt.Println("下载出错" + temp + ":" + e.Error())
							return
						}
						e = util.SaveToFile(dir+"/"+filename+".jpg", imgsrc)
						if e == nil {
							fmt.Println("成功保存在" + dir + "/" + filename)
						}
						//util.Sleep(1)
						//fmt.Println("暂停1秒")
					}
				}
			})

		}

	}
}
