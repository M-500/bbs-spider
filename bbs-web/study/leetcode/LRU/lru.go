//@Author: wulinlin
//@Description:
//@File:  lru
//@Version: 1.0.0
//@Date: 2024/05/07 02:53

package LRU

// 双端链表的节点
type node struct {
	prev, next *node
	data       int
}

type LRUCache struct {
	maxSize int // 最大容量
	head    *node
	tail    *node
	hashMap map[int]*node // 哈希表
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		maxSize: capacity,
		head: &node{
			prev: nil,
			next: nil,
			data: 0,
		},
		tail:    nil,
		hashMap: make(map[int]*node, 0),
	}
}

func (this *LRUCache) Get(key int) int {
	// 从map中获取元素
	val, ok := this.hashMap[key]
	if !ok {
		return -1
	}
	// 说明有值，要将这个节点移动到栈顶
	this.moveToNo1(val)
	return val.data
}

func (this *LRUCache) Put(key int, value int) {
	if len(this.hashMap) >= this.maxSize {
		this.removeLast()
	}
	n := &node{
		data: value,
		next: this.head.next,
		prev: this.head,
	}
	this.hashMap[key] = n
	this.head.next.prev = n
	this.head.next = n
}
func (this *LRUCache) removeLast() {
	// 1. 移除链表的最后一个元素
	// 2. 删除最后一个元素在map中的记录
}
func (this *LRUCache) moveToNo1(cur *node) {
	if cur == nil {
		return
	}
	if cur.prev == this.head {
		// 说明已经是第一个了，不用管了
		return
	}
	cur.prev.next = cur.next
	cur.next.prev = cur.prev
	cur.next = this.head.next
	cur.prev = this.head
	this.head.next = cur
}
