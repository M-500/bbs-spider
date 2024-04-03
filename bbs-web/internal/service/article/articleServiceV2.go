package article

import (
	"bbs-web/internal/domain"
	"bbs-web/internal/repository/article"
	"bbs-web/pkg/logger"
	"context"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-02 9:54

type articleServiceV2 struct {
	articleService
	l         logger.Logger
	writeRepo article.ArtWriterRepo
	readRepo  article.ArticleReaderRepository
}

// Publish
//
//	@Description: 只重写一个方法
func (svc *articleServiceV2) Publish(ctx context.Context, art domain.Article) (int64, error) {
	var (
		err error
		id  = art.Id
	)

	// 对于制作库来说
	if art.Id > 0 {
		err = svc.writeRepo.Update(ctx, art)
	} else {
		id, err = svc.writeRepo.Create(ctx, art)
	}
	if err != nil {
		return 0, err
	}
	art.Id = id // 要同步ID
	// 下面开始操作从库了 如何同步？同步出错了怎么办？
	/*
		1. 重试几次(日常开发最常用的就是重试)，重试失败就接入告警系统，然后人工介入(你没办法做到100%的正确，也做不到)
		2. 接入消息队列，走异步，先保存到本地文件
		3. 走Canal 关于canal可以参考 https://zhuanlan.zhihu.com/p/177001630
	*/
	for i := 0; i < 5; i++ {
		id, err = svc.readRepo.Save(ctx, art)
		if err == nil {
			break
		}
		svc.l.Error("同步线上库失败，部分失败", logger.Int64("article_id", id), logger.Error(err))
		time.Sleep(time.Duration(1))
	}
	if err == nil {
		svc.l.Error("同步线上库全部失败", logger.Int64("article_id", id), logger.Error(err))
		return 0, err
	}
	return id, err
}
