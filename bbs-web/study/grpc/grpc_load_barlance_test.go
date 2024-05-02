package grpc

import (
	"bbs-web/study/grpc/hello"
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"testing"
	"time"
)

// @Description
// @Author 代码小学生王木木

type BalancerTestSuite struct {
	suite.Suite
	client *clientv3.Client // 使用etcd来做服务注册和发现
}

func (s *BalancerTestSuite) SetupSuite() {
	client, err := clientv3.NewFromURL("192.168.1.52:12379")
	require.NoError(s.T(), err)
	s.client = client

}
func (s *BalancerTestSuite) TestPickFirst() {
	var a uint32 = 4294967295 // uint32 的最大值为 4294967295
	fmt.Println(a)            // 输出 4294967295

	a++            // 对 uint32 类型的变量进行增量操作
	fmt.Println(a) // 输出 0，因为溢出后继续从 0 开始增量
}

func (s *BalancerTestSuite) TestOverFlowUint32() {
	var a uint32 = 4294967295 // uint32 的最大值为 4294967295
	fmt.Println(a)            // 输出 4294967295

	a++            // 对 uint32 类型的变量进行增量操作
	fmt.Println(a) // 输出 0，因为溢出后继续从 0 开始增量
}

func (s *BalancerTestSuite) TestClient() {
	etcdResolver, err := resolver.NewBuilder(s.client)
	if err != nil {
		panic(err)
	}
	// URL的规范 scheme:///xxx
	dial, err := grpc.Dial("etcd:///service/user/t1",
		grpc.WithResolvers(etcdResolver),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	client := hello.NewHelloServiceClient(dial)
	for i := 0; i < 10; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		sayHello, err := client.SayHello(ctx, &hello.HelloRequest{Name: "李银河"})
		cancel()
		if err != nil {
			panic(err)
		}
		s.T().Log("响应", sayHello)
	}

}

func (s *BalancerTestSuite) TestRobinClient() {
	etcdResolver, err := resolver.NewBuilder(s.client)
	if err != nil {
		panic(err)
	}
	// URL的规范 scheme:///xxx
	dial, err := grpc.Dial("etcd:///service/user",
		grpc.WithResolvers(etcdResolver),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	client := hello.NewHelloServiceClient(dial)
	for i := 0; i < 10; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		sayHello, err := client.SayHello(ctx, &hello.HelloRequest{Name: "李银河"})
		cancel()
		if err != nil {
			panic(err)
		}
		s.T().Log("响应", sayHello)
	}

}

func (s *BalancerTestSuite) TestClient() {
	etcdResolver, err := resolver.NewBuilder(s.client)
	if err != nil {
		panic(err)
	}
	// URL的规范 scheme:///xxx
	dial, err := grpc.Dial("etcd:///service/user",
		grpc.WithResolvers(etcdResolver),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	client := hello.NewHelloServiceClient(dial)
	for i := 0; i < 10; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		sayHello, err := client.SayHello(ctx, &hello.HelloRequest{Name: "李银河"})
		cancel()
		if err != nil {
			panic(err)
		}
		s.T().Log("响应", sayHello)
	}

}

// TestServer
//
//	@Description: 启动gRPC服务端
//	@receiver s
func (s *BalancerTestSuite) TestServer() {
	go func() {
		s.startServer("192.168.1.51:8090")
	}()
	s.startServer("192.168.1.51:8091")
}

func (s *BalancerTestSuite) startServer(addr string) {

	var ttl int64 = 30
	var target = "service/user/t1"
	listen, err := net.Listen("tcp", addr)
	require.NoError(s.T(), err)

	// 1. 创建租约
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	leaseResp, err := s.client.Grant(timeout, ttl)
	require.NoError(s.T(), err)

	// 2. 创建endpoint
	manager, err := endpoints.NewManager(s.client, target)
	if err != nil {
		panic(err)
	}
	withTimeout, c := context.WithTimeout(context.Background(), time.Second)
	defer c()

	key := target + "/" + addr
	err = manager.AddEndpoint(withTimeout, key, endpoints.Endpoint{
		Addr: addr,
	}, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		panic(err)
	}
	// 3. 开启goroutine 自动续约
	ctx, cancelAlive2 := context.WithCancel(context.Background())
	go func() {
		_, err2 := s.client.KeepAlive(ctx, leaseResp.ID)
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
func TestBalancerTestSuite(t *testing.T) {
	suite.Run(t, new(BalancerTestSuite))
}
