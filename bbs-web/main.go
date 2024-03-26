package main

import "flag"

/*
对标网站 https://mlog.club/
github https://github.com/mlogclub/bbs-go
*/

var configFile = flag.String("config", "etc/dev.yaml", "配置文件路径")

func main() {
	engine := InitWebServer(*configFile)

	engine.Run(":8181")
}
