package wrr

import (
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
)

// @Description
// @Author 代码小学生王木木

const Name = "wrr_round_robin"

func init() {
	// base.NewBalancerBuilder 用于将我们自定义的Builder转换为 balancer.Builder
	balancer.Register(base.NewBalancerBuilder(Name, &PickerBuilder{}, base.Config{
		HealthCheck: true,
	}))
}

type PickerBuilder struct {
}

func (p *PickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	// 构建连接列表
	connList := make([]*node, 0, len(info.ReadySCs))
	// sc => SubConn
	// sci => SubConnInfo
	for sc, sci := range info.ReadySCs {
		conn := &node{
			conn: sc,
		}
		md, ok := sci.Address.Metadata.(map[string]any)
		if ok {
			weightVal, _ := md["weight"]
			weight, _ := weightVal.(float64)
			conn.weight = int(weight)
		}
		if conn.weight == 0 {
			conn.weight = 10 // 给个默认值，不给也没事
		}
		conn.currentWeight = conn.weight // 当前权重给一个默认值
		connList = append(connList, conn)
	}
	return &WrrPicker{
		conns: connList,
	}
}

type WrrPicker struct {
	conns []*node
}

// Pick
//
//	@Description: 实现基于权重的负载均衡算法
//	@receiver w
//	@param info
//	@return balancer.PickResult
//	@return error
func (w *WrrPicker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	if len(w.conns) == 0 {
		// 没有任何节点信息，要报错
		return balancer.PickResult{}, balancer.ErrNoSubConnAvailable
	}
	total := 0 // 总权重
	// 初始化总权重 初始化currentWeight
	for _, conn := range w.conns {
		total += conn.weight
		conn.currentWeight = conn.currentWeight + conn.weight
	}

	maxNode := w.conns[0]
	for _, conn := range w.conns {
		if conn.currentWeight > maxNode.currentWeight {
			maxNode = conn
		}
	}
	maxNode.currentWeight = maxNode.currentWeight - total // 这一步不要漏了
	return balancer.PickResult{
		SubConn: maxNode.conn,
		// RPC请求回调的时候，会调用这个方法
		Done: func(info balancer.DoneInfo) {
			// 如果希望通过调用结果来动态调整权重，可以在这个方法里做文章
		},
	}, nil
}

// 节点信息
type node struct {
	weight        int              // 权重信息
	currentWeight int              // 当前权重
	conn          balancer.SubConn // 连接信息
}
