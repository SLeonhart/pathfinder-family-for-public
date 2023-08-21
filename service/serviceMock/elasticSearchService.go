// Code generated by MockGen. DO NOT EDIT.
// Source: iElasticSearchService.go

// Package serviceMock is a generated GoMock package.
package serviceMock

import (
	context "context"
	model "pathfinder-family/model"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockIElasticSearchService is a mock of IElasticSearchService interface.
type MockIElasticSearchService struct {
	ctrl     *gomock.Controller
	recorder *MockIElasticSearchServiceMockRecorder
}

// MockIElasticSearchServiceMockRecorder is the mock recorder for MockIElasticSearchService.
type MockIElasticSearchServiceMockRecorder struct {
	mock *MockIElasticSearchService
}

// NewMockIElasticSearchService creates a new mock instance.
func NewMockIElasticSearchService(ctrl *gomock.Controller) *MockIElasticSearchService {
	mock := &MockIElasticSearchService{ctrl: ctrl}
	mock.recorder = &MockIElasticSearchServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIElasticSearchService) EXPECT() *MockIElasticSearchServiceMockRecorder {
	return m.recorder
}

// ClearOld mocks base method.
func (m *MockIElasticSearchService) ClearOld(ctx context.Context, updateTime time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearOld", ctx, updateTime)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearOld indicates an expected call of ClearOld.
func (mr *MockIElasticSearchServiceMockRecorder) ClearOld(ctx, updateTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearOld", reflect.TypeOf((*MockIElasticSearchService)(nil).ClearOld), ctx, updateTime)
}

// Get mocks base method.
func (m *MockIElasticSearchService) Get(ctx context.Context, searchString string) ([]model.ElasticResultHitsHits, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, searchString)
	ret0, _ := ret[0].([]model.ElasticResultHitsHits)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIElasticSearchServiceMockRecorder) Get(ctx, searchString interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIElasticSearchService)(nil).Get), ctx, searchString)
}

// UpdatePathfinderSearch mocks base method.
func (m *MockIElasticSearchService) UpdatePathfinderSearch(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePathfinderSearch", ctx)
}

// UpdatePathfinderSearch indicates an expected call of UpdatePathfinderSearch.
func (mr *MockIElasticSearchServiceMockRecorder) UpdatePathfinderSearch(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePathfinderSearch", reflect.TypeOf((*MockIElasticSearchService)(nil).UpdatePathfinderSearch), ctx)
}

// Upsert mocks base method.
func (m *MockIElasticSearchService) Upsert(ctx context.Context, updateTime time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", ctx, updateTime)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockIElasticSearchServiceMockRecorder) Upsert(ctx, updateTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockIElasticSearchService)(nil).Upsert), ctx, updateTime)
}
