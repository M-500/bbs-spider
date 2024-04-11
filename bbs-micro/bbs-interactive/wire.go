//go:build wireinject

package main

import (
	"bbs-micro/bbs-interactive/events"
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
	ioc.InitSaramaClient,
	ioc.InitConsumer,
	ioc.InitGRPCXServer,
	events.NewInteractiveReadEventBatchConsumer,
)

var interactiveSvcProvider = wire.NewSet(
	dao.NewInteractiveDao,
	cache.NewRedisInteractiveCache,
	repository.NewInteractiveRepo,
	service.NewInteractiveService,
)

func InitApp(path string) *App {
	wire.Build(thirdPartySet,
		interactiveSvcProvider,
		grpc.NewInteractiveServiceServer,
		wire.Struct(new(App), "*"))
	return new(App)
}
