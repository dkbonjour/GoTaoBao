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
)

func main() {
	fmt.Println(`
	---------------------------------------------
	|	亲爱的朋友，你好！
	|	欢迎使用皮卡秋秋制作的小工具
	|	友好超乎你想象！
	|	如果觉得好，给我一个star！
	|	https://github.com/hunterhug/GoTaoBao
	|	QQ：459527502
	---------------------------------------------
	`)

	for {
		fmt.Println(`
	-------温柔的提示框---------
	|天猫淘宝搜索框小工具: 请按 1 |
	|天猫淘宝啥图片小工具: 请按 2 |
	|天猫淘主图视频小工具: 请按 3 |
	|更多待续更多待续更多: 请按 x |
	--------------------------
		`)
		choice := util.Input("* 请你输入你要使用的功能:", "0")
		switch choice {
		case "1":
			src.SearchMain()
		case "2":
			src.DownloadPicMain()
		case "3":
			src.DownloadVideoMain()
		case "0":
			hello()
		default:
			hello()

		}
	}
}

func hello() {
	fmt.Println(`
	--
	- - -
	-
	--- -- - -------------
	---
	----------输入错误----- - - -- - -
	-----  -
	-
	-  --   - -- - --
	-  - -囧 -------
	   - - --    ---
	`)
}
