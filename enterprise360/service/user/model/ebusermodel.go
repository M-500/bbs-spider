package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ EbUserModel = (*customEbUserModel)(nil)

type (
	// EbUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEbUserModel.
	EbUserModel interface {
		ebUserModel
	}

	customEbUserModel struct {
		*defaultEbUserModel
	}
)

// NewEbUserModel returns a model for the database table.
func NewEbUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) EbUserModel {
	return &customEbUserModel{
		defaultEbUserModel: newEbUserModel(conn, c, opts...),
	}
}
