package validator

import (
	"bbs-web/internal/migrator"
	"bbs-web/internal/migrator/events"
	"bbs-web/pkg/logger"
	"bbs-web/pkg/utils/zifo/slice"
	"context"
	s2 "github.com/ecodeclub/ekit/slice"
	"gorm.io/gorm"
	"time"
)

// @Description
// @Author 代码小学生王木木

// Validator T
// @Description: T 泛型约束 必须是实现了migrator.Entity接口的类型
type Validator[T migrator.Entity] struct {
	// 校验  以 xxx 为准
	base   *gorm.DB
	target *gorm.DB

	lg logger.Logger

	p events.Producer

	batchSize int
}

// Validate
//
//	@Description: 调用者可以通过ctx来控制校验程序退出,全量校验就是一行一行的比对，故逐行请求数据库(进阶:可否批量查询)
//	@receiver v
//	@param ctx
func (v *Validator[T]) Validate(ctx context.Context) {
	offset := -1
	for {
		ctx1, cancel := context.WithTimeout(context.Background(), time.Second*5)
		offset++ // 秒啊，进来就更新offset，比较好控制。因为后面有很多的continue和return
		var src T
		//err := v.base.Offset(offset).Model(t).First(&t).Error
		err := v.base.Offset(offset).Order("id").First(&src).Error
		switch err {
		case gorm.ErrRecordNotFound:
			// 比对完毕，没有数据了,全量校验结束，
			return
		case nil:
			// 没有问题，查询到了数据 需从target查询数据，并且校验
			var dst T
			err1 := v.target.Where("id = ?", src.ID()).First(&dst).Error
			cancel()
			switch err1 {
			case nil:
				// 找到了对应的数据，要开始比对了  1. src == dst 直接比较 这是不允许的。
				// 2. 利用反射来比较 reflect.DeepEqual(src,dst)  --> 原则上是可以的
				//if !reflect.DeepEqual(src,dst){
				//
				//}
				// 3. 你来告诉我怎么比
				if !src.CompareTo(dst) {
					v.notify(ctx1, src.ID(), events.InconsistentEventTypeNEQ)
				}
			case gorm.ErrRecordNotFound:
				// 意味着目标库里没有这一行数据
				v.notify(ctx1, src.ID(), events.InconsistentEventTypeTargetMissing)
			default:
				// 放纵做法: 老子不管，我认为没问题
				v.lg.Error("查询数据，查询他target错误", logger.Error(err))
				continue
				// 保守做法，必须管，必须修复，大不了白修复了嘛，错杀一千不放过一个  上报Kafka
				// return
			}
			continue
		default:
			// 未知错误  数据库错误  选择忽略 或者 中断执行
			v.lg.Error("校验数据，查询base错误", logger.Error(err))
			continue
		}
	}
}

// validateTargetToBase
//
//	@Description: 通过target修复base中不存在的数据，==> 原表数据删除了，但是target表尚存  => 删除target中多余的数据
//	@receiver v
//	@param ctx
//	@param id
func (v *Validator[T]) validateTargetToBase(ctx context.Context) {
	offset := -v.batchSize
	for {
		offset = offset + v.batchSize
		dbCtx, cancel := context.WithTimeout(context.Background(), time.Second)
		var srcs []T
		err := v.target.WithContext(ctx).Offset(offset).Limit(v.batchSize).Order("id").Find(&srcs).Error
		cancel()
		if len(srcs) == 0 {
			return
		}
		switch err {
		case gorm.ErrRecordNotFound:
			// 说明结束了 没有数据了
			return
		case nil:
			// 查询源表
			ids := slice.Map(srcs, func(idx int, src T) int64 {
				return src.ID()
			})
			var targets []T
			err = v.base.Where("id IN ?", ids).Find(&targets).Error
			switch err {
			case gorm.ErrRecordNotFound:
				// 全没有
				v.notifyBaseMissing(dbCtx, ids)
			case nil:
				srcIds := slice.Map(targets, func(idx int, src T) int64 {
					return src.ID()
				})
				diff := s2.DiffSet(ids, srcIds) // 计算差集
				v.notifyBaseMissing(dbCtx, diff)
				continue
			default:
				continue
			}

		default:
			// 其他错误
			v.lg.Error("校验数据，查询target错误", logger.Error(err))
			continue
		}
		if len(srcs) < v.batchSize {
			// 没有数据了 要退出
			return
		}
	}
}

func (v *Validator[T]) notifyBaseMissing(ctx context.Context, ids []int64) {
	for _, id := range ids {
		v.notify(ctx, id, events.InconsistentEventTypeSourceMissing)
	}
}

func (v *Validator[T]) notify(ctx context.Context, id int64, ty string) {
	event := events.InconsistentEvent{
		ID:        id,
		Direction: "src",
		Type:      ty,
	}
	err := v.p.ProduceInconsistentEvent(ctx, event)
	if err != nil {
		// 发送kafka失败怎么办？ 重试？记日志？上报监控？写入MYSQL？都可以
		v.lg.Error("写入kafka失败", logger.Error(err))
	}
}
