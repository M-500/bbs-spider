package handler

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 11:45

import (
	"bbs-web/internal/domain"
	"bbs-web/internal/repository/dao"
	"bbs-web/internal/service"
	"bbs-web/internal/web/jwtx"
	"bbs-web/internal/web/vo"
	"bbs-web/pkg/ginplus"
	"errors"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc    service.IUserService
	capSvc service.ICaptchaSvc
	jwtx.JwtHandler
}

func NewUserHandler(svc service.IUserService, acpSvc service.ICaptchaSvc, j jwtx.JwtHandler) *UserHandler {
	return &UserHandler{svc: svc, capSvc: acpSvc, JwtHandler: j}
}

func (h *UserHandler) PwdLogin(ctx *gin.Context, req vo.PwdLoginReq) (ginplus.Result, error) {
	check := h.capSvc.CheckCaptcha(req.CaptchaId, req.CaptchaCode, true)
	if !check {
		return ginplus.Result{
			Code: 501001,
			Msg:  "验证码不正确",
		}, errors.New("验证码不正确")
	}
	user, err := h.svc.Login(ctx, req.UserName, req.Password)
	if err != nil {
		return ginplus.Result{
			Code: 501002,
			Msg:  err.Error(),
		}, err
	}
	token, err := h.GetJWTToken(ctx, user.Id)
	if err != nil {
		return ginplus.Result{
			Code: 501003,
			Msg:  err.Error(),
		}, err
	}
	return ginplus.Result{
		Data: "Bearer " + token,
	}, nil
}

// Register
//
//	@Description: 用户注册
func (h *UserHandler) Register(ctx *gin.Context, req vo.RegisterUserReq) (ginplus.Result, error) { // Register
	//  @Description:
	//  @receiver h
	//  @param ctx
	//  @param req
	//  @return ginplus.Result
	//  @return error
	// 调用server方法
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
		}, errors.New("验证码不正确")
	}
	err := h.svc.SignUp(ctx, domain.UserInfo{
		Id:       0,
		UserName: req.UserName,
		Password: req.Password,
	})

	if err == dao.ErrUserDuplicate {
		return ginplus.Result{
			Code: 501000,
			Msg:  err.Error(),
		}, err
	}

	// 校验两次密码
	if err != nil {
		return ginplus.Result{
			Code: 501005,
			Msg:  "注册失败，系统异常",
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
