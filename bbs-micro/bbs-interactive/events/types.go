package events

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-10 18:45
import "context"

type Consumer interface {
	// ConsumerReadEvent()
	Start() error
}

type Producer interface {
	// 阅读事件(领域事件)
	ProduceReadEvent(ctx context.Context, evt ReadEvent) error
}

const TopicString = "read_article"

type ReadEvent struct {
	Uid int64 // 用户ID
	Aid int64 // 文章id
}

type ReadEventBatch struct {
	Uids []int64 // 用户ID
	Aids []int64 // 文章id
}
