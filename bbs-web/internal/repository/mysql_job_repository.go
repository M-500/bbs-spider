package repository

import (
	"bbs-web/internal/domain"
	"bbs-web/internal/repository/dao"
	"context"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-15 21:19

type JobRepository interface {
	Preempt(ctx context.Context) (domain.Job, error)
	Release(ctx context.Context, id int64) error
	UpdateUpdateTime(ctx context.Context, id int64) error
	UpdateNextTime(ctx context.Context, id int64, time time.Time) error
	Stop(ctx context.Context, id int64) error
}

type jobRepository struct {
	dao dao.JobDAO
}

func (j *jobRepository) Preempt(ctx context.Context) (domain.Job, error) {
	jb, err := j.dao.Preempt(ctx)
	if err != nil {

	}

	return j.toDomain(jb), err
}

func (j *jobRepository) Release(ctx context.Context, id int64) error {
	return j.dao.Release(ctx, id)
}

func (j *jobRepository) UpdateUpdateTime(ctx context.Context, id int64) error {
	return j.dao.UpdateChangeTime(ctx, id)
}

func (j *jobRepository) UpdateNextTime(ctx context.Context, id int64, t time.Time) error {
	return j.dao.UpdateNextTime(ctx, id, t)
}

func (j *jobRepository) Stop(ctx context.Context, id int64) error {
	return nil
}

func (j *jobRepository) toDomain(model dao.JobModel) domain.Job {
	return domain.Job{
		ID:           int64(model.ID),
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
		Cfg:          model.Cfg,
		Status:       model.Status,
		NextExecTime: model.NextExecTime,
		Version:      model.Version,
	}
}
