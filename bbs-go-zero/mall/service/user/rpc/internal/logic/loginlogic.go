package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"mall/common/cryptx"
	"mall/service/user/rpc/internal/svc"
	"mall/service/user/rpc/types/user"

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
	// 1. 查询数据库中是否存在该用户
	exist, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil {
		// 这里可以对错误进行进一步细分
		return nil, status.Error(500, "系统错误")
	}
	// 2. 校验密码是否正确
	encryptPwd := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if encryptPwd != exist.Password {
		return nil, status.Error(500, "密码错误")
	}
	return &user.LoginResponse{
		Id:     int64(exist.Id),
		Name:   exist.Name,
		Gender: int64(exist.Gender),
		Mobile: exist.Mobile,
	}, nil
}
