package service

import (
	"bbs-web/internal/domain"
	"bbs-web/internal/service/article"
	"bbs-web/internal/service/svcmocks"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-28 17:13

func TestBatchRankingService_TopN(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name string
		mock func(ctrl *gomock.Controller) (article.IArticleService, InteractiveService)

		wantErr  error
		wantData []domain.Article
	}{
		{
			name: "计算成功",
			mock: func(ctrl *gomock.Controller) (article.IArticleService, InteractiveService) {
				artSvc := svcmocks.NewMockIArticleService(ctrl)
				intrSvc := svcmocks.NewMockInteractiveService(ctrl)
				// 先模拟批量获取数据
				// 先模拟第一批
				artSvc.EXPECT().ListPub(gomock.Any(), gomock.Any(), 0, 2).
					Return([]domain.Article{
						{Id: 1, Utime: now},
						{Id: 2, Utime: now},
					}, nil)
				// 模拟第二批
				artSvc.EXPECT().ListPub(gomock.Any(), gomock.Any(), 2, 2).
					Return([]domain.Article{
						{Id: 3, Utime: now},
						{Id: 4, Utime: now},
					}, nil)
				// 模拟第三批
				artSvc.EXPECT().ListPub(gomock.Any(), gomock.Any(), 4, 2).
					// 没数据了
					Return([]domain.Article{}, nil)

				// 第一批的点赞数据
				intrSvc.EXPECT().GetByIds(gomock.Any(), "article", []int64{1, 2}).
					Return(map[int64]domain.Interactive{
						1: {LikeCnt: 1},
						2: {LikeCnt: 2},
					}, nil)

				// 第二批的点赞数据
				intrSvc.EXPECT().GetByIds(gomock.Any(), "article", []int64{3, 4}).
					Return(map[int64]domain.Interactive{
						3: {LikeCnt: 3},
						4: {LikeCnt: 4},
					}, nil)

				// 第三批的点赞数据
				intrSvc.EXPECT().GetByIds(gomock.Any(), "article", []int64{}).
					Return(map[int64]domain.Interactive{}, nil)
				return artSvc, intrSvc
			},
			wantData: []domain.Article{
				{Id: 4, Utime: now},
				{Id: 3, Utime: now},
				{Id: 2, Utime: now},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			artSvc, intrSvc := tt.mock(ctrl)
			svc := &batchRankingService{
				artSvc:    artSvc,
				intrSvc:   intrSvc,
				batchSize: 2,
				queueCap:  3,
				scoreFn: func(t time.Time, likeCnt int64) float64 {
					return float64(likeCnt)
				},
			}
			arts, err := svc.topN(context.Background())
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantData, arts)
		})
	}
}
