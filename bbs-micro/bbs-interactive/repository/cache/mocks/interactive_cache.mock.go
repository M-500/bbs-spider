// Code generated by MockGen. DO NOT EDIT.
// Source: ./interactive_cache.go

// Package cachemocks is a generated GoMock package.
package cachemocks

import (
	domain "bbs-micro/bbs-interactive/domain"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRedisInteractiveCache is a mock of RedisInteractiveCache interface.
type MockRedisInteractiveCache struct {
	ctrl     *gomock.Controller
	recorder *MockRedisInteractiveCacheMockRecorder
}

// MockRedisInteractiveCacheMockRecorder is the mock recorder for MockRedisInteractiveCache.
type MockRedisInteractiveCacheMockRecorder struct {
	mock *MockRedisInteractiveCache
}

// NewMockRedisInteractiveCache creates a new mock instance.
func NewMockRedisInteractiveCache(ctrl *gomock.Controller) *MockRedisInteractiveCache {
	mock := &MockRedisInteractiveCache{ctrl: ctrl}
	mock.recorder = &MockRedisInteractiveCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisInteractiveCache) EXPECT() *MockRedisInteractiveCacheMockRecorder {
	return m.recorder
}

// DecrLikeCntIfPresent mocks base method.
func (m *MockRedisInteractiveCache) DecrLikeCntIfPresent(ctx context.Context, biz string, bizId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecrLikeCntIfPresent", ctx, biz, bizId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DecrLikeCntIfPresent indicates an expected call of DecrLikeCntIfPresent.
func (mr *MockRedisInteractiveCacheMockRecorder) DecrLikeCntIfPresent(ctx, biz, bizId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecrLikeCntIfPresent", reflect.TypeOf((*MockRedisInteractiveCache)(nil).DecrLikeCntIfPresent), ctx, biz, bizId)
}

// Get mocks base method.
func (m *MockRedisInteractiveCache) Get(ctx context.Context, biz string, bizId int64) (domain.Interactive, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, biz, bizId)
	ret0, _ := ret[0].(domain.Interactive)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRedisInteractiveCacheMockRecorder) Get(ctx, biz, bizId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRedisInteractiveCache)(nil).Get), ctx, biz, bizId)
}

// IncrCollCntIfPresent mocks base method.
func (m *MockRedisInteractiveCache) IncrCollCntIfPresent(ctx context.Context, biz string, bizId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrCollCntIfPresent", ctx, biz, bizId)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrCollCntIfPresent indicates an expected call of IncrCollCntIfPresent.
func (mr *MockRedisInteractiveCacheMockRecorder) IncrCollCntIfPresent(ctx, biz, bizId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrCollCntIfPresent", reflect.TypeOf((*MockRedisInteractiveCache)(nil).IncrCollCntIfPresent), ctx, biz, bizId)
}

// IncrLikeCntIfPresent mocks base method.
func (m *MockRedisInteractiveCache) IncrLikeCntIfPresent(ctx context.Context, biz string, bizId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrLikeCntIfPresent", ctx, biz, bizId)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrLikeCntIfPresent indicates an expected call of IncrLikeCntIfPresent.
func (mr *MockRedisInteractiveCacheMockRecorder) IncrLikeCntIfPresent(ctx, biz, bizId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrLikeCntIfPresent", reflect.TypeOf((*MockRedisInteractiveCache)(nil).IncrLikeCntIfPresent), ctx, biz, bizId)
}

// IncrReadCntIfPresent mocks base method.
func (m *MockRedisInteractiveCache) IncrReadCntIfPresent(ctx context.Context, biz string, bizId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrReadCntIfPresent", ctx, biz, bizId)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrReadCntIfPresent indicates an expected call of IncrReadCntIfPresent.
func (mr *MockRedisInteractiveCacheMockRecorder) IncrReadCntIfPresent(ctx, biz, bizId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrReadCntIfPresent", reflect.TypeOf((*MockRedisInteractiveCache)(nil).IncrReadCntIfPresent), ctx, biz, bizId)
}

// Set mocks base method.
func (m *MockRedisInteractiveCache) Set(ctx context.Context, biz string, bizId int64, intr domain.Interactive) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, biz, bizId, intr)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockRedisInteractiveCacheMockRecorder) Set(ctx, biz, bizId, intr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockRedisInteractiveCache)(nil).Set), ctx, biz, bizId, intr)
}
