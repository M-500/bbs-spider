package repository

import (
	"bbs-web/internal/domain"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-15 21:19

type JobRepository interface {
	Preempt(ctx context.Context) (domain.Job, error)
}

type jobRepository struct {
}

func (j *jobRepository) Preempt(ctx context.Context) (domain.Job, error) {
	//TODO implement me
	panic("implement me")
}
