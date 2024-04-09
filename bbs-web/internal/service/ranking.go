package service

import (
	"bbs-web/internal/domain"
	"bbs-web/internal/service/article"
	"context"
	"log"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-28 17:06

type RankingService interface {
	TopN(ctx context.Context) error
	topN(ctx context.Context) ([]domain.Article, error)
	//TopN(ctx context.Context,n int64) error
	//TopN(ctx context.Context, n int64) ([]any, error) // 后续写这个逻辑
}

type batchRankingService struct {
	artSvc  article.IArticleService
	intrSvc InteractiveService
}

func NewBatchRankingService(artSvc article.IArticleService, intrSvc InteractiveService) RankingService {
	return &batchRankingService{
		artSvc:  artSvc,
		intrSvc: intrSvc,
	}
}
func (svc *batchRankingService) TopN(ctx context.Context) error {
	n, err := svc.topN(ctx)
	if err != nil {
		return err
	}
	// 存起来
	log.Println(n)
	return err
}

func (svc *batchRankingService) topN(ctx context.Context) ([]domain.Article, error) {
	var (
		offset = 15
		limit  = 0
	)
	// 1. 先拿一批数据
	for {
		now := time.Now()
		_, err := svc.artSvc.ListPub(ctx, now, offset, limit)
		if err != nil {

		}
		// 要去对应的点赞数据

		//svc.artSvc.GetByIds(ctx, "", )

		// 合并计算socre
		// 排序
		// 最后得出结果
	}
	return nil, nil
}
