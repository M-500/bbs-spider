package logic

import (
	"context"

	"enterprise360/service/search/rpc/internal/svc"
	"enterprise360/service/search/rpc/types/search"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncCompanyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSyncCompanyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncCompanyLogic {
	return &SyncCompanyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 同步企业信息的接口
func (l *SyncCompanyLogic) SyncCompany(in *search.SyncCompanyRequest) (*search.SyncCompanyResponse, error) {
	// todo: add your logic here and delete this line

	return &search.SyncCompanyResponse{}, nil
}
