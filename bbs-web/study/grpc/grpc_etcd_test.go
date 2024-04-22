//@Author: wulinlin
//@Description:
//@File:  grpc_etcd_testr
//@Version: 1.0.0
//@Date: 2024/04/21 16:14

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

type EtcdTestSuide struct {
	suite.Suite
	client *clientv3.Client // 使用etcd来做服务注册和发现
}

func (s *EtcdTestSuide) SetupSuite() {
	client, err := clientv3.NewFromURL("192.168.1.52:12379")
	require.NoError(s.T(), err)
	s.client = client

}

type Svc struct {
	hello.UnimplementedHelloServiceServer
}

func (s *Svc) SayHello(ctx context.Context, request *hello.HelloRequest) (*hello.HelloResponse, error) {
	return &hello.HelloResponse{
		Res: "你好哇，" + request.Name,
	}, nil
}

func (s Svc) mustEmbedUnimplementedHelloServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (s *EtcdTestSuide) TestClient() {
	etcdResolver, err := resolver.NewBuilder(s.client)
	if err != nil {
		panic(err)
	}
	// URL的规范 scheme:///xxx
	dial, err := grpc.Dial("etcd:///service/user",
		grpc.WithResolvers(etcdResolver),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	client := hello.NewHelloServiceClient(dial)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	sayHello, err := client.SayHello(ctx, &hello.HelloRequest{Name: "李银河"})
	if err != nil {
		panic(err)
	}
	s.T().Log("响应", sayHello)
}

func (s *EtcdTestSuide) TestServer() {

	l, err := net.Listen("tcp", "192.168.1.51:8090")
	require.NoError(s.T(), err)

	// 这个Context是控制创建租约的超时时间
	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second)
	defer cancel2()
	// 创建租约
	// ttl 租期 秒为单位  会在1/3的时候就会触发续约
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
	addr := "192.168.1.51:8090"

	// key 就是指的实例 key ：1. 如果有instance id 就用instance id 2.用本地IP+Port
	//key := fmt.Sprintf("")
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
	// 万一注册信息有变动的话 怎么办  比如注册端口改了？  开线程监听啦
	//go func() {
	//	ticker := time.NewTicker(time.Second)
	//	for now := range ticker.C {
	//		ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second)
	//		defer cancel1()
	//		// 覆盖的语义
	//		em.AddEndpoint(ctx1, key, endpoints.Endpoint{
	//			Addr:     addr,
	//			Metadata: now.String(), // 元数据信息 分组信息，权重信息，机房信息，负载信息等等
	//		}, clientv3.WithLease(leaseResp.ID)) // 更新的时候 也要带上租约信息
	//		// 注意 不能在一个UpdateWithOpts中，对一个key只能进行一次操作，不然会报错  存疑
	//		//em.Update(ctx, []*endpoints.UpdateWithOpts{
	//		//	{
	//		//		Update: endpoints.Update{
	//		//			Op:       endpoints.Delete,
	//		//			Key:      key,
	//		//			Endpoint: endpoints.Endpoint{Addr: addr},
	//		//		},
	//		//	},
	//		//})
	//		cancel1()
	//	}
	//}()

	server := grpc.NewServer()
	hello.RegisterHelloServiceServer(server, &Svc{})

	err = server.Serve(l)
	if err != nil {
		panic(err)
	}
	// 退出的时候 记得Delete
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	/**
	删除的步骤  1.取消续约  2. 将自己从注册中心删除 2.停止服务
	*/
	kaCancel()                        // 取消续约
	err = em.DeleteEndpoint(ctx, key) // 取消注册
	server.GracefulStop()             // 优雅退出
}

func TestEtcd(t *testing.T) {
	suite.Run(t, new(EtcdTestSuide))
}
