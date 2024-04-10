package main

import (
	intrv1 "bbs-micro/api/proto/gen/proto/intr/v1"
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
)

var configFile = flag.String("config", "bbs-interactive/etc/dev.yaml", "配置文件路径")

func main() {
	server := grpc.NewServer()
	intrSvc := InitApp(*configFile)

	intrv1.RegisterInteractiveServiceServer(server, intrSvc)

	l, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err)
	}

	// 阻塞等待RPC调用
	err = server.Serve(l)
	log.Println(err)
}
