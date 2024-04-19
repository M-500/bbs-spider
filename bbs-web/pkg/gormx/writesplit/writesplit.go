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

func (w *WriteSplit) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	// 开启事务只能开在master上，没有什么好说得
	return w.master.(gorm.TxBeginner).BeginTx(ctx, opts)
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
	/**
	轮询，加权轮询，平滑的加权轮询，随机，加权随机，动态判定slaves的健康情况(永远挑最快响应的那个slave做处理)... 骚操作自己悟
	*/
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
