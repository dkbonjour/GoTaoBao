/*
   Created by jinhan on 17-8-25.
   Tip:
   Update:
*/
package src

import (
	"github.com/hunterhug/GoSpider/util"
	"strings"
)

func cancle() string {
	return strings.ToLower(util.Input("是否退出该功能: (Y/y),默认N", "n"))
}

func TripAll(a string)string{
	a=strings.Replace(a," ","",-1)
	a=strings.Replace(a,"\n","",-1)
	a=strings.Replace(a,"\r","",-1)
	a=strings.Replace(a,"\t","",-1)
	return a
}