package repository

import (
	"bbs-web/internal/repository/dao"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-07 19:28

type InteractiveRepo interface {
	IncrReadCnt(ctx context.Context, biz string, bizId int64) error
}

type interactiveRepo struct {
	dao dao.InteractiveDao
}

func (repo *interactiveRepo) IncrReadCnt(ctx context.Context, biz string, bizId int64) error {
	// 要考虑缓存方案了

	return nil
}
