package service

import (
	"bbs-web/internal/domain"
	"bbs-web/internal/repository"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 15:24

type IUserService interface {
	SignUp(ctx context.Context, user domain.UserInfo) error
}

type userService struct {
	repo repository.IUserRepo
}

func (u *userService) SignUp(ctx context.Context, user domain.UserInfo) error {
	return u.repo.CreateUser(ctx, user)
}

func NewUserService(repo repository.IUserRepo) IUserService {
	return &userService{
		repo: repo,
	}
}
