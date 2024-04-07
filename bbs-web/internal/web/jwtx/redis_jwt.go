package jwtx

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-03 12:10

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

var (
	AtKey = []byte("95osj3fUD7fo0mlYdDbncXz4VD2igvf0")
	RtKey = []byte("95osj3fUD7fo0mlYdDbncXz4VD2igvfx")
)

var (
	NotLogin    = errors.New("用户未登录")
	TokenExpire = errors.New("Token已过期")
)

type RedisJwtHandler struct {
	cmd redis.Cmdable
}

func (r *RedisJwtHandler) queryFromHeader(ctx *gin.Context) string {
	tokenHeader := ctx.GetHeader("Authorization")
	segs := strings.Split(tokenHeader, " ")
	if len(segs) != 2 {
		return ""
	}
	return segs[1]
}

func (r *RedisJwtHandler) queryFromURL(ctx *gin.Context) string {
	// url=/xx/xx?token=xxxx
	return ctx.Query("token")
}

func (r *RedisJwtHandler) queryFromFormData(ctx *gin.Context) string {
	return ctx.PostForm("token")
}

func (r *RedisJwtHandler) getKey(uid int64) string {
	return fmt.Sprintf("login_hold:%d", uid)
}

func (r *RedisJwtHandler) ExtractToken(ctx *gin.Context) string {
	var tkStr string
	// 从header中获取
	if tkStr = r.queryFromHeader(ctx); len(tkStr) != 0 {
		return tkStr
	}
	if tkStr = r.queryFromURL(ctx); len(tkStr) != 0 {
		return tkStr
	}
	if tkStr = r.queryFromFormData(ctx); len(tkStr) != 0 {
		return tkStr
	}
	return tkStr
}

func (r *RedisJwtHandler) GetJWTToken(ctx *gin.Context, uid int64) (string, error) {
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		},
		Id: uid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, err := token.SignedString(AtKey)
	if err != nil {
		return "", err
	}
	// TODO 是否需要将当前用户的token塞进redis中？
	r.cmd.Set(ctx, r.getKey(uid), tokenStr, time.Hour*24)
	return tokenStr, nil
}

func (r *RedisJwtHandler) ParseToken(ctx *gin.Context, tokenStr string) (UserClaims, error) {
	claims := UserClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(AtKey), nil
	})
	if err != nil {
		return UserClaims{}, NotLogin
	}
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return UserClaims{}, TokenExpire
	}
	if claims.UserAgent != ctx.Request.UserAgent() {

	}
	if token == nil || !token.Valid || claims.Id <= 0 {
		return UserClaims{}, NotLogin
	}
	// TODO 是否需要校验Redis的Token是否过期？
	return claims, nil
}

func NewRedisJWTHandler(cmd redis.Cmdable) JwtHandler {
	return &RedisJwtHandler{
		cmd: cmd,
	}
}
