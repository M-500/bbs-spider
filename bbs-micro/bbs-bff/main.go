package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var configFile = flag.String("config", "bbs-bff/etc/dev.yaml", "配置文件路径")

func main() {
	app := InitWebServer(*configFile)
	engine := app.server
	initPrometheus()
	//ioc.SetUpOTEL(ioc.AppConfig)

	// 启动kafka消费者，
	for _, consumer := range app.consumers {
		err := consumer.Start()
		if err != nil {
			panic(err)
		}
	}
	engine.Run(":8181")
}

func initPrometheus() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		// 监听端口，可以做成可配置的
		http.ListenAndServe(":8899", nil)
	}()
}
