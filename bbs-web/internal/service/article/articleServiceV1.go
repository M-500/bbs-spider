package article

import (
	event "bbs-web/internal/events/article"
	"bbs-web/internal/repository/article"
	"bbs-web/pkg/logger"
	"context"
	"time"

	"bbs-web/internal/domain"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 15:59

type articleService struct {
	repo      article.ArticleRepository
	l         logger.Logger
	writeRepo article.ArtWriterRepo
	readRepo  article.ArticleReaderRepository
	//userSvc   service.IUserService

	producer event.Producer
}

func NewArticleService(repo article.ArticleRepository, l logger.Logger, writeRepo article.ArtWriterRepo, readRepo article.ArticleReaderRepository) IArticleService {
	return &articleService{repo: repo, l: l, writeRepo: writeRepo, readRepo: readRepo}
}

func NewArticleServiceV1(writeRepo article.ArtWriterRepo, readRepo article.ArticleReaderRepository, l logger.Logger) IArticleService {
	return &articleService{
		l: l, writeRepo: writeRepo, readRepo: readRepo}
}

func (svc *articleService) Save(ctx context.Context, art domain.Article) (int64, error) {
	art.Status = domain.ArticleStatusUnpublished // service层控制状态
	if art.Id > 0 {
		err := svc.repo.Update(ctx, art)
		return art.Id, err
	}
	return svc.repo.Create(ctx, art)
}

func (svc *articleService) Withdraw(ctx context.Context, art domain.Article) error {
	art.Status = domain.ArticleStatusPrivate // 把状态设置为私有
	return svc.repo.SyncStatus(ctx, art.Id, art.Author.Id, art.Status)
}

// Publish
//
//	@Description: Service层不处理如何同步数据，交给Repo层去处理
func (svc *articleService) Publish(ctx context.Context, art domain.Article) (int64, error) {
	art.Status = domain.ArticleStatusPublished // service层控制状态
	return svc.repo.Sync(ctx, art)
}

// PublishV1
//
//	@Description: Service来处理协调如何同步数据，具体依赖两个Repo来实现同步
func (svc *articleService) PublishV1(ctx context.Context, art domain.Article) (int64, error) {
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
	if err != nil {
		svc.l.Error("同步线上库全部失败", logger.Int64("article_id", id), logger.Error(err))
		return 0, err
	}
	return id, err
}
func (svc *articleService) List(ctx context.Context, uid int64, offset int, limit int) ([]domain.Article, error) {
	return svc.repo.List(ctx, uid, offset, limit)
}

func (svc *articleService) ListPub(ctx context.Context, start time.Time, offset, limit int) ([]domain.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (svc *articleService) GetById(ctx context.Context, id int64) (domain.Article, error) {
	return svc.repo.GetByID(ctx, id)
}

func (svc *articleService) GetPublishedById(ctx context.Context, id, uid int64) (domain.Article, error) {
	art, err := svc.repo.GetPublishedById(ctx, id, uid)
	if err == nil {
		go func() {
			err2 := svc.producer.ProduceReadEvent(ctx, event.ReadEvent{
				// 就算kafka消费者要用article的其他字段，让他通过id去查，你不要带过去，因为你带过去的数据，在他使用过程中可能会被修改！
				Uid: uid,
				Aid: art.Id,
			})
			if err2 != nil {
				svc.l.Error("发送读者阅读事件失败", logger.Error(err),
					logger.Int64("articleId", art.Id),
					logger.Int64("userId", uid))
			}
		}()
	}
	return art, err
	//// 组装User
	//user, err := svc.userSvc.FindById(ctx, uid)
	//if err != nil {
	//	return domain.Article{}, err
	//}
	//return domain.Article{
	//	Id:      art.Id,
	//	Title:   art.Title,
	//	Content: art.Content,
	//	Author: domain.Author{
	//		Id:       user.Id,
	//		UserName: user.UserName,
	//		NickName: user.NickName,
	//		BirthDay: user.BirthDay,
	//		Avatar:   user.Avatar,
	//	},
	//	Status:      art.Status,
	//	Summary:     art.Summary,
	//	ContentType: art.ContentType,
	//	Cover:       art.Cover,
	//	Ctime:       art.Ctime,
	//	Utime:       art.Utime,
	//}, nil

}

func (svc *articleService) GetByIds(ctx context.Context, biz string, ids []int64) ([]domain.Article, error) {
	//TODO implement me
	panic("implement me")
}

//func (svc *articleService) ListPubArtsByAuthId(ctx context.Context, uid int64) ([]domain.Article, error) {
//	//TODO implement me
//	panic("implement me")
//}
