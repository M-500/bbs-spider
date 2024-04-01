package service

import (
	"bbs-web/internal/repository/article"
	"context"
	"time"

	"bbs-web/internal/domain"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 15:59

type IArticleService interface {
	Save(ctx context.Context, art domain.Article) (int64, error)
	Withdraw(ctx context.Context, art domain.Article) error
	Publish(ctx context.Context, art domain.Article) (int64, error)
	PublishV1(ctx context.Context, art domain.Article) (int64, error)
	List(ctx context.Context, uid int64, offset int, limit int) ([]domain.Article, error)
	ListPub(ctx context.Context, start time.Time, offset, limit int) ([]domain.Article, error)
	GetById(ctx context.Context, id int64) (domain.Article, error)
	GetByIds(ctx context.Context, biz string, ids []int64) ([]domain.Article, error)
	GetPublishedById(ctx context.Context, id, uid int64) (domain.Article, error)
}

type articleService struct {
	repo article.ArticleRepository
}

func NewArticleService(repo article.ArticleRepository) IArticleService {
	return &articleService{
		repo: repo,
	}
}

func (svc *articleService) Save(ctx context.Context, art domain.Article) (int64, error) {
	art.Status = domain.ArticleStatusUnpublished
	if art.Id > 0 {
		err := svc.repo.Update(ctx, art)
		return art.Id, err
	}
	return svc.repo.Create(ctx, art)
}

func (svc *articleService) Withdraw(ctx context.Context, art domain.Article) error {
	//TODO implement me
	panic("implement me")
}

func (svc *articleService) Publish(ctx context.Context, art domain.Article) (int64, error) {
	// 制作库新增数据
	id, err := svc.repo.Create(ctx, art)
	if err != nil {

	}
	// 线上库同步数据

}

func (svc *articleService) PublishV1(ctx context.Context, art domain.Article) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (svc *articleService) List(ctx context.Context, uid int64, offset int, limit int) ([]domain.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (svc *articleService) ListPub(ctx context.Context, start time.Time, offset, limit int) ([]domain.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (svc *articleService) GetById(ctx context.Context, id int64) (domain.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (svc *articleService) GetPublishedById(ctx context.Context, id, uid int64) (domain.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (svc *articleService) GetByIds(ctx context.Context, biz string, ids []int64) ([]domain.Article, error) {
	//TODO implement me
	panic("implement me")
}
