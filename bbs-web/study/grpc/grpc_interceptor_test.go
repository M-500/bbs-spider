package grpc

import (
	"bbs-web/study/grpc/hello"
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"testing"
	"time"
)

// @Description
// @Author 代码小学生王木木

var interceptor1 grpc.UnaryServerInterceptor = func(ctx context.Context, req any, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp any, err error) {
	fmt.Println("第1个拦截器开始")
	a, err := handler(ctx, req)
	fmt.Println("第1个拦截器结束")
	return a, err
}

var interceptor2 grpc.UnaryServerInterceptor = func(ctx context.Context, req any, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp any, err error) {
	fmt.Println("第2个拦截器开始")
	a, err := handler(ctx, req)
	fmt.Println("第2个拦截器结束")
	return a, err
}

func TestServer(t *testing.T) {
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptor1, interceptor2)) // 传递多个拦截器
	defer func() {
		server.GracefulStop() // 优雅退出
	}()
	svc := &Svc1{}
	hello.RegisterHelloServiceServer(server, svc)
	// 创建监听器
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	err = server.Serve(l)
}

var clientFirst grpc.UnaryClientInterceptor = func(ctx context.Context, method string,
	req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	// 编写拦截逻辑
	err := invoker(ctx, method, req, reply, cc, opts...) // 指定调用

	return err
}

func TestClient(t *testing.T) {
	dial, err := grpc.Dial("localhost:8888",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(clientFirst),
	)
	require.NoError(t, err)
	client := hello.NewHelloServiceClient(dial)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	sayHello, err := client.SayHello(ctx, &hello.HelloRequest{Name: "测试"})
	require.NoError(t, err)
	t.Log(sayHello)
}

type Svc1 struct {
	hello.UnimplementedHelloServiceServer
	Name string
}

func (s *Svc1) SayHello(ctx context.Context, request *hello.HelloRequest) (*hello.HelloResponse, error) {
	return &hello.HelloResponse{
		Res: "你好哇，" + request.Name,
	}, nil
}

func (s Svc1) mustEmbedUnimplementedHelloServiceServer() {
	//TODO implement me
	panic("implement me")
}
