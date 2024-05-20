//@Author: wulinlin
//@Description:
//@File:  sync_pool_test
//@Version: 1.0.0
//@Date: 2024/05/19 10:52

package sync_stu

import (
	"sync"
	"testing"
)

// 假设这是一个大对象
type BigObj struct {
	Name string
	Age  int
}

func TestUseSyncPool(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return BigObj{}
		},
	}

	data, ok := pool.Get().(BigObj) // 从对象池获取一个对象
	defer pool.Put(data)            // 归还对象给对象池
	if !ok {

	}
	data.Name = ""

}
