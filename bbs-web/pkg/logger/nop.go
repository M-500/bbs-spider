package logger

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-01 18:00
type NopLogger struct {
}

func NewNoOpLogger() *NopLogger {
	return &NopLogger{}
}
func (n *NopLogger) Debug(msg string, args ...Field) {
}

func (n *NopLogger) Info(msg string, args ...Field) {
}

func (n *NopLogger) Warn(msg string, args ...Field) {
}

func (n *NopLogger) Error(msg string, args ...Field) {
}
