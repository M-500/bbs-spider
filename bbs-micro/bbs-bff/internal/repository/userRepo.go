package repository

import (
	"bbs-micro/bbs-bff/internal/domain"
	"bbs-micro/bbs-bff/internal/repository/dao"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 15:25

type IUserRepo interface {
	CreateUser(ctx context.Context, u domain.UserInfo) error
	FindByUsername(ctx context.Context, username string) (domain.UserInfo, error)
	FindById(ctx context.Context, id int64) (domain.UserInfo, error)
}

type userRepo struct {
	dao dao.IUserDao
}

func (repo *userRepo) FindByUsername(ctx context.Context, username string) (domain.UserInfo, error) {
	user, err := repo.dao.FindByUserName(ctx, username)
	if err != nil {
		return domain.UserInfo{}, err
	}
	return repo.toDomain(user), nil
}

func (repo *userRepo) CreateUser(ctx context.Context, u domain.UserInfo) error {
	return repo.dao.Insert(ctx, dao.UserMode{
		Username: u.UserName,
		Password: u.Password,
		Nickname: u.UserName,
		IsAdmin:  0,
	})
}

func (repo *userRepo) FindById(ctx context.Context, id int64) (domain.UserInfo, error) {
	user, err := repo.dao.FindById(ctx, id)
	return repo.toDomain(user), err
}

func (repo *userRepo) toDomain(u dao.UserMode) domain.UserInfo {
	return domain.UserInfo{
		Id:       int64(u.Model.ID),
		UserName: u.Username,
		NickName: u.Nickname,
		Password: u.Password,
		BirthDay: u.Birthday,
		Avatar:   u.Avatar,
	}
}

func NewUserRepo(dao dao.IUserDao) IUserRepo {
	return &userRepo{dao: dao}
}
