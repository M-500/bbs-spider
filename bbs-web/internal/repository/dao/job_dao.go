//@Author: wulinlin
//@Description:
//@File:  job_dao
//@Version: 1.0.0
//@Date: 2024/04/15 22:18

package dao

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

type JobDAO interface {
	Preempt(ctx context.Context) (JobModel, error)
}

/**
考虑一下问题
1. 那些任务可以抢？
2. 那些任务已经被占用了
3. 哪些任务永远不会被调度
*/

type gormJobDAO struct {
	db *gorm.DB
}

func (g *gormJobDAO) Preempt(ctx context.Context) (JobModel, error) {
	now := time.Now()
	var res JobModel
	err := g.db.WithContext(ctx).
		Where("status = ? AND next_exec_time < ?", jobStatusWaiting, now).
		First(&res).Error
	if err != nil {
		// 这个处理有问题
		return JobModel{}, err
	}
	ans := g.db.WithContext(ctx).Where("id = ? AND version = ?", res.ID, res.Version).
		Updates(map[string]any{
			"status":     jobStatusRunning,
			"version":    res.Version + 1,
			"updated_at": now,
		})
	if ans.Error != nil {
		return JobModel{}, err
	}
	if ans.RowsAffected == 0 {
		return JobModel{}, errors.New("没有抢到")
	}
	return res, err
}
