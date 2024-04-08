package repository

import (
	"bbs-web/internal/repository/dao"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-07 19:28

type InteractiveRepo interface {
	IncrReadCnt(ctx context.Context, biz string, bizId int64) error
	IncrLike(ctx context.Context, biz string, id int64, uid int64) error
	DecrLike(ctx context.Context, biz string, id int64, uid int64) error
}

type interactiveRepo struct {
	dao dao.InteractiveDao
}

func NewInteractiveRepo(dao dao.InteractiveDao) InteractiveRepo {
	return &interactiveRepo{
		dao: dao,
	}
}

func (repo *interactiveRepo) IncrReadCnt(ctx context.Context, biz string, bizId int64) error {
	// 要考虑缓存方案了

	return repo.dao.IncrReadCnt(ctx, biz, bizId)
}

func (repo *interactiveRepo) IncrLike(ctx context.Context, biz string, id int64, uid int64) error {
	// 插入数据库 更新点赞计数 更新缓存
	err := repo.dao.IncrLikeInfo(ctx, biz, id, uid)
	if err != nil {
		return err
	}
	// 同步缓存
	return nil
}
func (repo *interactiveRepo) DecrLike(ctx context.Context, biz string, id int64, uid int64) error {
	// 插入数据库
	return repo.dao.IncrLikeCnt(ctx, biz, id, uid)
	// 同步缓存
}
