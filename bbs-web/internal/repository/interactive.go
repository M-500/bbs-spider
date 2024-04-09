package repository

import (
	"bbs-web/internal/domain"
	"bbs-web/internal/repository/cache"
	"bbs-web/internal/repository/dao"
	"bbs-web/pkg/logger"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-07 19:28

type InteractiveRepo interface {
	IncrReadCnt(ctx context.Context, biz string, bizId int64) error
	// BatchIncrReadCnt biz 和 bizId 长度必须一致
	BatchIncrReadCnt(ctx context.Context, biz []string, bizId []int64) error
	IncrLike(ctx context.Context, biz string, id int64, uid int64) error
	DecrLike(ctx context.Context, biz string, id int64, uid int64) error
	AddCollectionItem(ctx context.Context, biz string, id int64, cid int64, uid int64) error
	Get(ctx context.Context, biz string, id int64) (domain.Interactive, error)
	Liked(ctx context.Context, biz string, id int64, uid int64) (bool, error)
	Collected(ctx context.Context, biz string, id int64, uid int64) (bool, error)
}

type interactiveRepo struct {
	dao   dao.InteractiveDao
	cache cache.RedisInteractiveCache
	l     logger.Logger
}

func NewInteractiveRepo(dao dao.InteractiveDao, cache cache.RedisInteractiveCache, l logger.Logger) InteractiveRepo {
	return &interactiveRepo{
		dao:   dao,
		cache: cache,
		l:     l,
	}
}

func (repo *interactiveRepo) toDomain(model dao.InteractiveModel) domain.Interactive {
	return domain.Interactive{
		ReadCnt:    model.ReadCnt,
		LikeCnt:    model.LikeCnt,
		CollectCnt: model.CollectCnt,
		CommentCnt: model.CommentCnt,
	}

}
func (repo *interactiveRepo) AddCollectionItem(ctx context.Context, biz string, id int64, cid int64, uid int64) error {
	panic("")
}
func (repo *interactiveRepo) BatchIncrReadCnt(ctx context.Context, biz []string, bizId []int64) error {
	// 要不要检测 biz 和 bizId的长度是否相等
	err := repo.dao.BatchIncrReadCnt(ctx, biz, bizId)

	return err
}
func (repo *interactiveRepo) Get(ctx context.Context, biz string, bizId int64) (domain.Interactive, error) {
	// 要从缓存中拿出阅读数量 点赞 收藏 评论等
	intr, err := repo.cache.Get(ctx, biz, bizId)
	if err == nil {
		return intr, err
	}
	// 从数据库拿出阅读数量
	intrDao, err := repo.dao.Get(ctx, biz, bizId)
	if err != nil {
		return domain.Interactive{}, err
	}
	intr = repo.toDomain(intrDao)
	// 要不要同步缓存中的数据？肯定要的 其实可以异步的对吧
	//err = repo.cache.Set(ctx, biz, bizId, intr)
	go func() {
		err2 := repo.cache.Set(ctx, biz, bizId, intr)
		// 容错写法，直接打印一个日志就告辞了  对于点赞数这种缓存，业务是可以允许有差别的，而且影响不大
		repo.l.Error("写入缓存失败", logger.Error(err2), logger.String("biz", biz), logger.Int64("bizId", bizId))
	}()
	return intr, nil
}

func (repo *interactiveRepo) Liked(ctx context.Context, biz string, bizId int64, uid int64) (bool, error) {
	_, err := repo.dao.GetLikeInfo(ctx, biz, bizId, uid)
	switch err {
	case nil:
		return true, nil
	case dao.ErrRecordNotFound:
		return false, nil
	default:
		return false, err
	}

}
func (repo *interactiveRepo) Collected(ctx context.Context, biz string, bizId int64, uid int64) (bool, error) {
	_, err := repo.dao.GetCollectInfo(ctx, biz, bizId, uid)
	switch err {
	case nil:
		return true, nil
	case dao.ErrRecordNotFound:
		return false, nil
	default:
		return false, err
	}
}

func (repo *interactiveRepo) IncrReadCnt(ctx context.Context, biz string, bizId int64) error {
	// 要考虑缓存方案了

	return repo.dao.IncrReadCnt(ctx, biz, bizId)
}

func (repo *interactiveRepo) IncrLike(ctx context.Context, biz string, id int64, uid int64) error {
	// 插入数据库 更新点赞计数 更新缓存
	err := repo.dao.IncrLikeInfo(ctx, biz, id, uid)
	if err != nil {
		return err
	}
	// 同步缓存
	return repo.cache.IncrLikeCntIfPresent(ctx, biz, id)

}
func (repo *interactiveRepo) DecrLike(ctx context.Context, biz string, id int64, uid int64) error {
	// 插入数据库
	err := repo.dao.DelLikeInfo(ctx, biz, id, uid)
	if err != nil {
		return err
	}
	// 同步缓存
	return repo.cache.DecrLikeCntIfPresent(ctx, biz, id)
}
