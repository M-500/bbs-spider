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

type FailoverTestSuite struct {
	suite.Suite
	client *clientv3.Client
}

func (s *FailoverTestSuite) SetupSuite() {
	client, err := clientv3.NewFromURL("192.168.1.52:12379")
	require.NoError(s.T(), err)
	s.client = client

}
func TestFailover(t *testing.T) {
	suite.Run(t, new(FailoverTestSuite))
}

func (f *FailoverTestSuite) TestServer() {
	go func() {
		server := AlwaysFailServer{
			Name: "有问题",
		}
		f.startServer("192.168.1.51:18090", &server)
	}()
	go func() {
		server := Svc{
			Name: "正常",
		}
		f.startServer("192.168.1.51:18091", &server)
	}()
	server1 := Svc{
		Name: "正常",
	}
	f.startServer("192.168.1.51:18092", &server1)
}

// 初始化一个client客户端
func (f *FailoverTestSuite) TestClient() {
	etcdResolver, err := resolver.NewBuilder(f.client)
	if err != nil {
		panic(err)
	}
	svcCfg := `{
		  "loadBalancingConfig": [ { "round_robin": {} } ],
		  "methodConfig": [
			{
			  "name": [{ "service": "proto.HelloService"}],
			  "retryPolicy": {
				"maxAttempts":4,  				
				"initialBackoff":"0.01s",  		
				"maxBackoff":"0.1s",  			
				"backoffMultiplier":2.0,  		
				"retryableStatusCodes":["UNAVAILABLE","DEADLINE_EXCEEDED"]  
			  }
			}
		  ]
		}`
	dial, err := grpc.Dial("etcd:///service/user",
		grpc.WithResolvers(etcdResolver),
		grpc.WithDefaultServiceConfig(svcCfg),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	client := hello.NewHelloServiceClient(dial)
	for i := 0; i < 10; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		sayHello, err := client.SayHello(ctx, &hello.HelloRequest{Name: "李银河"})
		cancel()
		if err != nil {
			panic(err)
		}
		f.T().Log("响应", sayHello)
	}
}

func (s *FailoverTestSuite) startServer(addr string, svc hello.HelloServiceServer) {

	l, err := net.Listen("tcp", addr)
	require.NoError(s.T(), err)

	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second)
	defer cancel2()
	var ttl int64 = 30
	leaseResp, err := s.client.Grant(ctx2, ttl)
	if err != nil {
		panic(err)
	}

	// endpoint 以服务为维度，一个服务一个endpoint
	em, err := endpoints.NewManager(s.client, "service/user")
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	key := "service/user/" + addr

	// 在这一步之前 完成所有初始化工作 ==> 缓存预加载，配置预加载。。。。
	err = em.AddEndpoint(ctx, key, endpoints.Endpoint{
		Addr: addr,
	}, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		panic(err)
	}
	// 开启goroutine 自动续约
	kaCtx, kaCancel := context.WithCancel(context.Background())
	go func() {
		ch, err2 := s.client.KeepAlive(kaCtx, leaseResp.ID) // 自动续约，他会以ttl的1/3为标准去续约
		//s.client.KeepAliveOnce()  // 如果你想要控制续约的间隔，需要自己去调用 KeepAliveOnce 这个方法
		if err2 != nil {
			panic(err2)
		}
		for kaResp := range ch {
			fmt.Println("续约信息:", kaResp.String(), time.Now()) // 通常就是打印一下信息
		}
	}()

	server := grpc.NewServer()
	hello.RegisterHelloServiceServer(server, svc)

	err = server.Serve(l)
	if err != nil {
		panic(err)
	}
	// 退出的时候 记得Delete
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	kaCancel()                        // 取消续约
	err = em.DeleteEndpoint(ctx, key) // 取消注册
	server.GracefulStop()             // 优雅退出
}
