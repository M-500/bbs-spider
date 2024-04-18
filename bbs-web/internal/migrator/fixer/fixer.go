package fixer

import (
	"bbs-web/internal/migrator"
	"bbs-web/internal/migrator/events"
	"context"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// @Description
// @Author 代码小学生王木木

type Fixer[T migrator.Entity] struct {
	base    *gorm.DB
	target  *gorm.DB
	columns []string
}

// 有并发问题，但是问题不大！因为这个修复的逻辑会重复执行

// 最好的写法 不管三七二十一 我TM直接覆盖
func (f *Fixer[T]) Fix(ctx context.Context, evt events.InconsistentEvent) error {
	// 1. 先看源数据库有没有
	var t T
	err := f.base.WithContext(ctx).Where("id = ?", evt.ID).First(&t).Error
	switch err {
	case nil:
		// 说明base有数据 那么target要执行 upsert语义  如果有就更新，没有就插入
		return f.target.WithContext(ctx).Clauses(clause.OnConflict{
			DoUpdates: clause.AssignmentColumns(f.columns),
		}).Create(&t).Error
	case gorm.ErrRecordNotFound:
		// base没有数据了，那就去查target 删之
		return f.target.WithContext(ctx).Where("id = ?", evt.ID).Delete(&t).Error
	default:
		return errors.New("未知的不一致类型")
	}
}

// 好处: base和target在校验数据的时候，等你修复的时候不会变
func (f *Fixer[T]) FixV1(ctx context.Context, evt events.InconsistentEvent) error {
	switch evt.Type {
	case events.InconsistentEventTypeNEQ,
		events.InconsistentEventTypeTargetMissing:
		var t T
		err := f.base.WithContext(ctx).Where("id = ?", evt.ID).First(&t).Error
		switch err {
		case gorm.ErrRecordNotFound:
			// base找不到这行记录，那么target就要删除这一行记录
			return f.target.WithContext(ctx).Where("id = ?", evt.ID).Delete(&t).Error
		case nil:
			// 执行更新 需要Upsert语义
			return f.target.WithContext(ctx).Clauses(clause.OnConflict{
				//DoUpdates: clause.Assignments(map[string]interface{}{}),
				DoUpdates: clause.AssignmentColumns(f.columns),
			}).Where("id = ?", evt.ID).Updates(&t).Error
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

func (f *Fixer[T]) FixV2(ctx context.Context, evt events.InconsistentEvent) error {
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
