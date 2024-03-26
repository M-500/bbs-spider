package handler

import (
	"bbs-web/internal/service"
	"github.com/gin-gonic/gin"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 15:47

type ArticleHandler struct {
	svc service.IArticleService
}

func NewArticleHandler(svc service.IArticleService) *ArticleHandler {
	return &ArticleHandler{
		svc: svc,
	}
}

// Edit
//
//	@Description: 编辑文章
//	@receiver h
//	@param ctx
func (h *ArticleHandler) Edit(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "好好好",
	})
	return
}

// Withdraw
//
//	@Description: 下架
//	@receiver h
//	@param ctx
func (h *ArticleHandler) Withdraw(ctx *gin.Context) {

}

// Publish
//
//	@Description: 发布
//	@receiver h
//	@param ctx
func (h *ArticleHandler) Publish(ctx *gin.Context) {

}

// List
//
//	@Description: 查看列表
//	@receiver h
//	@param ctx
func (h *ArticleHandler) List(ctx *gin.Context) {

}

func (h *ArticleHandler) Detail(ctx *gin.Context) {

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
