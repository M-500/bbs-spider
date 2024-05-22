package logic

import (
	"context"
	"database/sql"
	"enterprise360/apps/user/model"
	"time"

	"enterprise360/apps/user/rpc/internal/svc"
	"enterprise360/apps/user/rpc/types/user"

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

func (l *RegisterLogic) Register(in *user.RegUserRequest) (*user.RegUserResponse, error) {
	now := time.Now()
	var u = model.EbUser{
		Username: in.Username,
		Password: in.Password,
		Mobile: sql.NullString{
			String: in.Mobile,
			Valid:  false,
		},
		CreateTime: now,
		UpdateTime: now,
	}
	res, err := l.svcCtx.DB.Insert(l.ctx, &u)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &user.RegUserResponse{
		Id: id,
	}, nil
}
