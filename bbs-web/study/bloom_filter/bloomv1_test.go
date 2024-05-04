//@Author: wulinlin
//@Description:
//@File:  bloomv1_test
//@Version: 1.0.0
//@Date: 2024/05/04 20:24

package bloom_filter

import (
	"github.com/bits-and-blooms/bloom/v3"
	"testing"
)

func TestUseBloomFilter(t *testing.T) {
	// 创建布隆过滤器
	bf := bloom.NewWithEstimates(10000, 0.01)

	bf.Add([]byte("Hello"))
	bf.AddString("world") // 字符串可以使用这个API

	bf.Test([]byte("Hello")) // 判断是否存在
	bf.Test([]byte("A"))
}
