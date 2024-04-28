package grpc

import (
	"bbs-web/study/grpc/hello"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// @Description
// @Author 代码小学生王木木

type AlwaysFailServer struct {
	hello.UnimplementedHelloServiceServer
	Name string
}

func (s *AlwaysFailServer) SayHello(ctx context.Context, request *hello.HelloRequest) (*hello.HelloResponse, error) {
	fmt.Println("确实请求了我，但是一定会报错")
	return &hello.HelloResponse{
		Res: "这个调用 永远失败！" + request.Name,
	}, status.Errorf(codes.Unavailable, "模拟服务异常")
}

func (s *AlwaysFailServer) mustEmbedUnimplementedHelloServiceServer() {
	//TODO implement me
	panic("implement me")
}
