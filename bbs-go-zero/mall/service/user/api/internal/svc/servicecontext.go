package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"mall/service/user/api/internal/config"
	"mall/service/user/rpc/types/user"
	"mall/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
