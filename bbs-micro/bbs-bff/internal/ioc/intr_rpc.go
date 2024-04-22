package ioc

import (
	intrv1 "bbs-micro/api/proto/gen/proto/intr/v1"
	"bbs-micro/bbs-bff/internal/service"
	"bbs-micro/bbs-bff/internal/web/client"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitEtcd(cfg *Config) *clientv3.Client {
	var EtcdCfg clientv3.Config
	EtcdCfg.Endpoints = cfg.EtcdCfg.Addr
	client, err := clientv3.New(EtcdCfg)
	if err != nil {
		panic(err)
	}
	return client
}

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-11 16:31
func InitInterGRPCClientV1(client *clientv3.Client, cfg *Config) intrv1.InteractiveServiceClient {

	builder, err := resolver.NewBuilder(client)
	if err != nil {
		panic(err)
	}
	opts := []grpc.DialOption{
		grpc.WithResolvers(builder),
	}
	rpcClienCfg := cfg.GRPCCfg.Client
	if rpcClienCfg.Secure {
		// 判断是否要加载证书之类的
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	dial, err := grpc.Dial(cfg.GRPCCfg.Client.Addr, opts...)
	if err != nil {
		panic(err)
	}
	// 初始化RPC客户端
	return intrv1.NewInteractiveServiceClient(dial)
}

// InitInterGRPCClient
//
//	@Description: 流量控制的客户端
//	@param localSvc
//	@param cfg
//	@return intrv1.InteractiveServiceClient
func InitInterGRPCClient(localSvc service.InteractiveService, cfg *Config) intrv1.InteractiveServiceClient {

	var opts []grpc.DialOption
	rpcClienCfg := cfg.GRPCCfg.Client

	if rpcClienCfg.Secure {
		// 判断是否要加载证书之类的
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	dial, err := grpc.Dial(cfg.GRPCCfg.Client.Addr, opts...)
	if err != nil {
		panic(err)
	}
	// 初始化RPC客户端
	remote := intrv1.NewInteractiveServiceClient(dial)

	local := client.NewInteractiveServiceAdapter(localSvc)
	return client.NewGreyScaleInteractiveServiceClient(remote, local)
}
