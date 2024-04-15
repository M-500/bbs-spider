package job

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-15 15:43

type Job interface {
	Name() string
	Run() error
}
