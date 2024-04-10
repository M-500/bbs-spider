package repository

import (
	"bbs-micro/bbs-interactive/domain"
	"bbs-micro/bbs-interactive/repository/cache"
	"bbs-micro/bbs-interactive/repository/dao"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-10 12:50

type InteractiveRepo interface {
	IncrReadCnt(ctx context.Context, biz string, bizId int64) error
	// BatchIncrReadCnt biz 和 bizId 长度必须一致
	BatchIncrReadCnt(ctx context.Context, biz []string, bizId []int64) error
	IncrLike(ctx context.Context, biz string, id int64, uid int64) error
	DecrLike(ctx context.Context, biz string, id int64, uid int64) error
	AddCollectionItem(ctx context.Context, biz string, id int64, cid int64, uid int64) error
	Get(ctx context.Context, biz string, id int64) (domain.Interactive, error)
	Liked(ctx context.Context, biz string, id int64, uid int64) (bool, error)
	Collected(ctx context.Context, biz string, id int64, uid int64) (bool, error)
}

type interactiveRepo struct {
	dao   dao.InteractiveDao
	cache cache.RedisInteractiveCache
}

func NewInteractiveRepo(d dao.InteractiveDao, c cache.RedisInteractiveCache) InteractiveRepo {
	return &interactiveRepo{
		dao:   d,
		cache: c,
	}
}

func (i *interactiveRepo) IncrReadCnt(ctx context.Context, biz string, bizId int64) error {
	err := i.dao.IncrReadCnt(ctx, biz, bizId)
	if err != nil {

	}
	return i.cache.IncrReadCntIfPresent(ctx, biz, bizId)
}

func (i *interactiveRepo) BatchIncrReadCnt(ctx context.Context, biz []string, bizId []int64) error {
	//TODO implement me
	panic("implement me")
}

func (i *interactiveRepo) IncrLike(ctx context.Context, biz string, id int64, uid int64) error {
	//TODO implement me
	panic("implement me")
}

func (i *interactiveRepo) DecrLike(ctx context.Context, biz string, id int64, uid int64) error {
	//TODO implement me
	panic("implement me")
}

func (i *interactiveRepo) AddCollectionItem(ctx context.Context, biz string, id int64, cid int64, uid int64) error {
	//TODO implement me
	panic("implement me")
}

func (i *interactiveRepo) Get(ctx context.Context, biz string, id int64) (domain.Interactive, error) {
	//TODO implement me
	panic("implement me")
}

func (i *interactiveRepo) Liked(ctx context.Context, biz string, id int64, uid int64) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (i *interactiveRepo) Collected(ctx context.Context, biz string, id int64, uid int64) (bool, error) {
	//TODO implement me
	panic("implement me")
}
