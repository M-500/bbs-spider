package domain

import "time"

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-12 11:38

type Collect struct {
	UserId      int64     `json:"user_id"`
	CName       string    `json:"c_name"`
	Description string    `json:"description"`
	Sort        int64     `json:"sort"`
	ResourceNum int64     `json:"resource_num"`
	IsPub       bool      `json:"is_pub"`
	CommentNum  int64     `json:"comment_num"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
}
