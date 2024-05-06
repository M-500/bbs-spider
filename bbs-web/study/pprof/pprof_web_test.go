package pprof

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"testing"
	"time"
)

func AddSliceDemo() {
	slice := make([]int, 0)
	c := 1
	for i := 0; i < 100000; i++ {
		c = i + 1 + 2 + 3 + 4 + 5
		slice = append(slice, c)
	}
}

func DemoHandler(w http.ResponseWriter, r *http.Request) {
	go func() {
		AddSliceDemo()
	}()
	w.Write([]byte("正在异步调用"))
}
func DemoHandlerV1(w http.ResponseWriter, r *http.Request) {
	Demo()
	for i := 0; i < 100; i++ {
		time.Sleep(time.Millisecond)
		fmt.Println(i)
	}
	w.Write([]byte("正在异步调用"))
}

func Demo() {
	time.Sleep(time.Second)
}

func TestServer(t *testing.T) {
	http.HandleFunc("/demo", DemoHandler)
	http.HandleFunc("/demo1", DemoHandlerV1)
	http.ListenAndServe("localhost:8910", nil)
}
