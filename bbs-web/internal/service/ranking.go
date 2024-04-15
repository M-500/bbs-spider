package service

import (
	"bbs-web/internal/domain"
	"bbs-web/internal/service/article"
	"bbs-web/pkg/utils/zifo/slice"
	"fmt"
	"math"

	"context"
	"github.com/ecodeclub/ekit/queue"
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
	artSvc    article.IArticleService
	intrSvc   InteractiveService
	batchSize int
	queueCap  int

	scoreFn func(t int64, likeCnt int64) float64 // 要求不能返回负数
}

func NewBatchRankingService(artSvc article.IArticleService, intrSvc InteractiveService) *batchRankingService {
	return &batchRankingService{
		artSvc:    artSvc,
		intrSvc:   intrSvc,
		batchSize: 100,
		queueCap:  100,
		scoreFn: func(t int64, likeCnt int64) float64 {
			// Hacknews的公式
			return float64(likeCnt-1) / math.Pow(float64(likeCnt+2), 1.5)
		},
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
		offset = 0
		start  = time.Now()
		ddl    = start.Add(-7 * 24 * time.Hour)
	)
	now := time.Now()
	type Score struct {
		art   domain.Article
		score float64
	}

	topQueue := queue.NewPriorityQueue[Score](svc.queueCap,
		func(src Score, dst Score) int {
			if src.score > dst.score {
				return 1
			} else if src.score == dst.score {
				return 0
			} else {
				return -1
			}
		})

	for {
		// 1. 拿一批数据
		arts, err := svc.artSvc.ListPub(ctx, now, offset, svc.batchSize)
		if err != nil {
			return nil, err
		}
		// 没取到数据，或者获取的数据不够一个批次，或者这批数据的最大实践都超出了计算范围
		if len(arts) < svc.batchSize || arts[len(arts)-1].Utime.Before(ddl) {
			break
		}
		// 转换为ID数组
		ids := slice.Map[domain.Article, int64](arts, func(idx int, src domain.Article) int64 {
			return src.Id
		})
		// 要去对应的点赞数据
		intrs, err := svc.intrSvc.GetByIds(ctx, "article", ids)
		if err != nil {
			return nil, err
		}
		// 合并计算socre
		// 排序
		for _, art := range arts {
			intr, ok := intrs[art.Id]
			if !ok {
				// 没有点赞数据
				continue
			}
			score := svc.scoreFn(art.Utime.Unix(), intr.LikeCnt+2) // +2 为了规避负数问题
			fmt.Println("当前的分数", art.Id, score)
			node := Score{
				art:   art,
				score: score,
			}
			err = topQueue.Enqueue(node)
			// 如果堆满了
			if err == queue.ErrOutOfCapacity {
				minNode, _ := topQueue.Dequeue()
				if minNode.score < score {
					_ = topQueue.Enqueue(node)
				} else {
					_ = topQueue.Enqueue(minNode)
				}
			}
		}
		// 如果一批不够最小批次，那么认为应该结束了（只计算7天前的数据）

		offset = offset + len(arts) // 否则就计算下一批
	}
	res := make([]domain.Article, 0, topQueue.Len())
	for i := topQueue.Len() - 1; i >= 0; i-- {
		ele, _ := topQueue.Dequeue()
		res[i] = ele.art
	}
	return res, nil
}
