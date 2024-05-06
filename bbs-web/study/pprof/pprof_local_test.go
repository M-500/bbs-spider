package pprof

import (
	"log"
	"os"
	"runtime/pprof"
	"sync"
	"testing"
)

func Feb(x int) int {
	if x <= 0 {
		return 0
	}
	if x < 2 {
		return 1
	}
	return Feb(x-1) + Feb(x-2)
}

func AddSlice() {
	slice := make([]int, 0)
	c := 1
	for i := 0; i < 100000; i++ {
		c = i + 1 + 2 + 3 + 4 + 5
		slice = append(slice, c)
	}
}

func TestFeb(t *testing.T) {
	//采样cpu运行状态
	f, err := os.Create("cpu.pprof")
	if err != nil {
		log.Fatal("无法创建 CPU profile: ", err)
	}
	err = pprof.StartCPUProfile(f)
	if err != nil {
		log.Fatal("采集CPU信息失败: ", err)
	} // 采样CPU
	defer pprof.StopCPUProfile()

	f2, err2 := os.Create("mem.pprof")
	if err2 != nil {
		log.Fatal("无法创建 Memory profile: ", err)
	}
	err = pprof.WriteHeapProfile(f2)
	if err != nil {
		log.Fatal("采集内存信息失败: ", err)
	} // 采样内存
	//for i := 0; i < 50; i++ {
	//	AddSlice()
	//}
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		go func(i int) {
			wg.Add(1)
			defer func() {
				wg.Done()
			}()
			t.Log(i, Feb(i))
		}(i)
	}
	wg.Wait() // 等待所有协程退出
}
