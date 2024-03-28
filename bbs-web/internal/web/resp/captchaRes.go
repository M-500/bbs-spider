package resp

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-28 14:32

type CaptchaResponse struct {
	CaptchaID string `json:"captcha_id"`
	PicPath   string `json:"pic_path"`
}
