package ioc

import (
	grpc2 "bbs-micro/bbs-interactive/grpc"
	"bbs-micro/pkg/grpcx"
	"bbs-micro/pkg/logger"
	"google.golang.org/grpc"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-11 12:20

func InitGRPCXServer(cfg *Config, server *grpc2.InteractiveServiceServer, l logger.Logger) *grpcx.ServerX {
	svc := grpc.NewServer()
	server.Register(svc) // 注册server
	return &grpcx.ServerX{
		Port:     cfg.GRPCCfg.Port,
		EtcdAddr: cfg.GRPCCfg.EtcdAddr,
		Server:   svc,
		Name:     cfg.ServiceName,
		Lg:       l,
	}
}
