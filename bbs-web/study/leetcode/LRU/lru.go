//@Author: wulinlin
//@Description:
//@File:  lru
//@Version: 1.0.0
//@Date: 2024/05/07 02:53

package LRU

import "container/list"

type node struct {
}

type LRUCache struct {
	maxSize  int        // 最大容量
	linkList *list.List // 双向链表

}

func Constructor(capacity int) LRUCache {

}

func (this *LRUCache) Get(key int) int {

}

func (this *LRUCache) Put(key int, value int) {

}
