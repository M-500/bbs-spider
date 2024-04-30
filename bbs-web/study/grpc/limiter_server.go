package grpc

import (
	"bbs-web/pkg/limiter"
	"bbs-web/study/grpc/hello"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// @Description
// @Author 代码小学生王木木

type LimiterServer struct {
	limit                    limiter.Limiter
	hello.HelloServiceServer // 组合自己的Server
}

// SayHello 选择实现自己想要限流的方法 使用装饰器重写
func (l *LimiterServer) SayHello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	key := fmt.Sprintf("limiter:service:helloService:SayHello:%s", req.Name)
	ok, err := l.limit.Limit(ctx, key)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, status.Errorf(codes.ResourceExhausted, "触发限流")
	}
	return l.HelloServiceServer.SayHello(ctx, req)
}
