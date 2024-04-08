package dao

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-07 19:32

var (
	ErrRecordNotFound = gorm.ErrRecordNotFound
)

type InteractiveDao interface {
	IncrReadCnt(ctx context.Context, biz string, bizId int64) error
	IncrLikeInfo(ctx context.Context, biz string, id int64, uid int64) error
	DelLikeInfo(ctx context.Context, biz string, bizId int64, uid int64) error
	IncrLikeCnt(ctx context.Context, biz string, id int64, uid int64) error
	DecrLikeCnt(ctx context.Context, biz string, id int64, uid int64) error

	Get(ctx context.Context, biz string, id int64) (InteractiveModel, error)
	GetLikeInfo(ctx context.Context, biz string, id int64, uid int64) (UserLikeBizModel, error)
	GetCollectInfo(ctx context.Context, biz string, id int64, uid int64) (UserLikeBizModel, error)
}

type interactiveDao struct {
	db *gorm.DB
}

func NewInteractiveDao(db *gorm.DB) InteractiveDao {
	return &interactiveDao{
		db: db,
	}
}

// DelLikeInfo
//
//	@Description: 删除点赞
func (dao *interactiveDao) DelLikeInfo(ctx context.Context, biz string, bizId int64, uid int64) error {
	now := time.Now()
	err := dao.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1. 软删除某一条数据
		err := tx.Model(&UserLikeBizModel{}).
			Where("biz_id = ? AND biz = ? AND uid = ?", bizId, biz, uid).
			Updates(map[string]any{
				"deleted_at": now,
				"updated_at": now,
			}).Error
		if err != nil {
			return err
		}
		// 2. 总数-1
		return tx.Model(&InteractiveModel{}).Where("biz_id = ? AND biz = ? ", bizId, biz).
			Updates(map[string]any{
				"updated_at": now,
				"like_cnt":   gorm.Expr("like_cnt - 1"),
			}).Error
	})
	return err
}

// IncrLikeInfo
//
//	@Description: 新增点赞 以及更新点赞记录  你需要一张表来记录谁给某一篇文章点了赞
func (dao *interactiveDao) IncrLikeInfo(ctx context.Context, biz string, id int64, uid int64) error {
	now := time.Now()
	err := dao.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 是否需要校验重复点赞的问题？不用
		err := tx.Model(&UserLikeBizModel{}).Clauses(
			clause.OnConflict{
				DoUpdates: clause.Assignments(map[string]any{
					"updated_at": now,
				}),
			}).Create(&UserLikeBizModel{
			BizId: id,
			Biz:   biz,
			Uid:   uid,
		}).Error
		if err != nil {
			return err
		}
		// 更新点赞总数
		return tx.WithContext(ctx).Clauses(
			clause.OnConflict{
				DoUpdates: clause.Assignments(map[string]any{
					"like_cnt":   gorm.Expr("like_cnt + 1"),
					"updated_at": time.Now(),
				}),
			}).Create(&InteractiveModel{
			Biz:     biz,
			BizId:   id,
			LikeCnt: 1,
		}).Error
	})
	return err
}

func (dao *interactiveDao) IncrLikeCnt(ctx context.Context, biz string, id int64, uid int64) error {
	return nil
}

func (dao *interactiveDao) DecrLikeCnt(ctx context.Context, biz string, id int64, uid int64) error {
	return nil
}

func (dao *interactiveDao) IncrReadCnt(ctx context.Context, biz string, bizId int64) error {
	// 下面这种写法有大问题！ check do something
	//var intr InteractiveModel
	//err := dao.db.WithContext(ctx).Model(&InteractiveModel{}).Where("biz_id = ? AND biz = ?", bizId, biz).First(&intr).Error
	//if err != nil {
	//
	//}
	//cnt := intr.ReadCnt + 1
	//dao.db.WithContext(ctx).Updates(map[string]any{
	//	"read_cnt": cnt,
	//})
	// 数据库层 SQL支持update a = a + 1  实现Upsert语义
	createObj := InteractiveModel{
		Biz:     biz,
		BizId:   bizId,
		ReadCnt: 1,
	}
	return dao.db.WithContext(ctx).Clauses(
		clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.Assignments(map[string]any{
				"read_cnt":   gorm.Expr("read_cnt +1"),
				"updated_at": time.Now(),
			}),
		}).Create(&createObj).Error
}
func (dao *interactiveDao) Get(ctx context.Context, biz string, id int64) (InteractiveModel, error) {
	var data InteractiveModel
	err := dao.db.WithContext(ctx).Model(&InteractiveModel{}).Where("biz_id=? AND biz=?", id, biz).First(&data).Error
	return data, err
}
func (dao *interactiveDao) GetLikeInfo(ctx context.Context, biz string, id int64, uid int64) (UserLikeBizModel, error) {
	panic("")
}
func (dao *interactiveDao) GetCollectInfo(ctx context.Context, biz string, id int64, uid int64) (UserLikeBizModel, error) {
	panic("")
}
