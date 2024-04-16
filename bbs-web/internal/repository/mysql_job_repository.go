package repository

import (
	"bbs-web/internal/domain"
	"bbs-web/internal/repository/dao"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-15 21:19

type JobRepository interface {
	Preempt(ctx context.Context) (domain.Job, error)
	Release(ctx context.Context, id int64) error
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

func (j *jobRepository) toDomain(model dao.JobModel) domain.Job {
	return domain.Job{
		ID:           model.ID,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
		Cfg:          model.Cfg,
		Status:       model.Status,
		NextExecTime: model.NextExecTime,
		Version:      model.Version,
	}
}
