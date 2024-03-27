package main

import (
	"bbs-web/internal/ioc"
	"flag"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

/*
对标网站 https://mlog.club/
github https://github.com/mlogclub/bbs-go
*/

var configFile = flag.String("config", "etc/dev.yaml", "配置文件路径")

func main() {
	engine := InitWebServer(*configFile)
	initPrometheus()
	ioc.SetUpOTEL(ioc.AppConfig)
	engine.Run(":8181")
}
func initPrometheus() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		// 监听端口，可以做成可配置的
		http.ListenAndServe(":8899", nil)
	}()
}
