package grpc

import (
	intrv1 "bbs-micro/api/proto/gen/proto/intr/v1"
	"bbs-micro/bbs-interactive/domain"
	"bbs-micro/bbs-interactive/service"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-10 12:27

// InteractiveServiceServer
// @Description: 这里只是把service包装成一个grpc而已，和grpc有关的操作就在这里限定
type InteractiveServiceServer struct {
	svc service.InteractiveService
	intrv1.UnimplementedInteractiveServiceServer
}

func (i *InteractiveServiceServer) IncrReadCnt(ctx context.Context, request *intrv1.IncrReadCntRequest) (*intrv1.IncrReadCntResponse, error) {
	err := i.svc.IncrReadCnt(ctx, request.Biz, request.BizId)
	if err != nil {
		return nil, err
	}
	return &intrv1.IncrReadCntResponse{}, nil // 只要error为nil 就不要返回数据域为nil 这个是标准写法
}

func (i *InteractiveServiceServer) Like(ctx context.Context, request *intrv1.LikeRequest) (*intrv1.LikeResponse, error) {
	err := i.svc.Like(ctx, request.GetBiz(), request.GetBizId(), request.GetUid()) // 使用Getxx()方法 要比.xx更优雅
	if err != nil {
		return nil, err
	}
	return &intrv1.LikeResponse{}, nil // 只要error为nil 就不要返回数据域为nil
}

func (i *InteractiveServiceServer) CancelLike(ctx context.Context, request *intrv1.CancelLikeRequest) (*intrv1.CancelLikeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (i *InteractiveServiceServer) CollectArt(ctx context.Context, request *intrv1.CollectArtRequest) (*intrv1.CollectArtResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (i *InteractiveServiceServer) Get(ctx context.Context, request *intrv1.GetRequest) (*intrv1.GetResponse, error) {
	get, err := i.svc.Get(ctx, request.GetBiz(), request.GetBizId(), request.GetUid())
	if err != nil {
		return nil, err
	}
	return &intrv1.GetResponse{
		Intr: i.toDTO(get),
	}, nil
}

func (i *InteractiveServiceServer) GetByIds(ctx context.Context, request *intrv1.GetByIdsRequest) (*intrv1.GetByIdsResponse, error) {
	//ids, err := i.svc.GetByIds(ctx, request.GetBiz(), request.GetIds())
	//if err != nil {
	//	return nil, err
	//}
	//res := &intrv1.GetByIdsResponse{
	//	Intrs: make(map[string]intrv1.Interactive),
	//}
	//for i2, interactive := range ids {
	//
	//}
	//return res, nil
	//TODO implement me
	panic("implement me")
}

func (i *InteractiveServiceServer) mustEmbedUnimplementedInteractiveServiceServer() {
	//TODO implement me
	panic("implement me")
}

// toDTO
//
//	@Description: Data transfer object
func (i *InteractiveServiceServer) toDTO(item domain.Interactive) *intrv1.Interactive {
	return &intrv1.Interactive{
		Biz:        item.Biz,
		BizId:      item.BizId,
		ReadCnt:    item.ReadCnt,
		LikeCnt:    item.LikeCnt,
		CollectCnt: item.CollectCnt,
		CommentCnt: item.CommentCnt,
		Liked:      item.Liked,
		Collected:  item.Collected,
	}
}
