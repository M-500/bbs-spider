package dao

import (
	"context"
	"errors"
	"github.com/ecodeclub/ekit/syncx/atomicx"
)

// @Description
// @Author 代码小学生王木木

const (
	patternDstOnly  = "DST_ONLY"
	patternSrcOnly  = "SRC_ONLY"
	patternDstFirst = "DST_FIRST"
	patternSrcFirst = "SRC_FIRST"
)

type DoubleWriteInteractiveDAO struct {
	srcDao  InteractiveDao
	dstDao  interactiveDao
	pattern *atomicx.Value[string]
}

//func NewDoubleWriteInteractiveDAOV1(srcDB *gorm.DB, dstDB *gorm.DB) *DoubleWriteInteractiveDAO {
//	return &DoubleWriteInteractiveDAO{
//		srcDao:  NewInteractiveDao(srcDB),
//		dstDao:  NewInteractiveDao(dstDB),
//		pattern: atomicx.NewValueOf(patternSrcFirst),
//	}
//}

func NewDoubleWriteInteractiveDAO(srcDao InteractiveDao, dstDao interactiveDao) *DoubleWriteInteractiveDAO {
	return &DoubleWriteInteractiveDAO{srcDao: srcDao, dstDao: dstDao, pattern: atomicx.NewValueOf(patternSrcFirst)}
}

// 暴露一个方法 用于修改双写策略
func (d *DoubleWriteInteractiveDAO) UpdatePattern(path string) {
	d.pattern.Store(path)
}
func (d *DoubleWriteInteractiveDAO) IncrReadCnt(ctx context.Context, biz string, bizId int64) error {
	switch d.pattern.Load() {
	case patternSrcOnly:
		// 只操作src源表
		return d.srcDao.IncrReadCnt(ctx, biz, bizId)
	case patternSrcFirst:
		// 只读写源码阶段，但是会写目标表
		err := d.srcDao.IncrReadCnt(ctx, biz, bizId)
		if err != nil {
			// 源表都没有写成功，写个屁的目标表啊 出了问题只能等校验与修复程序
			return err
		}
		err = d.dstDao.IncrReadCnt(ctx, biz, bizId)
		if err != nil {
			// 这里要记录日志 因为写入目标表失败，不认为是一种失败，只需要记录日志就好了
			return nil
		}
		return nil
	case patternDstOnly:
		err := d.dstDao.IncrReadCnt(ctx, biz, bizId)
		if err != nil {
			return err
		}
		return nil
	case patternDstFirst:
		// 只读写源码阶段，但是会写目标表
		err := d.dstDao.IncrReadCnt(ctx, biz, bizId)
		if err != nil {
			// 源表都没有写成功，写个屁的目标表啊 出了问题只能等校验与修复程序
			return err
		}
		err = d.srcDao.IncrReadCnt(ctx, biz, bizId)
		if err != nil {
			// 这里要记录日志 因为写入源表失败，不认为是一种失败，只需要记录日志就好了
			return nil
		}
		return nil
	default:
		return errors.New("未知的双写模式")
	}
}

func (d *DoubleWriteInteractiveDAO) BatchIncrReadCnt(ctx context.Context, bizs []string, bizIds []int64) error {
	//TODO implement me
	panic("implement me")
}

func (d *DoubleWriteInteractiveDAO) IncrLikeInfo(ctx context.Context, biz string, id int64, uid int64) error {
	//TODO implement me
	panic("implement me")
}

func (d *DoubleWriteInteractiveDAO) DelLikeInfo(ctx context.Context, biz string, bizId int64, uid int64) error {
	//TODO implement me
	panic("implement me")
}

func (d *DoubleWriteInteractiveDAO) IncrLikeCnt(ctx context.Context, biz string, id int64, uid int64) error {
	//TODO implement me
	panic("implement me")
}

func (d *DoubleWriteInteractiveDAO) DecrLikeCnt(ctx context.Context, biz string, id int64, uid int64) error {
	//TODO implement me
	panic("implement me")
}

func (d *DoubleWriteInteractiveDAO) Get(ctx context.Context, biz string, id int64) (InteractiveModel, error) {
	switch d.pattern.Load() {
	case patternSrcOnly, patternSrcFirst:
		return d.srcDao.Get(ctx, biz, id)
	case patternDstOnly, patternDstFirst:
		return d.dstDao.Get(ctx, biz, id)
	default:
		return InteractiveModel{}, errors.New("未知的双写模式")
	}
}

func (d *DoubleWriteInteractiveDAO) GetLikeInfo(ctx context.Context, biz string, id int64, uid int64) (UserLikeBizModel, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DoubleWriteInteractiveDAO) GetCollectInfo(ctx context.Context, biz string, id int64, uid int64) (UserCollectBizModel, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DoubleWriteInteractiveDAO) QueryCollectList(ctx context.Context, uid, limit, offset int64) ([]CollectionModle, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DoubleWriteInteractiveDAO) InsertCollect(ctx context.Context, uid int64, cname string, desc string, isPub bool) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DoubleWriteInteractiveDAO) InsertCollectInfo(ctx context.Context, biz string, uid, cid, bizId int64) error {
	//TODO implement me
	panic("implement me")
}
