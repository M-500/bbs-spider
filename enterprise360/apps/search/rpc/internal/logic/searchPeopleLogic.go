package logic

import (
	"context"

	"enterprise360/apps/search/rpc/internal/svc"
	"enterprise360/apps/search/rpc/types/search"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchPeopleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchPeopleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchPeopleLogic {
	return &SearchPeopleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchPeopleLogic) SearchPeople(in *search.SearchRequest) (*search.SearchPeopleResponse, error) {
	// todo: add your logic here and delete this line

	return &search.SearchPeopleResponse{}, nil
}
