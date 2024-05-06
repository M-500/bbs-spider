//@Author: wulinlin
//@Description: https://leetcode.cn/problems/design-skiplist/description/
//@File:  skiplist
//@Version: 1.0.0
//@Date: 2024/05/06 22:35

package skiplist

type skipNode struct {
	val     int
	forward []*skipNode
}

type Skiplist struct {
	head  *skipNode
	level int
}

func Constructor() Skiplist {

}

func (this *Skiplist) Search(target int) bool {

}

func (this *Skiplist) Add(num int) {

}

func (this *Skiplist) Erase(num int) bool {

}
