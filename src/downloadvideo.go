package src

import (
	"fmt"
	"github.com/hunterhug/GoSpider/util"
	"github.com/hunterhug/GoSpider/util/xid"
	"path/filepath"
	"regexp"
	"strings"
)

// 详情页主图
func DownloadVideoMain() {
	for {
		fmt.Println(`
	-------------------------------
	欢迎使用强大的视频下载小工具
	你只需按照提示进行即可！
	联系QQ：459527502
	----------------------------------
	`)
		fmt.Println("请输入天猫淘宝链接")
		fmt.Println("如:https://item.taobao.com/item.htm?id=536609211643")
		url := util.Input("请输入：", "")
		downlodvideo(TripAll(url))
		if cancle() == "y" {
			break
		}
	}
}

func downlodvideo(urlmany string) {
	爬虫.Url = urlmany
	urlhost := strings.Split(urlmany, "//")
	if len(urlhost) != 2 {
		fmt.Println("网站错误：开头必须为http://或https://")
		return
	}
	content, err := 爬虫.Get()
	if err != nil {
		fmt.Println(err.Error())
		return
	} else {
		v := Parsevideo(content)
		if string(v) == "" {
			fmt.Println("没有找到视频")
			return
		} else {
			vv := "http:" + string(v)
			fmt.Println("正在下载视频:" + vv)
			爬虫.Url = vv
			vc, err := 爬虫.Get()
			if err != nil {
				fmt.Println("爬视频失败:" + err.Error())
			} else {
				id := xid.New().String()
				fileonly := util.TodayString(5) + "*" + id + ".swf"
				ferr := util.SaveToFile(filepath.Join(".", fileonly), vc)
				if ferr != nil {
					fmt.Println("保存视频失败：" + ferr.Error())
				} else {
					fmt.Printf("保存成功，大小%v字节,请打开:%s\n", len(vc), fileonly)
				}
			}
		}
	}
}

func Parsevideo(data []byte) []byte {
	parsereg := "(//.*[.]swf)"
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
