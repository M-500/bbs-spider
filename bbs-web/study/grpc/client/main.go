package main

import (
	_ "bbs-web/pkg/grpcx/balancer/wrr"
	"bbs-web/study/grpc/hello"
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

// @Description
// @Author 代码小学生王木木

func main() {
	etcdClient, err := clientv3.NewFromURL("192.168.1.52:12379")
	if err != nil {
		panic(err)
	}
	etcdResolver, err := resolver.NewBuilder(etcdClient)
	if err != nil {
		panic(err)
	}
	// URL的规范 scheme:///xxx
	dial, err := grpc.Dial("etcd:///service/user/t1",
		grpc.WithResolvers(etcdResolver),
		grpc.WithDefaultServiceConfig(`{
 			"loadBalancingConfig": [ 
				{ "wrr_round_robin": {} } 
			]
		}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	client := hello.NewHelloServiceClient(dial)
	for i := 0; i < 10; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		sayHello, err := client.SayHello(ctx, &hello.HelloRequest{Name: "李银河"})
		cancel()
		if err != nil {
			panic(err)
		}
		fmt.Println("响应", sayHello.Res)
	}
}
