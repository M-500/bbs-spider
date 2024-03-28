package handler

import (
	"bbs-web/internal/service"
	"bbs-web/internal/web/resp"
	"bbs-web/pkg/ginplus"
	"github.com/gin-gonic/gin"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-28 14:25

type CaptchaHandler struct {
	svc service.ICaptchaSvc
}

func NewCaptchaHandler(svc1 service.ICaptchaSvc) *CaptchaHandler {
	return &CaptchaHandler{
		svc: svc1,
	}
}

func (h *CaptchaHandler) ImageCaptcha(ctx *gin.Context) (ginplus.Result, error) {
	id, path, err := h.svc.MakeCaptcha()
	if err != nil {
		return ginplus.Result{
			Msg: "系统错误",
		}, err
	}
	return ginplus.Result{
		Data: resp.CaptchaResponse{
			CaptchaID: id,
			PicPath:   path,
		},
	}, nil
}
