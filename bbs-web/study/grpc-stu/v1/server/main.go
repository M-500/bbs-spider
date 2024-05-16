package main

import (
	"bbs-web/study/grpc-stu/v1/pb/hello"
	"context"
	"google.golang.org/grpc"
	"net"
)

// @Description
// @Author 代码小学生王木木

type HelloService struct {
	hello.UnimplementedHelloServiceServer
}

func (s *HelloService) SayHello(ctx context.Context, req *hello.HelloReq) (*hello.HelloResp, error) {
	return &hello.HelloResp{Reply: "说Hello"}, nil
}

func main() {
	// 创建TCP监听器
	lis, err := net.Listen("tcp", "18848")
	if err != nil {
		panic(err)
	}

	// 创建gPRC Server
	ser := grpc.NewServer()
	// 将自己实现的服务注册到 gRPCServer中
	hello.RegisterHelloServiceServer(ser, &HelloService{})
	// 启动RPC服务
	ser.Serve(lis)
}
