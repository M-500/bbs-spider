package LRU

import "testing"

// @Description
// @Author 代码小学生王木木

func TestLRUCache_Put(t *testing.T) {
	l := Constructor(2)
	l.Put(1, 1)
	l.Put(2, 2)
	t.Log(l.Get(1))
	l.Put(3, 3)

	t.Log(l.Get(2))
}
