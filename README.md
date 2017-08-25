# 淘宝天猫商品信息采集小工具

使用速度快的Golang语言，编译成单一exe二进制文件，方便快捷！

依赖[Project:Marmot(Tubo) - Golang Web Spider/Crawler/Scrapy Package | 爬虫库 ](https://github.com/hunterhug/GoSpider)

开发中～

## 项目情况

请安装Golang环境，然后直接`./build.sh`，直接点击二进制，如`exe`即可

运行命令行显示类似：

```
$ ./GoTaoBao_linux_amd64 

        ---------------------------------------------
        |       亲爱的朋友，你好！
        |       欢迎使用皮卡秋秋制作的小工具
        |       友好超乎你想象！
        |       如果觉得好，给我一个star！
        |       https://github.com/hunterhug/GoTaoBao
        |       QQ：459527502
        ---------------------------------------------
        

        -------温柔的提示框---------
        |天猫淘宝搜索框小工具: 请按 1 |
        |天猫淘宝啥图片小工具: 请按 2 |
        |更多待续更多待续更多: 请按 x |
        --------------------------
                
* 请你输入你要使用的功能:
```

### 目录

```
tree -L 2
├── build.sh
├── data
├── doc
│   ├── doc.png
│   └── img.png
├── GoTaoBao_linux_386
├── GoTaoBao_linux_amd64
├── GoTaoBao_windows_386.exe
├── GoTaoBao_windows_amd64.exe
├── main.go
├── README.md
├── src
│   ├── downloadpic.go
│   ├── search.go
│   ├── search_test.go
│   └── util.go
└── 图片
    ├── taobao
    └── 默认保存
```

## 淘宝天猫关键字框搜索小工具（开发中）

![doc.png](doc/doc.png)

## 淘宝天猫啥啥图片小工具（开发结束）

![doc.png](doc/img.png)

# Support

如果你觉得项目帮助到你,欢迎请我喝杯咖啡,或加QQ：459527502

微信
![微信](https://raw.githubusercontent.com/hunterhug/hunterhug.github.io/master/static/jpg/wei.png)

支付宝
![支付宝](https://raw.githubusercontent.com/hunterhug/hunterhug.github.io/master/static/jpg/ali.png)