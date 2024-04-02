package service

import (
	"bbs-web/internal/service/article"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"testing"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-28 17:13

func TestBatchRankingService_TopN(t *testing.T) {
	tests := []struct {
		name string
		mock func(ctrl *gomock.Controller) article.articleService

		wantErr error
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			artSvc := tt.mock(ctrl)
			svc := NewBatchRankingService(artSvc)
			err := svc.TopN(context.Background())
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
