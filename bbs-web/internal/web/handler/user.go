package handler

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 11:45

import (
	"bbs-web/internal/domain"
	"bbs-web/internal/service"
	"bbs-web/internal/web/vo"
	"bbs-web/pkg/ginplus"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc    service.IUserService
	capSvc service.ICaptchaSvc
}

func NewUserHandler(svc service.IUserService, acpSvc service.ICaptchaSvc) *UserHandler {
	return &UserHandler{svc: svc, capSvc: acpSvc}
}

func (h *UserHandler) PwdLogin(ctx *gin.Context) {
	type Req struct {
		UserName    string `json:"user_name"`
		Password    string `json:"password"`
		CaptchaCode string `json:"captcha_code"`
		CaptchaId   string `json:"captcha_id"`
	}
}

func (h *UserHandler) Register(ctx *gin.Context, req vo.RegisterUserReq) (ginplus.Result, error) {
	// 调用server方法
	fmt.Println(req.RPassword, req.Password, req.UserName)
	if req.Password != req.RPassword {
		return ginplus.Result{
			Msg: "两次输入密码不一致",
		}, errors.New("用户两次密码不相同")
	}
	check := h.capSvc.CheckCaptcha(req.CaptchaId, req.CaptchaCode, true)
	if !check {
		return ginplus.Result{
			Code: 501001,
			Msg:  "验证码不正确",
		}, errors.New("用户两次密码不相同")
	}
	err := h.svc.SignUp(ctx, domain.UserInfo{
		Id:       0,
		UserName: req.UserName,
		NickName: "",
		Password: req.Password,
	})
	// 校验两次密码
	if err != nil {
		return ginplus.Result{
			Msg: "注册失败，系统异常",
		}, err
	}
	return ginplus.Result{
		Msg: "注册成功",
	}, nil
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
