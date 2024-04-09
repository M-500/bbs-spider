package service

import (
	"bbs-web/internal/domain"
	"bbs-web/internal/service/article"
	"bbs-web/pkg/utils/zifo/slice"
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

	scoreFn func(t time.Time, likeCnt int64) float64 // 要求不能返回负数
}

func NewBatchRankingService(artSvc article.IArticleService, intrSvc InteractiveService) *batchRankingService {
	return &batchRankingService{
		artSvc:    artSvc,
		intrSvc:   intrSvc,
		batchSize: 100,
		queueCap:  100,
		scoreFn: func(t time.Time, likeCnt int64) float64 {
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
	)
	type Score struct {
		art   domain.Article
		score float64
	}
	topQueue := queue.NewConcurrentPriorityQueue[Score](svc.queueCap, func(src Score, dst Score) int {
		if src.score > dst.score {
			return 1
		} else if src.score == dst.score {
			return 0
		}
		return -1
	})

	for {
		now := time.Now()
		// 1. 先拿一批数据
		arts, err := svc.artSvc.ListPub(ctx, now, offset, svc.batchSize)
		if err != nil {
			return nil, err
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
			score := svc.scoreFn(art.Utime, intr.LikeCnt+2) // +2 为了规避负数问题

			// 从小根堆中拿到热度最低的

			// 坑一: 堆中的元素要是不满100个，就有bug，因为不存在插入，一直在替换
			//dequeue, err := topQueue.Dequeue()
			//if err != nil {
			//	// 不管他 妈的！
			//}
			//if dequeue.s core > score {
			//	topQueue.Enqueue(dequeue)
			//} else {
			//	topQueue.Enqueue(Score{
			//		art:   art,
			//		score: score,
			//	})
			//}
			// 已经满了
			if topQueue.Len() == svc.queueCap {
				dequeue, err := topQueue.Dequeue()
				if err != nil {
					// 不管他 妈的！
				}
				if dequeue.score > score {
					topQueue.Enqueue(dequeue)
				} else {
					topQueue.Enqueue(Score{
						art:   art,
						score: score,
					})
				}
			}
			// 没有满，无脑追加
			err := topQueue.Enqueue(Score{
				art:   art,
				score: score,
			})
			if err != nil {
				// 有可能堆满了，也有可能其他错误
			}
		}
		// 最后得出结果
		if len(arts) < svc.batchSize {
			// 如果一批不够最小批次，那么认为应该结束了
			break
		}
		offset = offset + len(arts) // 否则就计算下一批
	}
	res := make([]domain.Article, 0, topQueue.Len())
	for i := topQueue.Len() - 1; i >= 0; i-- {
		ele, _ := topQueue.Dequeue()
		res[i] = ele.art
	}
	return res, nil
}
