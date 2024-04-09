package service

import (
	"bbs-web/internal/domain"
	"bbs-web/internal/service/article"
	"bbs-web/internal/service/svcmocks"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-28 17:13

func TestBatchRankingService_TopN(t *testing.T) {
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

				return artSvc, intrSvc
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			artSvc, intrSvc := tt.mock(ctrl)
			svc := NewBatchRankingService(artSvc, intrSvc)
			arts, err := svc.topN(context.Background())
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantData, arts)
		})
	}
}
