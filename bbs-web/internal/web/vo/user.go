package vo

type LoginVo struct {
}

type RegisterUserReq struct {
	UserName    string `json:"username"`
	Password    string `json:"password"`
	RPassword   string `json:"rpassword"`
	CaptchaCode string `json:"captcha_code"`
	CaptchaId   string `json:"captcha_id"`
}
