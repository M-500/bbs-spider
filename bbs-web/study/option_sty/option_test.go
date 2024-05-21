package option_sty

// @Description
// @Author 代码小学生王木木

type MyOption func(opt *DBConfig)

type DBConfig struct {
	Database string // 数据库名
	Host     string // IP地址
	Port     int    // 端口
	UserName string // 用户名
	Pwd      string // 密码
	CharSet  string // 字符集
	MaxConn  int    // 连接池大小
}

func NewDBConfig(database string, host string, port int, userName string, pwd string, charSet string, maxConn int) *DBConfig {
	return &DBConfig{
		Database: database,
		Host:     host,
		Port:     port,
		UserName: userName,
		Pwd:      pwd,
		CharSet:  charSet,
		MaxConn:  maxConn}
}
