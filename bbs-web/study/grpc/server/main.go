package main

import (
	"bbs-web/study/grpc/hello"
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/grpc"
	"net"
	"time"
)

// @Description
// @Author 代码小学生王木木

type Svc struct {
	hello.UnimplementedHelloServiceServer
	Name string
}

func (s *Svc) SayHello(ctx context.Context, request *hello.HelloRequest) (*hello.HelloResponse, error) {
	return &hello.HelloResponse{
		Res: "你好哇，" + request.Name + ", from " + s.Name,
	}, nil
}

func (s Svc) mustEmbedUnimplementedHelloServiceServer() {
	//TODO implement me
	panic("implement me")
}

func main() {
	// 初始化Etcd
	etcdClient, err := clientv3.NewFromURL("192.168.1.52:12379")
	if err != nil {
		panic(err)
	}
	go func() {
		startServer(etcdClient, "192.168.1.51:8090", 20)
	}()
	go func() {
		startServer(etcdClient, "192.168.1.51:8091", 50)
	}()
	startServer(etcdClient, "192.168.1.51:8092", 40)
}

func startServer(client *clientv3.Client, addr string, weight int) {
	var ttl int64 = 30
	var target = "service/user/t1"
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	// 1. 创建租约
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	leaseResp, err := client.Grant(timeout, ttl)
	if err != nil {
		panic(err)
	}

	// 2. 创建endpoint
	manager, err := endpoints.NewManager(client, target)
	if err != nil {
		panic(err)
	}
	withTimeout, c := context.WithTimeout(context.Background(), time.Second)
	defer c()

	key := target + "/" + addr
	err = manager.AddEndpoint(withTimeout, key, endpoints.Endpoint{
		Addr: addr,
		Metadata: map[string]any{
			"weight": weight,
		},
	}, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		panic(err)
	}
	// 3. 开启goroutine 自动续约
	ctx, cancelAlive2 := context.WithCancel(context.Background())
	go func() {
		_, err2 := client.KeepAlive(ctx, leaseResp.ID)
		if err2 != nil {
			panic(err2)
		}
		//for kaResp := range alive {
		//	fmt.Println("续约信息:", kaResp.String(), time.Now()) // 通常就是打印一下信息
		//}
	}()

	server := grpc.NewServer()
	hello.RegisterHelloServiceServer(server, &Svc{
		Name: addr, // 用地址来标识
	})
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
	// 退出的时候 记得Delete
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	cancelAlive2()
	err = manager.DeleteEndpoint(ctx, key)
	server.GracefulStop()
}
