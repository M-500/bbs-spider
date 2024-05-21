package logic

import (
	"context"
	"encoding/json"
	"mall/service/user/rpc/types/user"

	"mall/service/user/api/internal/svc"
	"mall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// 从上下文中获取token
	userId, err := l.ctx.Value("uid").(json.Number).Int64()
	if err != nil {
		// 这个错误可以忽略，可以确保不会出错
	}
	info, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{Id: userId})
	if err != nil {
		return nil, err
	}
	return &types.UserInfoResponse{
		Id:     info.Id,
		Name:   info.Name,
		Gender: info.Gender,
		Mobile: info.Mobile,
	}, nil
}
