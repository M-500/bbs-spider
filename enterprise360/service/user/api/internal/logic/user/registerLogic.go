package user

import (
	"context"
	"enterprise360/service/user/rpc/types/user"

	"enterprise360/service/user/api/internal/svc"
	"enterprise360/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// register
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	u := user.RegUserRequest{
		Username: req.UserName,
		Password: req.Password,
		Mobile:   req.Phone,
	}
	res, err := l.svcCtx.UserRpc.Register(l.ctx, &u)
	if err != nil {
		return nil, err
	}
	return &types.RegisterResp{
		Id: res.Id,
	}, nil
}
