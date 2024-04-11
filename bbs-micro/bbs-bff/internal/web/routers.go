package web

import (
	"bbs-micro/bbs-bff/internal/web/handler"
	"bbs-micro/bbs-bff/internal/web/jwtx"
	"bbs-micro/bbs-bff/internal/web/vo"
	gp "bbs-micro/pkg/ginplus"
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
	engine.POST("/sign-up", gp.WrapJson[vo.RegisterUserReq](r.userHdl.Register)) // 注册用户
	engine.POST("/pwd-login", gp.WrapJson[vo.PwdLoginReq](r.userHdl.PwdLogin))   // 账号密码登录

	engine.GET("/code", gp.Wrap(r.codeHdl.ImageCaptcha)) // 获取图片验证码

	articleGroup := engine.Group("/articles") // 文章创作者关注的接口
	{
		articleGroup.POST("/edit", gp.WrapBodyAndToken[vo.ArticleReq, jwtx.UserClaims](r.artHdl.Edit))       // 新建文章
		articleGroup.POST("/:id/withdraw", gp.WrapToken[jwtx.UserClaims](r.artHdl.Withdraw))                 // 下架某一篇文章
		articleGroup.POST("/publish", gp.WrapBodyAndToken[vo.ArticleReq, jwtx.UserClaims](r.artHdl.Publish)) // 发布某一篇文章
		articleGroup.POST("/list", gp.WrapBodyAndToken[vo.ArticleListReq, jwtx.UserClaims](r.artHdl.List))   // 创作者查看自己的文章列表
		articleGroup.GET("/detail/:id", gp.WrapToken[jwtx.UserClaims](r.artHdl.Detail))                      // 作者查看文章详情
	}

	pub := engine.Group("/pub") // 读者关注的接口
	{
		pub.GET("/list/:id", gp.Wrap(r.artHdl.PubAuthorArtList))                                    //  用户查看某个作者的所有文章
		pub.GET("/:id", gp.WrapToken[jwtx.UserClaims](r.artHdl.PubDetail))                          // 读取文章详情
		pub.POST("/like", gp.WrapBodyAndToken[vo.LikeReq, jwtx.UserClaims](r.artHdl.Like))          // 点赞/取消点赞某一篇文章
		pub.POST("/collect", gp.WrapBodyAndToken[vo.CollectReq, jwtx.UserClaims](r.artHdl.Collect)) // 收藏/取消收藏 某一篇文章
		pub.POST("/reward", r.artHdl.Reward)
	}
}
