package dao

import (
	"context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-10 12:50

type InteractiveDao interface {
	IncrReadCnt(ctx context.Context, biz string, bizId int64) error
	BatchIncrReadCnt(ctx context.Context, bizs []string, bizIds []int64) error
	IncrLikeInfo(ctx context.Context, biz string, id int64, uid int64) error
	DelLikeInfo(ctx context.Context, biz string, bizId int64, uid int64) error
	IncrLikeCnt(ctx context.Context, biz string, id int64, uid int64) error
	DecrLikeCnt(ctx context.Context, biz string, id int64, uid int64) error
	Get(ctx context.Context, biz string, id int64) (InteractiveModel, error)
	GetLikeInfo(ctx context.Context, biz string, id int64, uid int64) (UserLikeBizModel, error)
	GetCollectInfo(ctx context.Context, biz string, id int64, uid int64) (UserCollectBizModel, error)
}

type interactiveDao struct {
	db *gorm.DB
}

func NewInteractiveDao(db *gorm.DB) InteractiveDao {
	return &interactiveDao{
		db: db,
	}
}

func (dao *interactiveDao) IncrReadCnt(ctx context.Context, biz string, bizId int64) error {
	// 查看是否有点赞记录，如果有就将read_cnt字段+1 ，否则就创建一行记录，并将read_cnt 设置为 1 注意并发问题
	err := dao.db.WithContext(ctx).Model(&UserLikeBizModel{}).
		Clauses(clause.OnConflict{
			DoUpdates: clause.Assignments(map[string]interface{}{
				"read_cnt": gorm.Expr("`read_cnt` + 1"),
			}),
		}).
		Create(&UserCollectBizModel{
			BizId: bizId,
			Biz:   biz,
			Uid:   1,
		}).Error
	return err
}

func (dao *interactiveDao) BatchIncrReadCnt(ctx context.Context, bizs []string, bizIds []int64) error {
	//TODO implement me
	panic("implement me")
}

func (dao *interactiveDao) IncrLikeInfo(ctx context.Context, biz string, id int64, uid int64) error {
	//TODO implement me
	panic("implement me")
}

func (dao *interactiveDao) DelLikeInfo(ctx context.Context, biz string, bizId int64, uid int64) error {
	//TODO implement me
	panic("implement me")
}

func (dao *interactiveDao) IncrLikeCnt(ctx context.Context, biz string, id int64, uid int64) error {
	//TODO implement me
	panic("implement me")
}

func (dao *interactiveDao) DecrLikeCnt(ctx context.Context, biz string, id int64, uid int64) error {
	//TODO implement me
	panic("implement me")
}

func (dao *interactiveDao) Get(ctx context.Context, biz string, id int64) (InteractiveModel, error) {
	//TODO implement me
	panic("implement me")
}

func (dao *interactiveDao) GetLikeInfo(ctx context.Context, biz string, id int64, uid int64) (UserLikeBizModel, error) {
	//TODO implement me
	panic("implement me")
}

func (dao *interactiveDao) GetCollectInfo(ctx context.Context, biz string, id int64, uid int64) (UserCollectBizModel, error) {
	//TODO implement me
	panic("implement me")
}
