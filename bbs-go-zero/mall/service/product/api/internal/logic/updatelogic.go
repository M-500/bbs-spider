package logic

import (
	"context"
	"mall/service/product/rpc/types/product"

	"mall/service/product/api/internal/svc"
	"mall/service/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UpdateRequest) (resp *types.UpdateResponse, err error) {
	_, err = l.svcCtx.ProductRpc.Update(l.ctx, &product.UpdateRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdateResponse{}, nil
}
