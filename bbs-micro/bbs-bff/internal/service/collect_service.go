package service

import (
	"bbs-micro/bbs-bff/internal/domain"
	"bbs-micro/bbs-bff/internal/repository"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-12 11:33

type ICollectService interface {
	GetByUid(ctx context.Context, uid, limit, offset int64) ([]domain.Collect, error)
	CreateCollect(ctx context.Context, uid int64, cname string, desc string, isPub bool) (int64, error)
	CollectEntity(ctx context.Context, biz string, uid, cid, bizId int64) (int64, error)
}

type collectService struct {
	repo repository.ICollectRepo
}

func NewCollectService(repo repository.ICollectRepo) ICollectService {
	return &collectService{repo: repo}
}

func (c *collectService) GetByUid(ctx context.Context, uid, limit, offset int64) ([]domain.Collect, error) {
	return c.repo.GetCollectListByID(ctx, uid, limit, offset)
}

func (c *collectService) CreateCollect(ctx context.Context, uid int64, cname string, desc string, isPub bool) (int64, error) {
	return c.repo.CreateCollect(ctx, uid, cname, desc, isPub)
}
func (c *collectService) CollectEntity(ctx context.Context, biz string, uid, cid, bizId int64) (int64, error) {
	return c.repo.CollectEntity(ctx, biz, uid, cid, bizId)
}
