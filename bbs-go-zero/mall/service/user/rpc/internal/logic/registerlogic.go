package logic

import (
	"context"
	"google.golang.org/grpc/status"
	"mall/common/cryptx"
	"mall/service/user/model"
	"time"

	"mall/service/user/rpc/internal/svc"
	"mall/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// 1. 判断手机号是否注册
	_, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err == nil {
		return nil, status.Error(100, "手机号已存在")
	}
	nowData := time.Now()
	if err == model.ErrNotFound {
		newUser := model.User{
			Name:       in.Name,
			Gender:     uint64(in.Gender),
			Mobile:     in.Mobile,
			Password:   cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
			CreateTime: nowData,
			UpdateTime: nowData,
		}
		insert, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		id, err := insert.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}
		return &user.RegisterResponse{
			Id:     id,
			Name:   newUser.Name,
			Gender: int64(newUser.Gender),
			Mobile: newUser.Mobile,
		}, nil
	}

	return &user.RegisterResponse{}, nil
}
