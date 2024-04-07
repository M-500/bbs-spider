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

type InteractiveDao interface {
	IncrReadCnt(ctx context.Context, biz string, bizId int64) error
	IncrLikeInfo(ctx context.Context, biz string, id int64, uid int64) error
	DelLikeInfo(ctx context.Context, biz string, bizId int64) error
}

type interactiveDao struct {
	db *gorm.DB
}

func (dao *interactiveDao) DelLikeInfo(ctx context.Context, biz string, bizId int64) error {

}

func (dao *interactiveDao) IncrLikeInfo(ctx context.Context, biz string, id int64, uid int64) error {
	return dao.db.WithContext(ctx).Clauses(
		clause.OnConflict{
			DoUpdates: clause.Assignments(map[string]any{
				"like_cnt":   gorm.Expr("read_cnt +1"),
				"updated_at": time.Now(),
			}),
		}).Create(&InteractiveModel{
		Biz:        biz,
		BizId:      id,
		ReadCnt:    0,
		LikeCnt:    1,
		CollectCnt: 0,
		CommentCnt: 0,
	}).Error
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
	return dao.db.WithContext(ctx).Clauses(
		clause.OnConflict{
			DoUpdates: clause.Assignments(map[string]any{
				"read_cnt":   gorm.Expr("read_cnt +1"),
				"updated_at": time.Now(),
			}),
		}).Create(&InteractiveModel{
		Biz:        biz,
		BizId:      bizId,
		ReadCnt:    1,
		LikeCnt:    0,
		CollectCnt: 0,
		CommentCnt: 0,
	}).Error
}

func NewInteractiveDao() InteractiveDao {
	return &interactiveDao{}
}
