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

// BeginTx
//
//	@Description: 支持事务的实现
func (d *DoubleWritePool) BeginTx(ctx context.Context, opts *sql.TxOptions) (gorm.ConnPool, error) {
	switch d.pattern.Load() {
	case patternSrcOnly:
		// 断言是否是txBeginner
		tx, err := d.src.(gorm.TxBeginner).BeginTx(ctx, opts)
		return &DoubleWritePool{
			src: tx,
		}, err
	case patternSrcFirst:
		txSrc, err := d.src.(gorm.TxBeginner).BeginTx(ctx, opts)
		if err != nil {
			return nil, err
		}
		txDst, err := d.dst.(gorm.TxBeginner).BeginTx(ctx, opts)
		if err != nil {
			return nil, nil
		}
		return &DoubleWritePool{
			src: txSrc,
			dst: txDst,
		}, nil
	case patternDstFirst:
		txDst, err := d.dst.(gorm.TxBeginner).BeginTx(ctx, opts)
		if err != nil {
			return nil, err
		}
		txSrc, err := d.src.(gorm.TxBeginner).BeginTx(ctx, opts)
		if err != nil {
			return nil, nil
		}
		return &DoubleWritePool{
			src: txSrc,
			dst: txDst,
		}, nil
	case patternDstOnly:
		tx, err := d.dst.(gorm.TxBeginner).BeginTx(ctx, opts)
		return &DoubleWritePool{
			dst: tx,
		}, err
	}
	return nil, errors.New("未知的双写模式")
}

// PrepareContext
//
//	@Description: Prepare的语句会进来这里  预编译语句
func (d *DoubleWritePool) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	//TODO implement me
	panic("不支持 PrepareContext")
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
		if d.dst == nil {
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
		if d.src == nil {
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

type DoubleWritePoolTx struct {
	src     *sql.Tx
	dst     *sql.Tx
	pattern string
}

func (d *DoubleWritePoolTx) Commit() error {
	switch d.pattern {
	case patternSrcOnly:
		return d.src.Commit()
	case patternSrcFirst:
		err := d.src.Commit()
		if err != nil {
			// 源库的事务失败了， 目标库上的事务要不要提交 ==> 不要
			return err
		}
		if d.dst != nil {
			err = d.dst.Commit()
			if err != nil {
				// 吞掉错误，记录日志
			}
		}
		return nil
	case patternDstFirst:
		err := d.dst.Commit()
		if err != nil {
			return err
		}
		if d.src != nil {
			err = d.src.Commit()
			if err != nil {
				// 吞掉错误，记录日志
			}
		}
		return nil
	case patternDstOnly:
		return d.dst.Commit()
	}
	return errors.New("未知的双写模式")
}

func (d *DoubleWritePoolTx) Rollback() error {
	//TODO implement me
	panic("implement me")
}

func (d *DoubleWritePoolTx) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	panic("implement me")
}

func (d *DoubleWritePoolTx) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	switch d.pattern {
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

func (d *DoubleWritePoolTx) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	switch d.pattern {
	case patternSrcOnly, patternSrcFirst:
		return d.src.QueryContext(ctx, query, args...)
	case patternDstFirst, patternDstOnly:
		return d.dst.QueryContext(ctx, query, args...)
	default:
		panic("未知的双写模式")
		return nil, errors.New("未知的双写模式")
	}
}

func (d *DoubleWritePoolTx) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	switch d.pattern {
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
