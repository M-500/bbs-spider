package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 15:13

var ErrUserDuplicate = errors.New("用户名冲突")

type IUserDao interface {
	FindById(ctx context.Context, uid int64) (UserMode, error)
	FindByUserName(ctx context.Context, username string) (UserMode, error)
	Insert(ctx context.Context, mode UserMode) error
}

type userDao struct {
	db *gorm.DB
}

func (dao *userDao) FindById(ctx context.Context, uid int64) (UserMode, error) {
	var user UserMode
	err := dao.db.WithContext(ctx).Where("id = ?", uid).First(&user).Error
	return user, err
}
func (dao *userDao) FindByUserName(ctx context.Context, username string) (UserMode, error) {
	var user UserMode
	err := dao.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
	return user, err
}

func (dao *userDao) Insert(ctx context.Context, mode UserMode) error {
	err := dao.db.WithContext(ctx).Create(&mode).Error
	// 一定会唯一索引冲突，这里要记录一下
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			// 邮箱冲突 or 手机号码冲突
			return ErrUserDuplicate
		}
	}
	if err != nil {
		return err
	}
	return nil
}
func NewUserDao(db *gorm.DB) IUserDao {
	return &userDao{
		db: db,
	}
}
