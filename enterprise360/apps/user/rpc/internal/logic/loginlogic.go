package logic

import (
	"context"
	"enterprise360/apps/user/model"
	"google.golang.org/grpc/status"

	"enterprise360/apps/user/rpc/internal/svc"
	"enterprise360/apps/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	res, err := l.svcCtx.DB.FindOneByUsername(l.ctx, in.Username)
	if err == model.ErrNotFound {
		return nil, status.Error(500, "用户名不存在")
	}
	if err != nil {
		logx.WithContext(l.ctx).Error("登录模块根据用户名查询DB错误，存疑")
		return nil, status.Error(500, "系统错误")
	}

	return &user.LoginResponse{
		Id:       res.Id,
		Username: res.Username,
		Password: res.Password,
	}, nil
}
