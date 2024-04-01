package handler

import (
	"bbs-web/internal/service"
	"bbs-web/internal/web/vo"
	"bbs-web/pkg/ginplus"
	"bbs-web/pkg/logger"
	"github.com/gin-gonic/gin"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 15:47

type ArticleHandler struct {
	svc service.IArticleService
	log logger.Logger
}

func NewArticleHandler(svc service.IArticleService, l logger.Logger) *ArticleHandler {
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
func (h *ArticleHandler) Edit(ctx *gin.Context, req vo.ArticleReq) (ginplus.Result, error) {

	// 获取用户
	//get := ctx.MustGet(constant.JWT_USET_Key)
	//claims, ok := c.(ijwt.UserClaims) 做类型断言
	// 超时控制
	id, err := h.svc.Save(ctx.Request.Context(), req.ToDomain(1))
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
