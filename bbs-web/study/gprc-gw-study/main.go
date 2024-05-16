package main

import (
	"bbs-web/study/gprc-gw-study/handler"
	proto "bbs-web/study/gprc-gw-study/pb/hello"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

// @Description
// @Author 代码小学生王木木

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:10086")
	if err != nil {
		log.Fatalf("failed to listen : %v\n", err)
	}
	s := grpc.NewServer()
	proto.RegisterHelloServer(s, &handler.HelloServer{})

	// grpc-gateway 注册服务
	ctx := context.Background()
	mux := runtime.NewServeMux()
	err = proto.RegisterHelloHandlerServer(ctx, mux, &handler.HelloServer{})
	if err != nil {
		log.Fatalf("failed to RegisterHelloHandlerServer error = %v\n", err)
	}
	// 启动服务
	httpMux := http.NewServeMux()
	httpMux.Handle("/", mux)
	// 配置http服务
	serverHttp := &http.Server{
		Addr: "127.0.0.1:8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mux.ServeHTTP(w, r)
		}),
	}
	// 启动http服务
	err = serverHttp.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to ListenAndServe error = %v\n", err)
	}
	err = s.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
