package handler

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 11:45

import "github.com/gin-gonic/gin"

type UserHandler struct {
}

func (h *UserHandler) PwdLogin(ctx *gin.Context) {
	type Req struct {
		UserName    string `json:"user_name"`
		Password    string `json:"password"`
		CaptchaCode string `json:"captcha_code"`
		CaptchaId   string `json:"captcha_id"`
	}
}

func (h *UserHandler) Register(ctx *gin.Context) {
	type Req struct {
		UserName    string `json:"user_name"`
		Password    string `json:"password"`
		RPassword   string `json:"r_password"`
		CaptchaCode string `json:"captcha_code"`
		CaptchaId   string `json:"captcha_id"`
	}
}

// Create
//
//	@Description: 创建管理员用户
//	@receiver h
//	@param ctx
func (h *UserHandler) Create(ctx *gin.Context) {

}

// GetUserInfo
//
//	@Description: 获取用户资料
//	@receiver h
//	@param ctx
func (h *UserHandler) GetUserInfo(ctx *gin.Context) {

}
