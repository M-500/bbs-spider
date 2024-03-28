package service

import (
	"bbs-web/internal/domain"
	"bbs-web/internal/repository/dao"
	"context"
	"errors"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 15:24

type IUserService interface {
	CreateUser(ctx context.Context, username, pwd string) (domain.UserInfo, error)
}

type userService struct {
	dao dao.IUserDao
}

func (u *userService) CreateUser(ctx context.Context, username, pwd string) (domain.UserInfo, error) {
	// 1 查询数据库的用户名是否存在
	exist, err := u.dao.FindByUserName(ctx, username)
	if err != nil {
		return domain.UserInfo{}, err
	}
	if exist {
		return domain.UserInfo{}, errors.New("用户已存在")
	}
	// 2. 创建用户
	// pwd 要加密
	insert, err := u.dao.Insert(ctx, dao.UserMode{
		Username: username,
		Password: pwd,
		Nickname: "用户saduio",
		IsAdmin:  1,
	})
	if err != nil {

	}
	return insert, nil
}

func NewUserService() IUserService {
	return &userService{}
}
