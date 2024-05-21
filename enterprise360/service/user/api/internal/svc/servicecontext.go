package svc

import (
	"enterprise360/service/user/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

// 全程带着Gonfig结构体到处跑
func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
