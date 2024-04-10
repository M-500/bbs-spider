package dep_setup

import (
	"bbs-micro/bbs-interactive/grpc"
	"bbs-micro/bbs-interactive/repository"
	"bbs-micro/bbs-interactive/repository/cache"
	"bbs-micro/bbs-interactive/repository/dao"
	"bbs-micro/bbs-interactive/service"
	"github.com/google/wire"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-10 15:51

var thirdProvider = wire.NewSet(InitTestDB)
var interactiveSvcProvider = wire.NewSet(
	service.NewInteractiveService,
	dao.NewInteractiveDao,
	cache.NewRedisInteractiveCache,
	repository.NewInteractiveRepo,
)

func InitInteractiveGRPCServer() *grpc.InteractiveServiceServer {
	wire.Build(thirdProvider, interactiveSvcProvider, grpc.NewInteractiveServiceServer)

	return new(grpc.InteractiveServiceServer)
}
