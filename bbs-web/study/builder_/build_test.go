package builder_

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-11 11:37

type Config struct {
	Dsn         string
	MaxIdleConn int64
	MaxOpenConn int64
}

func NewConfigBuilder() *Config {

}
