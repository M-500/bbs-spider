package service

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-28 14:27

import (
	"errors"
	"github.com/mojocn/base64Captcha"
)

type ICaptchaSvc interface {
	MakeCaptcha() (string, string, error)
	CheckCaptcha(id, code string, clear bool) bool
}

type CaptchaService struct {
	store base64Captcha.Store
}

func NewCaptchaService() ICaptchaSvc {
	return &CaptchaService{
		store: base64Captcha.DefaultMemStore,
	}
}

func (c *CaptchaService) MakeCaptcha() (string, string, error) {
	driver := base64Captcha.NewDriverDigit(
		80,  // 高度 png 像素高度
		240, // 宽度 png 像素高度
		5,   // 验证码默认位数
		0.7, //单个数字的最大绝对倾斜因子
		180) // 背景圆圈的数量
	cp := base64Captcha.NewCaptcha(driver, c.store)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		return "", "", errors.New("生成验证码失败")
	}

	return id, b64s, nil
}

func (c *CaptchaService) CheckCaptcha(id, code string, clear bool) bool {
	if c.store.Verify(id, code, clear) {
		return true
	}
	return false
}
