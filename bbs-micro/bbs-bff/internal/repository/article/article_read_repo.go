package article

import (
	"bbs-micro/bbs-bff/internal/domain"
	"bbs-micro/bbs-bff/internal/repository/dao/article_dao"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-02 10:02

type ArticleReaderRepository interface {
	// Save 有就更新，没有就新建，即 upsert 的语义
	Save(ctx context.Context, art domain.Article) (int64, error)
}

type articleReader struct {
	dao article_dao.ReadDAO
}

func NewArticleReaderRepo(dao article_dao.ReadDAO) ArticleReaderRepository {
	return &articleReader{dao: dao}
}

func (a *articleReader) Save(ctx context.Context, art domain.Article) (int64, error) {
	//TODO implement me
	panic("implement me")
}
