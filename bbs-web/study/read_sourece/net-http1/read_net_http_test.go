package net_http1

import (
	"net/http"
	"testing"
)

// @Description
// @Author 代码小学生王木木

func TestNetHttp(t *testing.T) {
	http.HandleFunc("/hello", SayHello) // 注册SayHello函数
	http.HandleFunc("/help", SayHelp)   // 注册SayHelp函数
	http.ListenAndServe(":8023", nil)   // 启动服务
}
func SayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Shit Bro!"))
}
func SayHelp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Help me Bro !"))
}
