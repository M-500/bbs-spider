package handler

import (
	"bbs-web/internal/domain"
	"bbs-web/internal/service/article"
	"bbs-web/internal/web/jwtx"
	"bbs-web/internal/web/resp"
	"bbs-web/internal/web/vo"
	ginplus "bbs-web/pkg/ginplus"
	"bbs-web/pkg/logger"
	"bbs-web/pkg/utils/zifo/slice"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 15:47

type ArticleHandler struct {
	svc article.IArticleService
	log logger.Logger
}

func NewArticleHandler(svc article.IArticleService, l logger.Logger) *ArticleHandler {
	return &ArticleHandler{
		svc: svc,
		log: l,
	}
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
func (h *ArticleHandler) Publish(ctx *gin.Context, req vo.ArticleReq) (ginplus.Result, error) {
	publish, err := h.svc.Publish(ctx, req.ToDomain(1))
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
//	@Description: 点赞
//	@receiver h
//	@param ctx
func (h *ArticleHandler) Like(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "13e1",
	})
}

// PubDetail
//
//	@Description: 阅读
//	@receiver h
//	@param ctx
func (h *ArticleHandler) PubDetail(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "13e1",
	})
}

// Reward
//
//	@Description: 打赏
//	@receiver h
//	@param ctx
func (h *ArticleHandler) Reward(ctx *gin.Context) {

}
