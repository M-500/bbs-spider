package logic

import (
	"context"

	"enterprise360/service/search/rpc/internal/svc"
	"enterprise360/service/search/rpc/types/search"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchLogic) Search(in *search.SearchRequest) (*search.SearchResponse, error) {
	// todo: add your logic here and delete this line

	return &search.SearchResponse{}, nil
}
