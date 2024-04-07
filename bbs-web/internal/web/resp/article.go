package resp

import "time"

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-07 11:12

type ArticleResp struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	AuthorId    int64     `json:"authorId"`
	AuthorName  string    `json:"authorName"`
	Status      string    `json:"status"` // 状态这个东西，前端处理 后端处理都可以
	Summary     string    `json:"summary"`
	Content     string    `json:"content"`
	ContentType string    `json:"contentType"`
	Cover       string    `json:"cover"`
	Ctime       time.Time `json:"ctime"`
	Utime       time.Time `json:"utime"`
}
