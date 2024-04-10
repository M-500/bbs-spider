package grpc

import (
	intrv1 "bbs-micro/api/bbs-micro/api/proto/gen/proto/intr/v1"
	"context"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-10 12:27

// InteractiveServiceServer
// @Description: 这里只是把service包装成一个grpc而已，和grpc有关的操作就在这里限定
type InteractiveServiceServer struct {
}

func (i *InteractiveServiceServer) IncrReadCnt(ctx context.Context, request *intrv1.IncrReadCntRequest) (*intrv1.IncrReadCntResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (i *InteractiveServiceServer) Like(ctx context.Context, request *intrv1.LikeRequest) (*intrv1.LikeResponse, error) {
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
}

func (i *InteractiveServiceServer) GetByIds(ctx context.Context, request *intrv1.GetByIdsRequest) (*intrv1.GetByIdsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (i *InteractiveServiceServer) mustEmbedUnimplementedInteractiveServiceServer() {
	//TODO implement me
	panic("implement me")
}
