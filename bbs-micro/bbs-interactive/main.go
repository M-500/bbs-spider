package main

import (
	"flag"
)

var configFile = flag.String("config", "bbs-interactive/etc/dev.yaml", "配置文件路径")

func main() {
	app := InitApp(*configFile)
	for _, c := range app.consumer {
		err := c.Start()
		if err != nil {
			panic(err)
		}
	}
	app.server.Serve()
}
