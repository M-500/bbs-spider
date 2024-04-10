package service

import (
	"bbs-micro/bbs-interactive/domain"
	"context"
)

//go:generate mockgen -source=./interactive_service.go -package=svcmocks -destination=./svcmocks/interactive.mock.go
type InteractiveService interface {
	IncrReadCnt(ctx context.Context, biz string, bizId int64) error
	Like(ctx context.Context, biz string, bizId int64, uid int64) error
	CancelLike(ctx context.Context, biz string, bizId int64, id2 int64) error
	CollectArt(ctx context.Context, biz string, bizId int64, uId int64, cId int64) error
	Get(ctx context.Context, biz string, bizId int64, uid int64) (domain.Interactive, error)
	GetByIds(ctx context.Context, biz string, ids []int64) (map[int64]domain.Interactive, error)
}

type interactiveService struct {
}

func NewInteractiveService() InteractiveService {
	return &interactiveService{}
}

func (svc *interactiveService) IncrReadCnt(ctx context.Context, biz string, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (svc *interactiveService) Like(ctx context.Context, biz string, id int64, id2 int64) error {
	//TODO implement me
	panic("implement me")
}

func (svc *interactiveService) CancelLike(ctx context.Context, biz string, id int64, id2 int64) error {
	//TODO implement me
	panic("implement me")
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
