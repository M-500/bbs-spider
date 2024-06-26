package handler

/**
这个是单元测试 区别于 integration 目录下的测试文件，后者是集成测试
*/
import (
	"bbs-web/internal/domain"
	"bbs-web/internal/service"
	"bbs-web/internal/service/article"
	"bbs-web/internal/service/svcmocks"
	"bbs-web/internal/web/vo"
	"bbs-web/pkg/ginplus"
	"bbs-web/pkg/logger"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-01 15:43

func TestArticleHandler_Publish(t *testing.T) {
	serverTest := gin.Default()
	tests := []struct {
		name string

		mock func(ctrl *gomock.Controller) (article.IArticleService, service.InteractiveService, logger.Logger)

		reqBody string

		wantCode int

		wantBody ginplus.Result
	}{
		// 还有一些，1. 修改已有文章，并且发表成功  2. Bind返回错误  3. 找不到User 4.publish返回错误
		{
			name: "新建并发表",
			mock: func(ctrl *gomock.Controller) (article.IArticleService, service.InteractiveService, logger.Logger) {
				svc := svcmocks.NewMockIArticleService(ctrl)
				svc.EXPECT().Publish(gomock.Any(), domain.Article{
					Id:      0,
					Title:   "我的标题",
					Content: "##小别致听东西的啊不然呢",
					Author: domain.Author{
						Id: 1,
					},
					Status:      0,
					Summary:     "搞事情饿",
					ContentType: "article",
					Cover:       "",
					Ctime:       time.Time{},
					Utime:       time.Time{},
				}).Return(int64(1), nil)
				return svc, nil, nil
			},
			reqBody: `{
				"id":0,
				"content_type":"article",
				"title":"我的标题",
				"summary":"搞事情饿",
				"cover":"",
				"content":"##小别致听东西的啊不然呢"}`,
			wantCode: 200,
			wantBody: ginplus.Result{
				Data: float64(1), // 因为JSON会默认将整数转换为float64
				Msg:  "OK",
			},
		},
		{
			name: "发布失败",
			mock: func(ctrl *gomock.Controller) (article.IArticleService, service.InteractiveService, logger.Logger) {
				svc := svcmocks.NewMockIArticleService(ctrl)
				svc.EXPECT().Publish(gomock.Any(), domain.Article{
					Id:      0,
					Title:   "我的标题",
					Content: "##小别致听东西的啊不然呢",
					Author: domain.Author{
						Id: 1,
					},
					Status:      0,
					Summary:     "搞事情饿",
					ContentType: "article",
					Cover:       "",
					Ctime:       time.Time{},
					Utime:       time.Time{},
				}).Return(int64(0), errors.New("publish 失败"))
				return svc, nil, nil
			},
			reqBody: `{
				"id":0,
				"content_type":"article",
				"title":"我的标题",
				"summary":"搞事情饿",
				"cover":"",
				"content":"##小别致听东西的啊不然呢"}`,
			wantCode: 200,
			wantBody: ginplus.Result{
				Code: 510003, // 因为JSON会默认将整数转换为float64
				Msg:  "保存帖子失败",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			h := NewArticleHandler(tt.mock(ctrl))
			serverTest.POST("/articles/publish", ginplus.WrapJson[vo.ArticleReq](h.Publish))

			req, err := http.NewRequest(http.MethodPost, "/articles/publish", bytes.NewBuffer([]byte(tt.reqBody)))
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			// 这里你就可以继续使用 req
			resp := httptest.NewRecorder()

			serverTest.ServeHTTP(resp, req) // 使用测试套件里的Server对象
			assert.Equal(t, tt.wantCode, resp.Code)
			if resp.Code != 200 {
				return
			}
			require.NoError(t, err)
			var webRes ginplus.Result
			err = json.NewDecoder(resp.Body).Decode(&webRes)
			assert.Equal(t, tt.wantBody, webRes)

		})
	}
}
