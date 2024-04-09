package article

import "context"

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-09 11:10

type Producer interface {
	// 阅读事件(领域事件)
	ProduceReadEvent(ctx context.Context, evt ReadEvent) error
}

type ReadEvent struct {
	Uid int64 // 用户ID
	Aid int64 // 文章id
}
