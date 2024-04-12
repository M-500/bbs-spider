package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-12 12:09

var (
	ErrCollectDuplicate = errors.New("收藏夹名字冲突")
)

type ICollectDAO interface {
	QueryCollectList(ctx context.Context, uid, limit, offset int64) ([]CollectionModle, error)

	InsertCollect(ctx context.Context, uid int64, cname string, desc string, isPub bool) (int64, error)

	InsertCollectToBiz(ctx context.Context, biz string, uid, cid, bizId int64) (int64, error)
}

type collectDao struct {
	db *gorm.DB
}

func NewCollectDao(db *gorm.DB) ICollectDAO {
	return &collectDao{db: db}
}

func (dao *collectDao) QueryCollectList(ctx context.Context, uid, limit, offset int64) ([]CollectionModle, error) {
	var collectList []CollectionModle
	err := dao.db.WithContext(ctx).Model(&CollectionModle{}).Where("user_id = ?", uid).Find(&collectList).Error
	return collectList, err
}

func (dao *collectDao) InsertCollect(ctx context.Context, uid int64, cname string, desc string, isPub bool) (int64, error) {
	// upsert 语义 ? 还是依赖唯一索引的冲突？？？
	data := CollectionModle{
		UserId:      uid,
		CName:       cname,
		Description: desc,
		IsPub:       isPub,
	}
	//dao.db.WithContext(ctx).Where("user_id = ? AND c_name = ?", uid, cname)
	err := dao.db.WithContext(ctx).Create(&data).Error
	if sqlError, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo uint16 = 1062
		if sqlError.Number == uniqueConflictsErrNo {
			// 邮箱冲突 or 手机号码冲突
			return 0, ErrCollectDuplicate
		}
	}
	if err != nil {
		return 0, err
	}
	return int64(data.ID), nil
}

func (dao *collectDao) InsertCollectToBiz(ctx context.Context, biz string, uid, cid, bizId int64) (int64, error) {
	// upsert语义 并且操作两张表
	//dao.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
	//	newDao := NewCollectDao(tx)
	//
	//})
	return 0, nil
}
