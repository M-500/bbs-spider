package events

import "context"

// @Description
// @Author 代码小学生王木木

type Producer interface {
	ProduceInconsistentEvent(ctx context.Context, evt InconsistentEvent) error
}

type InconsistentEvent struct {
	ID int64
	// 用什么来修？ 取值为SRC意味着以源表为准，DST以目标表为准
	Direction string
	// 另外一个问题 有些时候，一些观测或者一些第三方，需要知道是什么引擎了不一致，我们要记录一下备注
	Type string // 可选
}

const (
	InconsistentEventTypeNEQ           = "neq"
	InconsistentEventTypeTargetMissing = "target_missing"
	InconsistentEventTypeSourceMissing = "source_missing"
)
