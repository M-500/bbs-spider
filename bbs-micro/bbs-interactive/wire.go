//go:build wireinject

package main

import (
	"bbs-micro/bbs-interactive/grpc"
	"bbs-micro/bbs-interactive/ioc"
	"bbs-micro/bbs-interactive/repository"
	"bbs-micro/bbs-interactive/repository/cache"
	"bbs-micro/bbs-interactive/repository/dao"
	"bbs-micro/bbs-interactive/service"
	"github.com/google/wire"
)

var thirdPartySet = wire.NewSet(
	ioc.InitConfig,
	ioc.InitLogger,
	ioc.InitRedis,
	ioc.InitDatabase,
)

var interactiveSvcProvider = wire.NewSet(
	dao.NewInteractiveDao,
	cache.NewRedisInteractiveCache,
	repository.NewInteractiveRepo,
	service.NewInteractiveService,
)

func InitApp(path string) *grpc.InteractiveServiceServer {
	wire.Build(thirdPartySet, interactiveSvcProvider, grpc.NewInteractiveServiceServer)
	return new(grpc.InteractiveServiceServer)
}
