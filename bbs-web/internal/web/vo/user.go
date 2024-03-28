package vo

type LoginVo struct {
}

type RegisterUserReq struct {
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	RPassword   string `json:"r_password"`
	CaptchaCode string `json:"captcha_code"`
	CaptchaId   string `json:"captcha_id"`
}