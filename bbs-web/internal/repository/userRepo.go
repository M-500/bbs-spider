package repository

import (
	"bbs-web/internal/domain"
	"bbs-web/internal/repository/dao"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 15:25

type IUserRepo interface {
	CreateUser(ctx context.Context, u domain.UserInfo) error
}

type userRepo struct {
	dao dao.IUserDao
}

func (repo *userRepo) CreateUser(ctx context.Context, u domain.UserInfo) error {
	//TODO implement me
	panic("implement me")
}

func NewUserRepo(dao dao.IUserDao) IUserRepo {
	return &userRepo{dao: dao}
}
