package service

import "context"

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-28 17:06

type RankingService interface {
	TopN(ctx context.Context) error
	//TopN(ctx context.Context,n int64) error
	//TopN(ctx context.Context, n int64) ([]any, error) // 后续写这个逻辑
}

//type BatchRankingService struct {
//	artSvc articleService
//}
//
//func NewBatchRankingService(artSvc articleService) *BatchRankingService {
//	return &BatchRankingService{artSvc: artSvc}
//}
//
//func (svc *BatchRankingService) TopN(ctx context.Context) error {
//	//// 1. 先拿一批数据
//	//for {
//	//	svc.artSvc.ListPub(context.Background())
//	//	// 要去对应的点赞数据
//	//
//	//	svc.artSvc.GetByIds()
//	//
//	//	// 合并计算socre
//	//	// 排序
//	//	// 最后得出结果
//	//}
//}
