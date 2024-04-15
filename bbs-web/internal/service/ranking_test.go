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
	nw := time.Now()
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
				artSvc.EXPECT().ListPub(gomock.Any(), nw, 0, 3).
					Return([]domain.Article{
						domain.Article{Id: 1, Utime: nw, Ctime: nw},
						domain.Article{Id: 2, Utime: nw, Ctime: nw},
						domain.Article{Id: 3, Utime: nw, Ctime: nw},
					}, nil)
				artSvc.EXPECT().ListPub(gomock.Any(), nw, 3, 3).
					Return([]domain.Article{}, nil)

				intrSvc.EXPECT().GetByIds(gomock.Any(), "article", []int64{
					1, 2, 3,
				}).Return(map[int64]domain.Interactive{
					1: domain.Interactive{BizId: 1, LikeCnt: 1},
					2: domain.Interactive{BizId: 2, LikeCnt: 2},
					3: domain.Interactive{BizId: 3, LikeCnt: 3},
				}, nil)
				intrSvc.EXPECT().GetByIds(gomock.Any(), "article", []int64{}).
					Return(map[int64]domain.Interactive{})
				return artSvc, intrSvc
			},
			wantData: []domain.Article{
				{Id: 3, Utime: nw, Ctime: nw},
				{Id: 2, Utime: nw, Ctime: nw},
				{Id: 1, Utime: nw, Ctime: nw},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			artSvc, intrSvc := tt.mock(ctrl)
			svc := NewBatchRankingService(artSvc, intrSvc)
			svc.batchSize = 3
			svc.queueCap = 3
			svc.scoreFn = func(t int64, likeCnt int64) float64 {
				return float64(likeCnt)
			}
			arts, err := svc.topN(context.Background())
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantData, arts)
		})
	}
}
