package grpcx

import (
	"bbs-web/pkg/logger"
	"bbs-web/pkg/netx"
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"time"

	_ "go.etcd.io/etcd/client/v3/naming/endpoints"
	_ "go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"net"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-11 11:56

type ServerX struct {
	*grpc.Server
	Port     int
	Name     string
	EtcdAddr []string
	client   *clientv3.Client
	ttl      int64
	kaCancel func()
	lg       logger.Logger
	key      string
	em       endpoints.Manager
}

func (s *ServerX) Serve() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", s.Port))
	if err != nil {
		return err
	}
	return s.Server.Serve(lis)
}

func (s *ServerX) register() error {
	// 1. 初始化etcd client
	client, err := clientv3.New(clientv3.Config{
		Endpoints: s.EtcdAddr,
	})
	if err != nil {
		return err
	}
	s.client = client
	// 2. 创建manager
	// endpoint 以服务为维度，一个服务一个endpoint
	em, err := endpoints.NewManager(s.client, fmt.Sprintf("service/%s", s.Name))
	if err != nil {
		return err
	}
	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second)
	defer cancel2()
	leaseResp, err := s.client.Grant(ctx2, s.ttl)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	addr := fmt.Sprintf("%s:%d", netx.GetOutIP(), s.Port)
	s.key = fmt.Sprintf("service/%s/%s", s.Name, addr)
	err = em.AddEndpoint(ctx, s.key, endpoints.Endpoint{
		Addr: addr,
	}, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		panic(err)
	}
	kaCtx, kaCancel := context.WithCancel(context.Background())
	s.kaCancel = kaCancel
	ch, err2 := s.client.KeepAlive(kaCtx, leaseResp.ID) // 自动续约，他会以ttl的1/3为标准去续约
	if err2 != nil {
		return err2
	}

	go func() {
		for kaResp := range ch {
			s.lg.Debug("续约信息", logger.String("详情", kaResp.String()))
		}
	}()
	return nil
}

func (s *ServerX) Close() error {
	// 取消续约
	if s.kaCancel != nil {
		s.kaCancel()
	}
	// 注销注册信息
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err := s.em.DeleteEndpoint(ctx, s.key)
	if err != nil {
		return err
	}
	s.GracefulStop() // 优雅退出gRPC服务
	return nil
}
