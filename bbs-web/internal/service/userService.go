package service

import (
	"bbs-web/internal/domain"
	"bbs-web/internal/repository"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 15:24

type IUserService interface {
	SignUp(ctx context.Context, user domain.UserInfo) error
	Login(ctx context.Context, username, password string) (domain.UserInfo, error)
}

type userService struct {
	repo repository.IUserRepo
}

func (u *userService) SignUp(ctx context.Context, user domain.UserInfo) error {
	// 加密在这里做
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	//user.B
	// 随机昵称也在这里做
	return u.repo.CreateUser(ctx, user)
}

func (u *userService) Login(ctx context.Context, username, password string) (domain.UserInfo, error) {
	user, err := u.repo.FindByUsername(ctx, username)
	if err != nil {
		return domain.UserInfo{}, errors.New("用户名不存在")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// DEBUG
		return domain.UserInfo{}, errors.New("密码不对")
	}

	return user, nil
}

func NewUserService(repo repository.IUserRepo) IUserService {
	return &userService{
		repo: repo,
	}
}
