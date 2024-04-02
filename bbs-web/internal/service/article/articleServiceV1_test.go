package article

import (
	"bbs-web/internal/domain"
	"bbs-web/internal/repository/article"
	"bbs-web/internal/repository/article/svcmocks"
	"bbs-web/pkg/logger"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-04-01 18:02

//func TestNewArticleService(t *testing.T) {
//	type args struct {
//		repo article.ArticleRepository
//	}
//	tests := []struct {
//		name string
//		args args
//		want IArticleService
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewArticleService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewArticleService() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func Test_articleService_GetById(t *testing.T) {
//	type fields struct {
//		repo article.ArticleRepository
//	}
//	type args struct {
//		ctx context.Context
//		id  int64
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    domain.Article
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			svc := &articleService{
//				repo: tt.fields.repo,
//			}
//			got, err := svc.GetById(tt.args.ctx, tt.args.id)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetById() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetById() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func Test_articleService_GetByIds(t *testing.T) {
//	type fields struct {
//		repo article.ArticleRepository
//	}
//	type args struct {
//		ctx context.Context
//		biz string
//		ids []int64
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    []domain.Article
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			svc := &articleService{
//				repo: tt.fields.repo,
//			}
//			got, err := svc.GetByIds(tt.args.ctx, tt.args.biz, tt.args.ids)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetByIds() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetByIds() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_articleService_GetPublishedById(t *testing.T) {
//	type fields struct {
//		repo article.ArticleRepository
//	}
//	type args struct {
//		ctx context.Context
//		id  int64
//		uid int64
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    domain.Article
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			svc := &articleService{
//				repo: tt.fields.repo,
//			}
//			got, err := svc.GetPublishedById(tt.args.ctx, tt.args.id, tt.args.uid)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetPublishedById() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetPublishedById() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_articleService_List(t *testing.T) {
//	type fields struct {
//		repo article.ArticleRepository
//	}
//	type args struct {
//		ctx    context.Context
//		uid    int64
//		offset int
//		limit  int
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    []domain.Article
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			svc := &articleService{
//				repo: tt.fields.repo,
//			}
//			got, err := svc.List(tt.args.ctx, tt.args.uid, tt.args.offset, tt.args.limit)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("List() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_articleService_ListPub(t *testing.T) {
//	type fields struct {
//		repo article.ArticleRepository
//	}
//	type args struct {
//		ctx    context.Context
//		start  time.Time
//		offset int
//		limit  int
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    []domain.Article
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			svc := &articleService{
//				repo: tt.fields.repo,
//			}
//			got, err := svc.ListPub(tt.args.ctx, tt.args.start, tt.args.offset, tt.args.limit)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("ListPub() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("ListPub() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func Test_articleService_Publish(t *testing.T) {
//	testCast := []struct {
//		name string
//
//		mock func(ctrl *gomock.Controller) article.ArticleRepository
//
//		article domain.Article
//
//		wantErr error
//		wantId  int64
//	}{
//		{
//			name: "新建发表成功",
//			mock: func(ctl *gomock.Controller) article.ArticleRepository {
//				return nil
//			},
//			article: domain.Article{
//				Title:   "赋能",
//				Content: "请求赋能，请求赋能！",
//				Author: domain.Author{
//					Id: 123,
//				},
//				Status:      0,
//				Summary:     "分割线",
//				ContentType: "blog",
//			},
//			wantErr: nil,
//			wantId:  1,
//		},
//	}
//	for _, tc := range testCast {
//		t.Run(tc.name, func(t *testing.T) {
//			ctrl := gomock.NewController(t)
//			defer ctrl.Finish()
//
//			svc := NewArticleService(tc.mock(ctrl))
//			aId, err := svc.Publish(context.Background(), tc.article)
//			assert.Equal(t, tc.wantErr, err)
//			assert.Equal(t, aId, tc.wantId)
//
//		})
//	}
//}

func Test_articleService_PublishV1(t *testing.T) {
	testCast := []struct {
		name string

		mock func(ctrl *gomock.Controller) (article.ArticleWriterRepo, article.ArticleReaderRepository)

		article domain.Article

		wantErr error
		wantId  int64
	}{
		// 这里测试用例还不够完善，应该是所有条件的笛卡尔积(测试覆盖率要达标)
		{
			name: "新建发表成功",
			mock: func(ctl *gomock.Controller) (article.ArticleWriterRepo, article.ArticleReaderRepository) {
				writerRepo := svcmocks.NewMockArticleWriterRepo(ctl)

				writerRepo.EXPECT().Create(gomock.Any(), domain.Article{
					Title:   "赋能",
					Content: "请求赋能，请求赋能！",
					Author: domain.Author{
						Id: 123,
					},
					Status:      0,
					Summary:     "分割线",
					ContentType: "blog",
				}).Return(int64(1), nil)

				readerRepo := svcmocks.NewMockArticleReaderRepository(ctl)
				readerRepo.EXPECT().Save(gomock.Any(), domain.Article{
					// 确保使用了制作库 ID
					Id:      1,
					Title:   "赋能",
					Content: "请求赋能，请求赋能！",
					Author: domain.Author{
						Id: 123,
					},
					Status:      0,
					Summary:     "分割线",
					ContentType: "blog",
				}).Return(int64(1), nil)
				return writerRepo, readerRepo
			},
			article: domain.Article{
				Title:   "赋能",
				Content: "请求赋能，请求赋能！",
				Author: domain.Author{
					Id: 123,
				},
				Status:      0,
				Summary:     "分割线",
				ContentType: "blog",
			},
			wantErr: nil,
			wantId:  1,
		},
		{
			name: "修改并发表成功",
			mock: func(ctl *gomock.Controller) (article.ArticleWriterRepo, article.ArticleReaderRepository) {
				writerRepo := svcmocks.NewMockArticleWriterRepo(ctl)

				writerRepo.EXPECT().Update(gomock.Any(), domain.Article{
					Id:      2,
					Title:   "赋能",
					Content: "请求赋能，请求赋能！",
					Author: domain.Author{
						Id: 123,
					},
					Status:      0,
					Summary:     "分割线",
					ContentType: "blog",
				}).Return(nil)

				readerRepo := svcmocks.NewMockArticleReaderRepository(ctl)
				readerRepo.EXPECT().Save(gomock.Any(), domain.Article{
					// 确保使用了制作库 ID
					Id:      2,
					Title:   "赋能",
					Content: "请求赋能，请求赋能！",
					Author: domain.Author{
						Id: 123,
					},
					Status:      0,
					Summary:     "分割线",
					ContentType: "blog",
				}).Return(int64(2), nil)
				return writerRepo, readerRepo
			},
			article: domain.Article{
				Id:      2,
				Title:   "赋能",
				Content: "请求赋能，请求赋能！",
				Author: domain.Author{
					Id: 123,
				},
				Status:      0,
				Summary:     "分割线",
				ContentType: "blog",
			},
			wantErr: nil,
			wantId:  2,
		},
		{
			// 修改-保存到制作库失败
			name: "新建-保存到制作库失败",
			mock: func(ctl *gomock.Controller) (article.ArticleWriterRepo, article.ArticleReaderRepository) {
				writerRepo := svcmocks.NewMockArticleWriterRepo(ctl)

				writerRepo.EXPECT().Create(gomock.Any(), domain.Article{
					Title:   "赋能",
					Content: "请求赋能，请求赋能！",
					Author: domain.Author{
						Id: 123,
					},
					Status:      0,
					Summary:     "分割线",
					ContentType: "blog",
				}).Return(int64(0), errors.New("mock db error"))
				readerRepo := svcmocks.NewMockArticleReaderRepository(ctl)
				return writerRepo, readerRepo
			},
			article: domain.Article{
				Title:   "赋能",
				Content: "请求赋能，请求赋能！",
				Author: domain.Author{
					Id: 123,
				},
				Status:      0,
				Summary:     "分割线",
				ContentType: "blog",
			},
			wantErr: errors.New("mock db error"),
			wantId:  0,
		},
		{
			name: "修改-保存到制作库失败",
			mock: func(ctl *gomock.Controller) (article.ArticleWriterRepo, article.ArticleReaderRepository) {
				writerRepo := svcmocks.NewMockArticleWriterRepo(ctl)

				writerRepo.EXPECT().Update(gomock.Any(), domain.Article{
					Id:      3,
					Title:   "赋能",
					Content: "请求赋能，请求赋能！",
					Author: domain.Author{
						Id: 123,
					},
					Status:      0,
					Summary:     "分割线",
					ContentType: "blog",
				}).Return(errors.New("mock db error"))
				readerRepo := svcmocks.NewMockArticleReaderRepository(ctl)
				return writerRepo, readerRepo
			},
			article: domain.Article{
				Id:      3,
				Title:   "赋能",
				Content: "请求赋能，请求赋能！",
				Author: domain.Author{
					Id: 123,
				},
				Status:      0,
				Summary:     "分割线",
				ContentType: "blog",
			},
			wantErr: errors.New("mock db error"),
			wantId:  0,
		},
		{
			name: "保存到制作库成功，重试到线上库成功",
			mock: func(ctl *gomock.Controller) (article.ArticleWriterRepo, article.ArticleReaderRepository) {
				writerRepo := svcmocks.NewMockArticleWriterRepo(ctl)
				writerRepo.EXPECT().Update(gomock.Any(), domain.Article{
					Id:      5,
					Title:   "赋能",
					Content: "请求赋能，请求赋能！",
					Author: domain.Author{
						Id: 123,
					},
					Status:      0,
					Summary:     "分割线",
					ContentType: "blog",
				}).Return(nil)

				readerRepo := svcmocks.NewMockArticleReaderRepository(ctl)
				readerRepo.EXPECT().Save(gomock.Any(), domain.Article{
					// 确保使用了制作库 ID
					Id:      5,
					Title:   "赋能",
					Content: "请求赋能，请求赋能！",
					Author: domain.Author{
						Id: 123,
					},
					Status:      0,
					Summary:     "分割线",
					ContentType: "blog",
				}).Return(int64(0), errors.New("mock db error"))
				readerRepo.EXPECT().Save(gomock.Any(), domain.Article{
					// 确保使用了制作库 ID
					Id:      5,
					Title:   "赋能",
					Content: "请求赋能，请求赋能！",
					Author: domain.Author{
						Id: 123,
					},
					Status:      0,
					Summary:     "分割线",
					ContentType: "blog",
				}).Return(int64(5), nil)
				return writerRepo, readerRepo
			},
			article: domain.Article{
				Id:      5,
				Title:   "赋能",
				Content: "请求赋能，请求赋能！",
				Author: domain.Author{
					Id: 123,
				},
				Status:      0,
				Summary:     "分割线",
				ContentType: "blog",
			},
			wantErr: nil,
			wantId:  5,
		},
		{
			name: "保存到制作库成功，且重试N次全部失败",
			mock: func(ctl *gomock.Controller) (article.ArticleWriterRepo, article.ArticleReaderRepository) {
				writerRepo := svcmocks.NewMockArticleWriterRepo(ctl)
				writerRepo.EXPECT().Update(gomock.Any(), domain.Article{
					Id:      5,
					Title:   "赋能",
					Content: "请求赋能，请求赋能！",
					Author: domain.Author{
						Id: 123,
					},
					Status:      0,
					Summary:     "分割线",
					ContentType: "blog",
				}).Return(nil)

				readerRepo := svcmocks.NewMockArticleReaderRepository(ctl)
				readerRepo.EXPECT().Save(gomock.Any(), domain.Article{
					// 确保使用了制作库 ID
					Id:      5,
					Title:   "赋能",
					Content: "请求赋能，请求赋能！",
					Author: domain.Author{
						Id: 123,
					},
					Status:      0,
					Summary:     "分割线",
					ContentType: "blog",
				}).Times(5).Return(int64(0), errors.New("mock db error"))
				return writerRepo, readerRepo
			},
			article: domain.Article{
				Id:      5,
				Title:   "赋能",
				Content: "请求赋能，请求赋能！",
				Author: domain.Author{
					Id: 123,
				},
				Status:      0,
				Summary:     "分割线",
				ContentType: "blog",
			},
			wantErr: errors.New("mock db error"),
			wantId:  0,
		},
	}
	for _, tc := range testCast {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			write, reader := tc.mock(ctrl)
			svc := NewArticleServiceV1(write, reader, logger.NewNoOpLogger())
			aId, err := svc.PublishV1(context.Background(), tc.article)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, aId, tc.wantId)

		})
	}
}

//
//func Test_articleService_Save(t *testing.T) {
//	type fields struct {
//		repo article.ArticleRepository
//	}
//	type args struct {
//		ctx context.Context
//		art domain.Article
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    int64
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			svc := &articleService{
//				repo: tt.fields.repo,
//			}
//			got, err := svc.Save(tt.args.ctx, tt.args.art)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("Save() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func Test_articleService_Withdraw(t *testing.T) {
//	type fields struct {
//		repo article.ArticleRepository
//	}
//	type args struct {
//		ctx context.Context
//		art domain.Article
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			svc := &articleService{
//				repo: tt.fields.repo,
//			}
//			if err := svc.Withdraw(tt.args.ctx, tt.args.art); (err != nil) != tt.wantErr {
//				t.Errorf("Withdraw() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
