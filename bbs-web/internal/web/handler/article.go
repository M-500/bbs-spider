package handler

import (
	"bbs-web/internal/service"
	"bbs-web/internal/web/vo"
	"fmt"
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
	var req vo.ArticleReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		ctx.JSON(401, gin.H{
			"msg": "用户输入的数据有问题",
		})
		return
	}
	// 获取用户
	//get := ctx.MustGet(constant.JWT_USET_Key)
	//claims, ok := c.(ijwt.UserClaims) 做类型断言
	id, err := h.svc.Save(ctx.Request.Context(), req.ToDomain(1))
	if err != nil {
		return
	}
	ctx.JSON(200, gin.H{
		"msg":  "好好好",
		"data": id,
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
