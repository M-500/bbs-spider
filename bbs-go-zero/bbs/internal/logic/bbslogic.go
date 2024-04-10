package logic

import (
	"context"

	"bbs-go-zero/bbs/internal/svc"
	"bbs-go-zero/bbs/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BbsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBbsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BbsLogic {
	return &BbsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BbsLogic) Bbs(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
