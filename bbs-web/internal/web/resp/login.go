package resp

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-07 12:56

type PWDLoginResp struct {
	Token    string `json:"token"`
	UserName string `json:"userName"`
	UserId   int64  `json:"userId"`
	Cover    string `json:"cover"`
}
