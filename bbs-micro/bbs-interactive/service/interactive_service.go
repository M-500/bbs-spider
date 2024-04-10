package service

import (
	"bbs-micro/bbs-interactive/domain"
	"bbs-micro/bbs-interactive/repository"
	"bbs-micro/pkg/logger"
	"context"
)

//go:generate mockgen -source=./interactive_service.go -package=svcmocks -destination=./svcmocks/interactive.mock.go
type InteractiveService interface {

	// IncrReadCnt
	//  @Description: 增加阅读计数
	//  @param ctx
	//  @param biz   业务标志
	//  @param bizId  业务ID
	//  @return error
	IncrReadCnt(ctx context.Context, biz string, bizId int64) error
	Like(ctx context.Context, biz string, bizId int64, uid int64) error
	CancelLike(ctx context.Context, biz string, bizId int64, uid int64) error
	CollectArt(ctx context.Context, biz string, bizId int64, uId int64, cId int64) error
	Get(ctx context.Context, biz string, bizId int64, uid int64) (domain.Interactive, error)
	GetByIds(ctx context.Context, biz string, ids []int64) (map[int64]domain.Interactive, error)
}

type interactiveService struct {
	repo repository.InteractiveRepo
	l    logger.Logger
}

func NewInteractiveService(r repository.InteractiveRepo, log logger.Logger) InteractiveService {
	return &interactiveService{
		repo: r,
		l:    log,
	}
}

func (svc *interactiveService) IncrReadCnt(ctx context.Context, biz string, id int64) error {
	return svc.repo.IncrReadCnt(ctx, biz, id)
}

func (svc *interactiveService) Like(ctx context.Context, biz string, id int64, uid int64) error {
	return svc.repo.IncrLike(ctx, biz, id, uid)
}

func (svc *interactiveService) CancelLike(ctx context.Context, biz string, id int64, id2 int64) error {
	return svc.repo.DecrLike(ctx, biz, id, id2)
}

func (svc *interactiveService) CollectArt(ctx context.Context, biz string, bizId int64, uId int64, cId int64) error {
	//TODO implement me
	panic("implement me")
}

func (svc *interactiveService) Get(ctx context.Context, biz string, id int64, uid int64) (domain.Interactive, error) {
	//TODO implement me
	panic("implement me")
}

func (svc *interactiveService) GetByIds(ctx context.Context, biz string, ids []int64) (map[int64]domain.Interactive, error) {
	//TODO implement me
	panic("implement me")
}
