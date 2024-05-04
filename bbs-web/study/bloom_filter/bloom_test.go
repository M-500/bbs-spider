//@Author: wulinlin
//@Description:
//@File:  bloom_test
//@Version: 1.0.0
//@Date: 2024/05/04 17:14

package bloom_filter

import (
	"context"
	"fmt"
	"github.com/bits-and-blooms/bloom/v3"
	"testing"
)

type BloomFilter interface {
	HashKey(ctx context.Context, key string) bool
}

type BloomFilterCache struct {
}

func (b *BloomFilterCache) HashKey(ctx context.Context, key string) bool {
	//TODO implement me
	panic("implement me")
}

func TestBloom(t *testing.T) {
	filter := bloom.NewWithEstimates(1000, 0.01)
	filter.Add([]byte("Wulinlin"))
	fmt.Println(filter.Test([]byte("wulinlin")))
	fmt.Println(filter.Test([]byte("Wulinlin")))
}
