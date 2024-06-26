package svc

import (
	"enterprise360/apps/user/model"
	"enterprise360/apps/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	DB model.EbUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		DB:     model.NewEbUserModel(conn, c.CacheRedis),
	}
}
