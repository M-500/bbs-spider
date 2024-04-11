package client

import (
	intrv1 "bbs-micro/api/proto/gen/proto/intr/v1"
	"bbs-micro/bbs-bff/internal/domain"
	"bbs-micro/bbs-bff/internal/service"
	"context"
	"google.golang.org/grpc"
)

// @Description  适配器流量分发
// @Author 代码小学生王木木
// @Date 2024-04-11 15:45

// 实现 intrv1.InteractiveServiceClient 这个接口

// InteractiveServiceAdapter
// @Description: 将一个本地的实现 伪装为GPRC客户端
type InteractiveServiceAdapter struct {
	svc service.InteractiveService
}

func (i *InteractiveServiceAdapter) IncrReadCnt(ctx context.Context, in *intrv1.IncrReadCntRequest, opts ...grpc.CallOption) (*intrv1.IncrReadCntResponse, error) {
	err := i.svc.IncrReadCnt(ctx, in.GetBiz(), in.GetBizId())
	return &intrv1.IncrReadCntResponse{}, err
}

func (i *InteractiveServiceAdapter) Like(ctx context.Context, in *intrv1.LikeRequest, opts ...grpc.CallOption) (*intrv1.LikeResponse, error) {
	err := i.svc.Like(ctx, in.GetBiz(), in.GetBizId(), in.GetUid())
	return &intrv1.LikeResponse{}, err
}

func (i *InteractiveServiceAdapter) CancelLike(ctx context.Context, in *intrv1.CancelLikeRequest, opts ...grpc.CallOption) (*intrv1.CancelLikeResponse, error) {
	err := i.svc.CancelLike(ctx, in.GetBiz(), in.GetBizId(), in.GetBizId())
	return &intrv1.CancelLikeResponse{}, err
}

func (i *InteractiveServiceAdapter) CollectArt(ctx context.Context, in *intrv1.CollectArtRequest, opts ...grpc.CallOption) (*intrv1.CollectArtResponse, error) {
	err := i.svc.CollectArt(ctx, in.GetBiz(), in.GetBizId(), in.GetUid(), in.GetCid())
	return &intrv1.CollectArtResponse{}, err
}

func (i *InteractiveServiceAdapter) Get(ctx context.Context, in *intrv1.GetRequest, opts ...grpc.CallOption) (*intrv1.GetResponse, error) {
	resp, err := i.svc.Get(ctx, in.GetBiz(), in.GetBizId(), in.GetUid())
	if err != nil {
		return &intrv1.GetResponse{}, err
	}
	return &intrv1.GetResponse{
		Intr: i.toDTO(resp),
	}, nil
}

func (i *InteractiveServiceAdapter) GetByIds(ctx context.Context, in *intrv1.GetByIdsRequest, opts ...grpc.CallOption) (*intrv1.GetByIdsResponse, error) {
	ids, err := i.svc.GetByIds(ctx, in.GetBiz(), in.GetIds())
	if err != nil {
		return &intrv1.GetByIdsResponse{}, err
	}
	m := make(map[int64]*intrv1.Interactive, len(ids))
	for k, v := range m {
		m[k] = v
	}
	return &intrv1.GetByIdsResponse{
		Intrs: m,
	}, nil
}

func (i *InteractiveServiceAdapter) toDTO(data domain.Interactive) *intrv1.Interactive {
	return &intrv1.Interactive{
		Biz:        data.Biz,
		BizId:      data.BizId,
		ReadCnt:    data.ReadCnt,
		LikeCnt:    data.LikeCnt,
		CollectCnt: data.CollectCnt,
		CommentCnt: data.CommentCnt,
		Liked:      data.Liked,
		Collected:  data.Collected,
	}
}
