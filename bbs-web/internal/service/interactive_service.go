package service

import (
	"bbs-web/internal/repository"
	"bbs-web/internal/repository/cache"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-07 19:13

type InteractiveService interface {
	IncrReadCnt(ctx context.Context, biz string, id int64) error
	Like(ctx context.Context, biz string, id int64, id2 int64) error
	CancelLike(ctx context.Context, biz string, id int64, id2 int64) error

	CollectArt(ctx context.Context, biz string, bizId int64, uId int64, cId int64) error
}

type interactiveService struct {
	repo  repository.InteractiveRepo
	cache cache.RedisInteractiveCache
}

func NewInteractiveService(repo repository.InteractiveRepo, cache cache.RedisInteractiveCache) InteractiveService {
	return &interactiveService{
		repo:  repo,
		cache: cache,
	}
}

func (i *interactiveService) IncrReadCnt(ctx context.Context, biz string, id int64) error {
	// 操作DB和操作缓存的顺序能换吗？？
	err := i.repo.IncrReadCnt(ctx, biz, id)
	if err != nil {
		return err
	}
	// 操作缓存  也可以用异步操作
	return i.cache.IncrReadCntIfPresent(ctx, biz, id)
}

func (i *interactiveService) Like(ctx context.Context, biz string, id int64, uid int64) error {

	return i.repo.IncrLike(ctx, biz, id, uid)
}

func (i *interactiveService) CancelLike(ctx context.Context, biz string, id int64, uid int64) error {
	return i.repo.DecrLike(ctx, biz, id, uid)
}

func (i *interactiveService) CollectArt(ctx context.Context, biz string, bizId int64, uId int64, cId int64) error {
	return nil
}
