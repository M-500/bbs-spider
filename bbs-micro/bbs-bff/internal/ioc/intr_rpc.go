package ioc

import (
	intrv1 "bbs-micro/api/proto/gen/proto/intr/v1"
	"bbs-micro/bbs-bff/internal/service"
	"bbs-micro/bbs-bff/internal/web/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-11 16:31

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
