package migrator

// @Description
// @Author 代码小学生王木木

type Entity interface {
	ID() int64

	// 定义比较方法
	CompareTo(t Entity) bool
}
