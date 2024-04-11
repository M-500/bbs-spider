package saramax

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-11 12:01

type Consumer interface {
	// ConsumerReadEvent()
	Start() error
}
