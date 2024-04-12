package repository

import (
	"bbs-micro/bbs-bff/internal/domain"
	"bbs-micro/bbs-bff/internal/repository/dao"
	"bbs-micro/pkg/utils/zifo/slice"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-12 12:08

type ICollectRepo interface {
	GetCollectListByID(ctx context.Context, uid, limit, offset int64) ([]domain.Collect, error)
	CreateCollect(ctx context.Context, uid int64, cname string, desc string, isPub bool) (int64, error)
}

type collectRepo struct {
	dao dao.ICollectDAO
}

func NewCollectRepo(dao dao.ICollectDAO) ICollectRepo {
	return &collectRepo{dao: dao}
}

func (c *collectRepo) GetCollectListByID(ctx context.Context, uid, limit, offset int64) ([]domain.Collect, error) {
	list, err := c.dao.QueryCollectList(ctx, uid, limit, offset)
	if err != nil {
		return nil, err
	}
	return slice.Map[dao.CollectionModle, domain.Collect](list, func(idx int, src dao.CollectionModle) domain.Collect {
		return c.toDomain(src)
	}), nil
}
func (c *collectRepo) CreateCollect(ctx context.Context, uid int64, cname string, desc string, isPub bool) (int64, error) {
	return c.dao.InsertCollect(ctx, uid, cname, desc, isPub)
}

func (c *collectRepo) toDomain(item dao.CollectionModle) domain.Collect {
	return domain.Collect{
		UserId:      item.UserId,
		CName:       item.CName,
		Description: item.Description,
		Sort:        item.Sort,
		ResourceNum: item.ResourceNum,
		IsPub:       item.IsPub,
		CommentNum:  item.CommentNum,
		CreateTime:  item.CreatedAt,
		UpdateTime:  item.UpdatedAt,
	}
}
