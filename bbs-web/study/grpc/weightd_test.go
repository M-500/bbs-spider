package grpc

import (
	"fmt"
	"sync"
	"testing"
)

// @Description
// @Author 代码小学生王木木

type Node struct {
	Name          string
	weight        int
	currentWeight int
}

func (n *Node) Invoke() {
	// 发起了RPC调用
	fmt.Println("选中了", n.Name, n.currentWeight, n.weight)
}

func TestSmoothWRR(t *testing.T) {
	nodes := []*Node{
		{
			Name:          "A",
			weight:        10,
			currentWeight: 10,
		},
		{
			Name:          "B",
			weight:        20,
			currentWeight: 20,
		},
		{
			Name:          "C",
			weight:        30,
			currentWeight: 30,
		},
	}
	var b = &Balancer{
		nodes: nodes,
	}
	for i := 0; i < 10; i++ {
		pick := b.pick()
		pick.Invoke()
	}
}

type Balancer struct {
	nodes []*Node
	lock  sync.Mutex
}

func (b *Balancer) pick() *Node {
	b.lock.Lock()
	defer b.lock.Unlock()
	total := 0
	for i := 0; i < len(b.nodes); i++ {
		total = total + b.nodes[i].weight
	}
	// 更新当前权重
	for _, node := range b.nodes {
		node.currentWeight = node.weight + node.currentWeight
	}
	// 挑选一个
	var target *Node
	for _, node := range b.nodes {
		if target == nil {
			target = node
			continue
		}
		if target.currentWeight < node.currentWeight {
			target = node
			continue
		}
	}
	// 选中后要将当前权重减少
	target.currentWeight = target.currentWeight - total
	return target
}
