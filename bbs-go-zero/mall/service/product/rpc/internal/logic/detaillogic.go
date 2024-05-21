package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"mall/service/product/model"

	"mall/service/product/rpc/internal/svc"
	"mall/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *product.DetailRequest) (*product.DetailResponse, error) {
	res, err := l.svcCtx.ProductModel.FindOne(l.ctx, uint64(in.Id))
	if err == model.ErrNotFound {
		return nil, status.Error(500, "产品不存在")
	}
	if err != nil {
		return nil, err
	}

	return &product.DetailResponse{
		Id:     int64(res.Id),
		Name:   res.Name,
		Desc:   res.Desc,
		Stock:  int64(res.Stock),
		Amount: int64(res.Amount),
		Status: int64(res.Status),
	}, nil
}
