package service

import (
	"bbs-web/internal/repository"
	"bbs-web/internal/repository/cache"
	"github.com/gin-gonic/gin"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-07 19:13

type InteractiveService interface {
	IncrReadCnt(ctx *gin.Context, biz string, id int64) error
	Like(ctx *gin.Context, biz string, id int64, id2 int64) error
	CancelLike(ctx *gin.Context, biz string, id int64, id2 int64) error
}

type interactiveService struct {
	repo  repository.InteractiveRepo
	cache cache.RedisInteractiveCache
}

func NewInteractiveService(repo repository.InteractiveRepo) InteractiveService {
	return &interactiveService{
		repo: repo,
	}
}

func (i *interactiveService) IncrReadCnt(ctx *gin.Context, biz string, id int64) error {
	// 操作DB和操作缓存的顺序能换吗？？
	err := i.repo.IncrReadCnt(ctx, biz, id)
	if err != nil {
		return err
	}
	// 操作缓存  也可以用异步操作
	return i.cache.IncrReadCntIfPresent(ctx, biz, id)
}

func (i *interactiveService) Like(ctx *gin.Context, biz string, id int64, id2 int64) error {
	return nil
}
func (i *interactiveService) CancelLike(ctx *gin.Context, biz string, id int64, id2 int64) error {
	return nil
}
