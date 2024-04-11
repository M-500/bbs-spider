package jwtx

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-03 11:34

type JwtHandler interface {

	// ExtractToken
	//  @Description: 提取token
	ExtractToken(ctx *gin.Context) string

	// GetJWTToken
	//  @Description: 设置JWT token
	GetJWTToken(ctx *gin.Context, uid int64) (string, error)

	// ParseToken
	//  @Description: 解析token
	ParseToken(ctx *gin.Context, tokenString string) (UserClaims, error)
}

type UserClaims struct {
	jwt.RegisteredClaims
	// 声明你自己的要放进去 token 里面的数据
	Id        int64
	UserAgent string
}
