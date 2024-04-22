//@Author: wulinlin
//@Description:
//@File:  grpc_etcd_testr
//@Version: 1.0.0
//@Date: 2024/04/21 16:14

package grpc

import (
	"bbs-web/study/grpc/hello"
	"context"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/grpc"
	"net"
	"testing"
	"time"
)

type EtcdTestSuide struct {
	suite.Suite
	client *clientv3.Client
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
	//TODO implement me
	panic("implement me")
}

func (s Svc) mustEmbedUnimplementedHelloServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (s *EtcdTestSuide) TestServer() {
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
	err = em.AddEndpoint(ctx, key, endpoints.Endpoint{
		Addr: addr,
	})
	if err != nil {
		panic(err)
	}

	l, err := net.Listen("tcp", "192.168.1.51:8090")
	require.NoError(s.T(), err)
	server := grpc.NewServer()
	hello.RegisterHelloServiceServer(server, &Svc{})

	err = server.Serve(l)
	if err != nil {
		panic(err)
	}

}

func TestEtcd(t *testing.T) {
	suite.Run(t, new(EtcdTestSuide))
}
