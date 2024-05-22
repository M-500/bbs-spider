package logic

import (
	"context"

	"enterprise360/service/search/rpc/internal/svc"
	"enterprise360/service/search/rpc/types/search"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCompanyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchCompanyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchCompanyLogic {
	return &SearchCompanyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchCompanyLogic) SearchCompany(in *search.SearchRequest) (*search.SearchResponse, error) {
	// todo: add your logic here and delete this line

	return &search.SearchResponse{}, nil
}
