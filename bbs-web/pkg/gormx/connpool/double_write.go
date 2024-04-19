package connpool

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ecodeclub/ekit/syncx/atomicx"
	"gorm.io/gorm"
)

// @Description 用装饰器模式
// @Author 代码小学生王木木

const (
	patternDstOnly  = "DST_ONLY"
	patternSrcOnly  = "SRC_ONLY"
	patternDstFirst = "DST_FIRST"
	patternSrcFirst = "SRC_FIRST"
)

type DoubleWritePool struct {
	src     gorm.ConnPool
	dst     gorm.ConnPool
	pattern *atomicx.Value[string]
}

// PrepareContext
//
//	@Description: Prepare的语句会进来这里  预编译语句
func (d *DoubleWritePool) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	//TODO implement me
	panic("implement me")
}

// ExecContext
//
//	@Description: Exec语句会进来这里
func (d *DoubleWritePool) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	switch d.pattern.Load() {
	case patternSrcOnly:
		// 只操作src源表
		return d.src.ExecContext(ctx, query, args...)
	case patternSrcFirst:
		// 只读写源码阶段，但是会写目标表
		res, err := d.src.ExecContext(ctx, query, args...)
		if err != nil {
			// 源表都没有写成功，写个屁的目标表啊 出了问题只能等校验与修复程序
			return res, err
		}
		res, err = d.dst.ExecContext(ctx, query, args...)
		if err != nil {
			// 这里要记录日志 因为写入目标表失败，不认为是一种失败，只需要记录日志就好了
			return res, nil
		}
		return res, nil
	case patternDstOnly:
		return d.dst.ExecContext(ctx, query, args...)
	case patternDstFirst:
		// 只读写源码阶段，但是会写目标表
		res, err := d.dst.ExecContext(ctx, query, args...)
		if err != nil {
			// 源表都没有写成功，写个屁的目标表啊 出了问题只能等校验与修复程序
			return res, err
		}
		res, err = d.src.ExecContext(ctx, query, args...)
		if err != nil {
			// 这里要记录日志 因为写入源表失败，不认为是一种失败，只需要记录日志就好了
			return res, nil
		}
		return res, nil
	default:
		panic("未知的双写模式")
		return nil, errors.New("未知的双写模式")
	}
}

// QueryContext
//
//	@Description: 查询语句会进来这里
func (d *DoubleWritePool) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	switch d.pattern.Load() {
	case patternSrcOnly, patternSrcFirst:
		return d.src.QueryContext(ctx, query, args...)
	case patternDstFirst, patternDstOnly:
		return d.dst.QueryContext(ctx, query, args...)
	default:
		panic("未知的双写模式")
		return nil, errors.New("未知的双写模式")
	}
}

// QueryRowContext
//
//	@Description: 查询一行的语句会进来这里
func (d *DoubleWritePool) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	switch d.pattern.Load() {
	case patternSrcOnly, patternSrcFirst:
		return d.src.QueryRowContext(ctx, query, args...)
	case patternDstFirst, patternDstOnly:
		return d.dst.QueryRowContext(ctx, query, args...)
	default:
		// 这里怎么返回一个error？ 没办法！只能panic
		panic("未知的双写模式")
		return &sql.Row{}
	}
}
