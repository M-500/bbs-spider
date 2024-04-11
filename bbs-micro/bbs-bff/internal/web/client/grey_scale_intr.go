package client

import (
	intrv1 "bbs-micro/api/proto/gen/proto/intr/v1"
	"context"
	"github.com/ecodeclub/ekit/syncx/atomicx"
	"google.golang.org/grpc"
	"math/rand"
)

// @Description  灰度调度
// @Author 代码小学生王木木
// @Date 2024-04-11 16:12

type GreyScaleInteractiveServiceClient struct {
	remote intrv1.InteractiveServiceClient
	local  intrv1.InteractiveServiceClient
	// 怎么控制流量的呢  一个请求过来，如何控制IPC调用 还是RPC调用 ？
	// 用随机数 + 阈值

	threshold *atomicx.Value[int32] // 一个阈值
}

func (g GreyScaleInteractiveServiceClient) IncrReadCnt(ctx context.Context, in *intrv1.IncrReadCntRequest, opts ...grpc.CallOption) (*intrv1.IncrReadCntResponse, error) {
	return g.client().IncrReadCnt(ctx, in, opts...)
}

func (g GreyScaleInteractiveServiceClient) Like(ctx context.Context, in *intrv1.LikeRequest, opts ...grpc.CallOption) (*intrv1.LikeResponse, error) {
	return g.client().Like(ctx, in, opts...)
}

func (g GreyScaleInteractiveServiceClient) CancelLike(ctx context.Context, in *intrv1.CancelLikeRequest, opts ...grpc.CallOption) (*intrv1.CancelLikeResponse, error) {
	return g.client().CancelLike(ctx, in, opts...)
}

func (g GreyScaleInteractiveServiceClient) CollectArt(ctx context.Context, in *intrv1.CollectArtRequest, opts ...grpc.CallOption) (*intrv1.CollectArtResponse, error) {
	return g.client().CollectArt(ctx, in, opts...)
}

func (g GreyScaleInteractiveServiceClient) Get(ctx context.Context, in *intrv1.GetRequest, opts ...grpc.CallOption) (*intrv1.GetResponse, error) {
	return g.client().Get(ctx, in, opts...)
}

func (g GreyScaleInteractiveServiceClient) GetByIds(ctx context.Context, in *intrv1.GetByIdsRequest, opts ...grpc.CallOption) (*intrv1.GetByIdsResponse, error) {
	return g.client().GetByIds(ctx, in, opts...)
}
func (g GreyScaleInteractiveServiceClient) UpdateThreshold(newThreshold int32) {
	g.threshold.Store(newThreshold)
}
func (g GreyScaleInteractiveServiceClient) client() intrv1.InteractiveServiceClient {
	threshold := g.threshold.Load()
	num := rand.Int31n(100) // 生成一个 0-100的随机数
	if num <= threshold {
		return g.remote
	}
	return g.local
}
