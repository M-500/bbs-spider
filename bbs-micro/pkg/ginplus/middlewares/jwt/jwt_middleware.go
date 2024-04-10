package jwt

import (
	"bbs-web/internal/web/jwtx"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-29 14:25

type LoginJWTMiddlewareBuilder struct {
	paths []string // 不需要JWT拦截的路由
	jwtx.JwtHandler
}

func NewLoginJWTMiddlewareBuilder(j jwtx.JwtHandler) *LoginJWTMiddlewareBuilder {
	return &LoginJWTMiddlewareBuilder{
		JwtHandler: j,
	}
}

func (l *LoginJWTMiddlewareBuilder) IgnorePaths(path string) *LoginJWTMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}

func (l *LoginJWTMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 不需要登录校验的
		for _, path := range l.paths {
			if c.Request.URL.Path == path {
				return
			}
		}
		tokenStr := l.ExtractToken(c)
		claims, err := l.ParseToken(c, tokenStr)
		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
		}
		c.Set("users", claims)
		c.Next()
	}
}
