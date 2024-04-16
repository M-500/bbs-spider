//@Author: wulinlin
//@Description:
//@File:  job_dao
//@Version: 1.0.0
//@Date: 2024/04/15 22:18

package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type JobDAO interface {
	Preempt(ctx context.Context) (JobModel, error)
	Release(ctx context.Context, id int64) error
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

func NewJobDAO(db *gorm.DB) JobDAO {
	return &gormJobDAO{
		db: db,
	}
}

func (g *gormJobDAO) Preempt(ctx context.Context) (JobModel, error) {
	for {
		now := time.Now()
		var res JobModel
		err := g.db.WithContext(ctx).
			Where("status = ? AND next_exec_time < ?", jobStatusWaiting, now).
			First(&res).Error
		if err != nil {
			// 这个处理有问题
			return JobModel{}, err
		}
		// 分布式任务调度系统
		// 1. 一次拉一批 2. 随机从某一条开始，从后开始抢占  3. 随机偏移量+取模/第一轮没有查到，偏移量归零

		// 这里是乐观锁实现 CAS  Compare And Swap  常见的面试装逼 => 用乐观锁取代 for update   forupdate性能差，容易引起死锁问题。 性能优化的套路
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
			//return JobModel{}, errors.New("没有抢到")
			continue
		}
		return res, err
	}
}

func (g *gormJobDAO) Release(ctx context.Context, id int64) error {

	return nil
}
