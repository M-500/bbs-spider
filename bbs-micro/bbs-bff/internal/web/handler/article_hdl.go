package handler

import (
	intrv1 "bbs-micro/api/proto/gen/proto/intr/v1"
	"bbs-micro/bbs-bff/internal/domain"
	"bbs-micro/bbs-bff/internal/service/article"
	"bbs-micro/bbs-bff/internal/web/jwtx"
	"bbs-micro/bbs-bff/internal/web/resp"
	"bbs-micro/bbs-bff/internal/web/vo"
	ginplus "bbs-micro/pkg/ginplus"
	"bbs-micro/pkg/logger"
	"bbs-micro/pkg/utils/zifo/slice"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"strconv"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 15:47

type ArticleHandler struct {
	svc article.IArticleService
	log logger.Logger
	//interSvc service.InteractiveService // 调用本地方法
	interSvc intrv1.InteractiveServiceServer // 改为RPC调用
	biz      string                          // 业务ID
}

func NewArticleHandler(svc article.IArticleService,
	intrSvc intrv1.InteractiveServiceServer,
	l logger.Logger) *ArticleHandler {
	return &ArticleHandler{
		svc:      svc,
		log:      l,
		interSvc: intrSvc,
		biz:      "article",
	}
}

func (h *ArticleHandler) PubAuthorArtList(ctx *gin.Context) (ginplus.Result, error) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return ginplus.Result{
			Msg: "参数错误",
		}, err
	}
	var (
		eg   errgroup.Group
		arts []domain.Article
	)

	eg.Go(func() error {
		var err error
		arts, err = h.svc.List(ctx, id, 1, 15)
		return err
	})

	er := eg.Wait()
	if er != nil {
		return ginplus.Result{
			Msg: "系统错误",
		}, err
	}
	// 组装内容
	return ginplus.Result{Data: arts}, nil
}

// Edit
//
//	@Description: 编辑文章
//	@receiver h
//	@param ctx
func (h *ArticleHandler) Edit(ctx *gin.Context, req vo.ArticleReq, c jwtx.UserClaims) (ginplus.Result, error) {
	id, err := h.svc.Save(ctx.Request.Context(), req.ToDomain(c.Id))
	if err != nil {
		h.log.Error("编辑文章出错", logger.Error(err))
		return ginplus.Result{
			Code: 510002,
			Msg:  "系统异常",
		}, err
	}
	return ginplus.Result{
		Data: id,
	}, nil
}

// Withdraw
//
//	@Description: 下架 撤回文章
//	@receiver h
//	@param ctx
func (h *ArticleHandler) Withdraw(ctx *gin.Context, user jwtx.UserClaims) (ginplus.Result, error) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return ginplus.Result{
			Msg: "参数错误",
		}, err
	}
	//user.Id
	err = h.svc.Withdraw(ctx, domain.Article{
		Id: i,
		Author: domain.Author{
			Id: user.Id,
		},
	})
	if err != nil {
		return ginplus.Result{}, err
	}
	return ginplus.Result{}, nil
}

// Publish
//
//	@Description: 发布
//	@receiver h
//	@param ctx
func (h *ArticleHandler) Publish(ctx *gin.Context, req vo.ArticleReq, user jwtx.UserClaims) (ginplus.Result, error) {
	publish, err := h.svc.Publish(ctx, req.ToDomain(user.Id))
	if err != nil {
		return ginplus.Result{
			Code: 510003,
			Msg:  "保存帖子失败",
		}, err
	}
	return ginplus.Result{
		Data: publish,
	}, nil
}

// List
//
//	@Description: 查看列表
//	@receiver h
//	@param ctx
func (h *ArticleHandler) List(ctx *gin.Context, req vo.ArticleListReq, user jwtx.UserClaims) (ginplus.Result, error) {
	list, err := h.svc.List(ctx, user.Id, req.PageNum, req.PageSize)
	if err != nil {
		return ginplus.Result{
			Code: 502001,
			Msg:  "",
		}, err
	}
	return ginplus.Result{
		Data: slice.Map[domain.Article, resp.ArticleResp](list, func(idx int, src domain.Article) resp.ArticleResp {
			return resp.ArticleResp{
				Id:          src.Id,
				Title:       src.Title,
				AuthorId:    src.Author.Id,
				AuthorName:  src.Author.UserName,
				Status:      src.Status.String(),
				Summary:     src.Summary,
				ContentType: src.ContentType,
				Cover:       src.Cover,
				Ctime:       src.Ctime,
				Utime:       src.Utime,
			}
		})}, nil
}

func (h *ArticleHandler) Detail(ctx *gin.Context, user jwtx.UserClaims) (ginplus.Result, error) {
	artIdStr := ctx.Param("id") // 从URL中获取id
	artId, err := strconv.ParseInt(artIdStr, 10, 64)
	if err != nil {
		return ginplus.Result{
			Code: 502002,
			Msg:  "参数错误",
		}, nil
	}
	art, err := h.svc.GetById(ctx, artId)
	if err != nil {
		return ginplus.Result{
			Code: 502005,
			Msg:  "系统错误",
		}, err
	}
	if art.Id != user.Id {
		// 说明非法访问 ，需要做反爬,或者上报风控系统
		return ginplus.Result{
			Code: 5002003,
			Msg:  "输入错误",
		}, fmt.Errorf("非法访问文章，创作者 ID 不匹配 %d", user.Id)
	}
	return ginplus.Result{
		Data: resp.ArticleResp{
			Id:          art.Id,
			Title:       art.Title,
			AuthorId:    art.Author.Id,
			AuthorName:  art.Author.UserName,
			Status:      art.Status.String(),
			Summary:     art.Summary,
			Content:     art.Content,
			ContentType: art.ContentType,
			Cover:       art.Cover,
			Ctime:       art.Ctime,
			Utime:       art.Utime,
		},
	}, nil
}

// Like
//
//	@Description: 点赞和取消点赞
//	@receiver h
//	@param ctx
func (h *ArticleHandler) Like(ctx *gin.Context, req vo.LikeReq, c jwtx.UserClaims) (ginplus.Result, error) {
	var err error
	if req.Like {
		_, err = h.interSvc.Like(ctx, &intrv1.LikeRequest{
			Biz:   h.biz,
			BizId: req.Id,
			Uid:   c.Id,
		})
	} else {
		_, err = h.interSvc.CancelLike(ctx, &intrv1.CancelLikeRequest{
			Biz:   h.biz,
			BizId: req.Id,
			Uid:   c.Id,
		})
	}
	if err != nil {
		return ginplus.Result{
			Code: 502005,
			Msg:  "系统错误",
		}, err
	}
	return ginplus.Result{Msg: "OK!"}, nil
}

func (h *ArticleHandler) Collect(ctx *gin.Context, req vo.CollectReq, c jwtx.UserClaims) (ginplus.Result, error) {
	//var err error
	//if req.Collect {
	//	err = h.interSvc.CollectArt(ctx)
	//} else {
	//
	//}
	return ginplus.Result{}, nil
}

// PubDetail
//
//	@Description: 阅读
//	@receiver h
//	@param ctx
func (h *ArticleHandler) PubDetail(ctx *gin.Context, user jwtx.UserClaims) (ginplus.Result, error) {
	idstr := ctx.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		return ginplus.Result{
			Code: 502004,
			Msg:  "参数错误",
		}, err
	}
	var eg errgroup.Group
	var article domain.Article
	var intr *intrv1.GetResponse

	// 这里异步获取文章的点赞数 收藏数 评论数等信息
	eg.Go(func() error {
		article, err = h.svc.GetPublishedById(ctx, id, user.Id)
		return err
	})
	eg.Go(func() error {
		intr, err = h.interSvc.Get(ctx, &intrv1.GetRequest{
			Biz:   h.biz,
			BizId: id,
			Uid:   user.Id,
		})
		return err
	})
	err = eg.Wait() // 任何一个地方返回的err 都会被捕捉到
	if err != nil {
		//
		return ginplus.Result{
			Code: 502005,
			Msg:  "系统错误",
		}, err
	}
	//// 异步增加阅读计数
	//go func() {
	//	// 阅读数+1  最好集成kafka来异步处理，减轻MySQL的压力 因为这里会回写MySQL以新增阅读量(下沉到了service层实现)
	//	// 1. 如果你想摆脱原本主链路的超时控制，你就创建一个新的
	//	// 2. 如果你不想，你就用 ctx
	//	err1 := h.interSvc.IncrReadCnt(ctx, h.biz, article.Id)
	//	if err1 != nil {
	//		h.log.Error("增加文章阅读数失败", logger.Error(err1), logger.Int64("Article_ID", article.Id))
	//	}
	//}()

	return ginplus.Result{Data: resp.ArticleResp{
		Id:          article.Id,
		Title:       article.Title,
		AuthorId:    article.Author.Id,
		AuthorName:  article.Author.UserName,
		Status:      article.Status.String(),
		Summary:     article.Summary,
		Content:     article.Content,
		LikeCnt:     intr.GetIntr().LikeCnt,
		ReadCnt:     intr.GetIntr().ReadCnt,
		CommentCnt:  intr.GetIntr().CommentCnt,
		CollectCnt:  intr.GetIntr().CollectCnt,
		ContentType: article.ContentType,
		Cover:       article.Cover,
		Ctime:       article.Ctime,
		Utime:       article.Utime,
	}}, nil
}

// Reward
//
//	@Description: 打赏
//	@receiver h
//	@param ctx
func (h *ArticleHandler) Reward(ctx *gin.Context) {

}
