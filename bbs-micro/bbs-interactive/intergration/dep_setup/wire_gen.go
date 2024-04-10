// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package dep_setup

import (
	"bbs-micro/bbs-interactive/grpc"
	"bbs-micro/bbs-interactive/repository"
	"bbs-micro/bbs-interactive/repository/cache"
	"bbs-micro/bbs-interactive/repository/dao"
	"bbs-micro/bbs-interactive/service"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitInteractiveGRPCServer() *grpc.InteractiveServiceServer {
	logger := InitLog()
	gormDB := InitTestDB()
	interactiveDao := dao.NewInteractiveDao(gormDB)
	cmdable := InitRedis()
	redisInteractiveCache := cache.NewRedisInteractiveCache(cmdable)
	interactiveRepo := repository.NewInteractiveRepo(logger, interactiveDao, redisInteractiveCache)
	interactiveService := service.NewInteractiveService(interactiveRepo, logger)
	interactiveServiceServer := grpc.NewInteractiveServiceServer(interactiveService)
	return interactiveServiceServer
}

// wire.go:

var thirdProvider = wire.NewSet(InitTestDB, InitRedis, InitLog)

var interactiveSvcProvider = wire.NewSet(dao.NewInteractiveDao, cache.NewRedisInteractiveCache, repository.NewInteractiveRepo, service.NewInteractiveService)