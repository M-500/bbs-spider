package web

import (
	"bbs-web/internal/web/handler"
	"github.com/gin-gonic/gin"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 15:28

type Router struct {
	//userHdl handler.UserHandler
	artHdl *handler.ArticleHandler
}

func NewRouter(artHdl *handler.ArticleHandler) *Router {
	return &Router{
		//userHdl: userHdl,
		artHdl: artHdl,
	}
}

func (r *Router) RegisterURL(engine *gin.Engine) {
	//engine.POST("", r.userHdl.PwdLogin) // 账号密码登录

	articleGroup := engine.Group("/articles")
	{
		articleGroup.POST("/edit", r.artHdl.Edit)
		articleGroup.POST("/withdraw", r.artHdl.Withdraw)
		articleGroup.POST("/publish", r.artHdl.Publish)
		articleGroup.POST("/list", r.artHdl.List)
		articleGroup.GET("/detail/:id", r.artHdl.Detail)
	}

	pub := engine.Group("/pub")
	{
		pub.GET("/:id", r.artHdl.PubDetail)
		pub.POST("/like", r.artHdl.Like)
		pub.POST("/reward", r.artHdl.Reward)
	}
}
