package netx

import "net"

// @Description
// @Author 代码小学生王木木

// GetOutIP
//
//	@Description: 获取本机的对外IP（局域网）
//	@return string
func GetOutIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80") // UDP发送DNS地址 国内可以用 114.114.114.114
	if err != nil {
		return ""
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
