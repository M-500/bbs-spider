package writesplit

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

// @Description 读写分离
// @Author 代码小学生王木木

type WriteSplit struct {
	master gorm.ConnPool
	slaves []gorm.ConnPool
}

func (w *WriteSplit) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	// 默认永远返回master的prepareContext  当然也可以默认返回slaves的数据 无所屌谓
	return w.master.PrepareContext(ctx, query)
}

func (w *WriteSplit) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	// 写操作直接操作Master 并无二意
	return w.master.ExecContext(ctx, query, args)
}

func (w *WriteSplit) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	// 读操作要考虑负载均衡！ 如果slave不止一个 那么必须负载均衡的去读取其他slaves
	if len(w.slaves) == 0 {
		return w.master.QueryContext(ctx, query, args...)
	}
	if len(w.slaves) == 1 {
		return w.slaves[0].QueryContext(ctx, query, args...)
	}
	// 否则就要用负载均衡的方式去读取slave啦
	//TODO implement me
	panic("implement me")
}

func (w *WriteSplit) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	if len(w.slaves) == 0 {
		return w.master.QueryRowContext(ctx, query, args...)
	}
	if len(w.slaves) == 1 {
		return w.slaves[0].QueryRowContext(ctx, query, args...)
	}
	// 否则就要用负载均衡的方式去读取slave啦
	//TODO implement me
	panic("implement me")
}
