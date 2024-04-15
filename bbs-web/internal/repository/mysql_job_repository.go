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
}

type jobRepository struct {
	dao dao.JobDAO
}

func (j *jobRepository) Preempt(ctx context.Context) (domain.Job, error) {
	jb, err := j.dao.Preempt(ctx)
	if err != nil {

	}
	return domain.Job{}, err
}
