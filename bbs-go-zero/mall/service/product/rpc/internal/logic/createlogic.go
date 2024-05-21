package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"mall/service/product/model"
	"time"

	"mall/service/product/rpc/internal/svc"
	"mall/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *product.CreateRequest) (*product.CreateResponse, error) {
	nw := time.Now()
	newProd := model.Product{
		Name:       in.Name,
		Desc:       in.Desc,
		Stock:      uint64(in.Stock),
		Amount:     uint64(in.Amount),
		Status:     uint64(in.Status),
		CreateTime: nw,
		UpdateTime: nw,
	}

	res, err := l.svcCtx.ProductModel.Insert(l.ctx, &newProd)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, "系统错误")
	}
	return &product.CreateResponse{
		Id: id,
	}, nil
}
