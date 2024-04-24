package grpc

import (
	"bbs-web/pkg/utils/zifo/slice"
	"fmt"
	"math/rand"
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
	//fmt.Println("选中了", n.Name, n.currentWeight)
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
		fmt.Printf("挑选前 nodes: %v\n", slice.Map(nodes, func(idx int, src *Node) Node {
			return *src
		}))
		pick := b.pick()
		fmt.Printf("挑选后 nodes: %v\n", slice.Map(nodes, func(idx int, src *Node) Node {
			return *src
		}))
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
	for _, node := range b.nodes {
		total += node.weight
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
		} else {
			if target.currentWeight <= node.currentWeight {
				target = node
			}
		}
	}
	// 选中后要将当前权重减少
	fmt.Println("选中了", target.Name, "当前权重", target.currentWeight, target.weight)
	target.currentWeight = target.currentWeight - total
	fmt.Println("选完后", target.Name, "当前权重减去总权重", target.currentWeight, target.weight)
	return target
}

// randomWeightPicker
//
//	@Description: 加权随机
//	@receiver b*Balancer
func (b *Balancer) randomWeightPicker() *Node {
	b.lock.Lock()
	defer b.lock.Unlock()
	// 计算总权重
	total := 0
	for _, node := range b.nodes {
		total += node.weight
	}
	r := rand.Int31n(int32(total))
	//// 求模运算
	//remainder := randomNumber % divisor
	for _, node := range b.nodes {
		r = r - int32(total)
		if r < 0 {
			return node
		}
	}
	panic("")
}
