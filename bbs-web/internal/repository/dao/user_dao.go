package dao

import (
	"bbs-web/internal/domain"
	"context"
	"gorm.io/gorm"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 15:13

type IUserDao interface {
	FindById(ctx context.Context, uid int64) (UserMode, error)
	FindByUserName(ctx context.Context, username string) (bool, error)
	Insert(ctx context.Context, mode UserMode) (domain.UserInfo, error)
}

type userDao struct {
	db *gorm.DB
}

func (dao *userDao) FindById(ctx context.Context, uid int64) (UserMode, error) {
	var user UserMode
	err := dao.db.WithContext(ctx).Where("id = ?", uid).First(&user).Error
	return user, err
}
func (dao *userDao) FindByUserName(ctx context.Context, username string) (bool, error) {
	tx := dao.db.WithContext(ctx).Where("username = ?", username)
	if tx.Error != nil {
		return false, tx.Error
	}
	if tx.RowsAffected >= 1 {
		return true, nil
	}
	return false, nil
}

func (dao *userDao) Insert(ctx context.Context, mode UserMode) (domain.UserInfo, error) {
	err := dao.db.WithContext(ctx).Create(&mode).Error
	if err != nil {
		return domain.UserInfo{}, err
	}
	return domain.UserInfo{
		Id:       int64(mode.ID),
		UserName: mode.Username,
		NickName: mode.Nickname,
	}, nil
}
func NewUserDao(db *gorm.DB) IUserDao {
	return &userDao{
		db: db,
	}
}
