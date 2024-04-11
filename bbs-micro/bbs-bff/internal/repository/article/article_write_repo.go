package article

import (
	"bbs-micro/bbs-bff/internal/domain"
	"bbs-micro/bbs-bff/internal/repository/dao/article_dao"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-02 10:00

type ArtWriterRepo interface {
	Create(ctx context.Context, art domain.Article) (int64, error)
	Update(ctx context.Context, art domain.Article) error
}

type artWriter struct {
	dao article_dao.WriteDAO
}

func (a *artWriter) Create(ctx context.Context, art domain.Article) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (a *artWriter) Update(ctx context.Context, art domain.Article) error {
	//TODO implement me
	panic("implement me")
}

func NewArtWriterRepo(dao article_dao.WriteDAO) ArtWriterRepo {
	return &artWriter{
		dao: dao,
	}
}
