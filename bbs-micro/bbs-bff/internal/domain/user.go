package domain

import "time"

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 15:13

type UserInfo struct {
	Id       int64
	UserName string
	NickName string
	Password string
	BirthDay *time.Time
	Avatar   string
	// 文章数
	// 粉丝数
	// 关注数量
	// 评论数量
}
