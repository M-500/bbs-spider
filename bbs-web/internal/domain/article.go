package domain

import "time"

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-26 16:00

// Author 在帖子这个领域内，是一个值对象
type Author struct {
	Id       int64
	UserName string
	NickName string
	BirthDay *time.Time
	Avatar   string
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

// Abstract
//
//	@Description: 根据内容获取摘要  如果数据库不存储摘要的话，就要用这个方法
//	@receiver a
//	@return string
func (a Article) Abstract() string {
	cs := []rune(a.Content)
	if len(cs) < 100 {
		return a.Content
	}
	return string(cs[:100])
}

// 使用衍生类型 方便在上面加方法
type ArticleStatus uint8

const (
	// ArticleStatusUnknown 为了避免零值之类的问题
	ArticleStatusUnknown     ArticleStatus = iota
	ArticleStatusUnpublished               // 未发表
	ArticleStatusPublished                 // 已发表
	ArticleStatusPrivate                   // 私有
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

// 状态的另一种方式

type ArticleStatusV1 struct {
	Val  uint8
	Name string
}

var (
	ArticleStatusV1Unknown = ArticleStatusV1{
		Val:  0,
		Name: "unknown",
	}
	//
)
