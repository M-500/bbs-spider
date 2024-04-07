package web

import (
	"bbs-web/internal/web/handler"
	"bbs-web/internal/web/jwtx"
	"bbs-web/internal/web/vo"
	"bbs-web/pkg/ginplus"
	"github.com/gin-gonic/gin"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 15:28

type Router struct {
	userHdl *handler.UserHandler
	artHdl  *handler.ArticleHandler
	codeHdl *handler.CaptchaHandler
}

func NewRouter(artHdl *handler.ArticleHandler, codeHdl *handler.CaptchaHandler, userHdl *handler.UserHandler) *Router {
	return &Router{
		//userHdl: userHdl,
		artHdl:  artHdl,
		codeHdl: codeHdl,
		userHdl: userHdl,
	}
}

func (r *Router) RegisterURL(engine *gin.Engine) {
	engine.POST("/sign-up", ginplus.WrapJson[vo.RegisterUserReq](r.userHdl.Register)) // 注册用户
	engine.POST("/pwd-login", ginplus.WrapJson[vo.PwdLoginReq](r.userHdl.PwdLogin))   // 账号密码登录

	engine.GET("/code", ginplus.Wrap(r.codeHdl.ImageCaptcha))

	articleGroup := engine.Group("/articles")
	{
		articleGroup.POST("/edit", ginplus.WrapJson[vo.ArticleReq](r.artHdl.Edit))
		articleGroup.POST("/:id/withdraw", ginplus.WrapToken[jwtx.UserClaims](r.artHdl.Withdraw))
		articleGroup.POST("/publish", ginplus.WrapJson[vo.ArticleReq](r.artHdl.Publish))
		articleGroup.POST("/list", ginplus.WrapBodyAndToken[vo.ArticleListReq, jwtx.UserClaims](r.artHdl.List))
		articleGroup.GET("/detail/:id", r.artHdl.Detail)
	}

	pub := engine.Group("/pub")
	{
		pub.GET("/:id", r.artHdl.PubDetail)
		pub.POST("/like", r.artHdl.Like)
		pub.POST("/reward", r.artHdl.Reward)
	}
}
