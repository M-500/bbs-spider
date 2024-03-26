package dao

import (
	"context"
	"gorm.io/gorm"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 15:13

type IUserDao interface {
	FindById(ctx context.Context, uid int64) (UserMode, error)
}

type userDao struct {
	db *gorm.DB
}

func (dao *userDao) FindById(ctx context.Context, uid int64) (UserMode, error) {
	var user UserMode
	err := dao.db.WithContext(ctx).Where("id = ?", uid).First(&user).Error
	return user, err
}

func NewUserDao(db *gorm.DB) IUserDao {
	return &userDao{
		db: db,
	}
}
