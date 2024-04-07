package service

import (
	"bbs-web/internal/repository"
	"github.com/gin-gonic/gin"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-07 19:13

type InteractiveService interface {
	IncrReadCnt(ctx *gin.Context, biz string, id int64) error
}

type interactiveService struct {
	repo repository.InteractiveRepo
}

func NewInteractiveService(repo repository.InteractiveRepo) InteractiveService {
	return &interactiveService{
		repo: repo,
	}
}

func (i *interactiveService) IncrReadCnt(ctx *gin.Context, biz string, id int64) error {
	return i.repo.IncrReadCnt(ctx, biz, id)
}
