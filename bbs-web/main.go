package main

import (
	"bbs-web/internal/ioc"
	"flag"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

/*
对标网站 https://mlog.club/
github https://github.com/mlogclub/bbs-go
*/

var configFile = flag.String("config", "etc/dev.yaml", "配置文件路径")

func main() {
	app := InitWebServer(*configFile)
	engine := app.server
	initPrometheus()
	ioc.SetUpOTEL(ioc.AppConfig)
	// 启动定时任务
	app.cron.Start()
	// 启动kafka消费者，
	for _, consumer := range app.consumers {
		err := consumer.Start()
		if err != nil {
			panic(err)
		}
	}
	engine.Run(":8181")

	// 退出
	ctx := app.cron.Stop()
	tm := time.NewTimer(time.Minute * 10)
	select {
	case <-tm.C:
	case <-ctx.Done(): // 可以考虑超时强制退出 防止有些任务执行特别长的时间
	}

}
func initPrometheus() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		// 监听端口，可以做成可配置的
		http.ListenAndServe(":8891", nil)
	}()
}
