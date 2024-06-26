// Code generated by MockGen. DO NOT EDIT.
// Source: E:\workspace\github.com\bbs-spider\bbs-web\internal\repository\article\article_write_repo.go

// Package svcmocks is a generated GoMock package.
package svcmocks

import (
	domain "bbs-web/internal/domain"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockArticleWriterRepo is a mock of ArticleWriterRepo interface.
type MockArticleWriterRepo struct {
	ctrl     *gomock.Controller
	recorder *MockArticleWriterRepoMockRecorder
}

// MockArticleWriterRepoMockRecorder is the mock recorder for MockArticleWriterRepo.
type MockArticleWriterRepoMockRecorder struct {
	mock *MockArticleWriterRepo
}

// NewMockArticleWriterRepo creates a new mock instance.
func NewMockArticleWriterRepo(ctrl *gomock.Controller) *MockArticleWriterRepo {
	mock := &MockArticleWriterRepo{ctrl: ctrl}
	mock.recorder = &MockArticleWriterRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleWriterRepo) EXPECT() *MockArticleWriterRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockArticleWriterRepo) Create(ctx context.Context, art domain.Article) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, art)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockArticleWriterRepoMockRecorder) Create(ctx, art interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockArticleWriterRepo)(nil).Create), ctx, art)
}

// Update mocks base method.
func (m *MockArticleWriterRepo) Update(ctx context.Context, art domain.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, art)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockArticleWriterRepoMockRecorder) Update(ctx, art interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockArticleWriterRepo)(nil).Update), ctx, art)
}
