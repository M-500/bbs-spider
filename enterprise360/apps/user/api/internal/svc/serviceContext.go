package svc

import (
	"enterprise360/apps/user/api/internal/config"
	"enterprise360/apps/user/rpc/types/user"
	"enterprise360/apps/user/rpc/usercenterservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.UserCenterServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: usercenterservice.NewUserCenterService(zrpc.MustNewClient(c.UserRpc)),
	}
}
