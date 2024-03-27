package domain

import "time"

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 16:00

// Author 在帖子这个领域内，是一个值对象
type Author struct {
	Id   int64
	Name string
}
type Article struct {
	Id      int64
	Title   string
	Content string
	// Author 要从用户来
	Author      Author
	Status      ArticleStatus
	Summary     string
	ContentType string
	Cover       string
	Ctime       time.Time
	Utime       time.Time
}

type ArticleStatus uint8

const (
	// ArticleStatusUnknown 为了避免零值之类的问题
	ArticleStatusUnknown ArticleStatus = iota
	ArticleStatusUnpublished
	ArticleStatusPublished
	ArticleStatusPrivate
)

func (s ArticleStatus) ToUint8() uint8 {
	return uint8(s)
}

func (s ArticleStatus) NonPublished() bool {
	return s != ArticleStatusPublished
}

func (s ArticleStatus) String() string {
	switch s {
	case ArticleStatusPrivate:
		return "private"
	case ArticleStatusUnpublished:
		return "unpublished"
	case ArticleStatusPublished:
		return "published"
	default:
		return "unknown"
	}
}
