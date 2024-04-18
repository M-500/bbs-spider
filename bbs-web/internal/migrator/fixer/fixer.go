package fixer

import (
	"bbs-web/internal/migrator"
	"bbs-web/internal/migrator/events"
	"context"
	"errors"
	"gorm.io/gorm"
)

// @Description
// @Author 代码小学生王木木

type Fixer[T migrator.Entity] struct {
	base   *gorm.DB
	target *gorm.DB
}

func (f *Fixer[T]) Fix(ctx context.Context, evt events.InconsistentEvent) error {
	switch evt.Type {
	case events.InconsistentEventTypeNEQ:
		// 意味着要更新
		var t T
		err := f.base.WithContext(ctx).Where("id = ?", evt.ID).First(&t).Error
		switch err {
		case gorm.ErrRecordNotFound:
			// base找不到这行记录，那么target就要删除这一行记录
			return f.target.WithContext(ctx).Where("id = ?", evt.ID).Delete(&t).Error
		case nil:
			// 执行更新 需要Upsert语义
			return f.target.WithContext(ctx).Where("id = ?", evt.ID).Updates(&t).Error
		default:
			return err
		}

	case events.InconsistentEventTypeTargetMissing:
		// 意味着要插入
		var t T
		err := f.base.WithContext(ctx).Where("id = ?", evt.ID).First(&t).Error
		switch err {
		case gorm.ErrRecordNotFound:
			// base也删除了这条数据，不用插入了
			return nil
		case nil:
			return f.target.Create(&t).Error
		default:
			return err
		}
	case events.InconsistentEventTypeSourceMissing:
		// 意味着要删除
		return f.target.WithContext(ctx).Where("id = ?", evt.ID).Updates(new(T)).Error
	default:
	}
	return errors.New("未知的不一致类型")
}
