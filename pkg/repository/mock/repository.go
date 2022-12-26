// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	entity "github.com/CyberAgentHack/server-performance-tuning-2023/pkg/entity"
	repository "github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
	gomock "github.com/golang/mock/gomock"
)

// MockEpisode is a mock of Episode interface.
type MockEpisode struct {
	ctrl     *gomock.Controller
	recorder *MockEpisodeMockRecorder
}

// MockEpisodeMockRecorder is the mock recorder for MockEpisode.
type MockEpisodeMockRecorder struct {
	mock *MockEpisode
}

// NewMockEpisode creates a new mock instance.
func NewMockEpisode(ctrl *gomock.Controller) *MockEpisode {
	mock := &MockEpisode{ctrl: ctrl}
	mock.recorder = &MockEpisodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEpisode) EXPECT() *MockEpisodeMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockEpisode) List(ctx context.Context, params *repository.ListEpisodesParams) (entity.Episodes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, params)
	ret0, _ := ret[0].(entity.Episodes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockEpisodeMockRecorder) List(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockEpisode)(nil).List), ctx, params)
}

// MockSeries is a mock of Series interface.
type MockSeries struct {
	ctrl     *gomock.Controller
	recorder *MockSeriesMockRecorder
}

// MockSeriesMockRecorder is the mock recorder for MockSeries.
type MockSeriesMockRecorder struct {
	mock *MockSeries
}

// NewMockSeries creates a new mock instance.
func NewMockSeries(ctrl *gomock.Controller) *MockSeries {
	mock := &MockSeries{ctrl: ctrl}
	mock.recorder = &MockSeriesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSeries) EXPECT() *MockSeriesMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockSeries) List(ctx context.Context, params *repository.ListSeriesParams) (entity.SeriesMulti, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, params)
	ret0, _ := ret[0].(entity.SeriesMulti)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockSeriesMockRecorder) List(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSeries)(nil).List), ctx, params)
}

// MockSeason is a mock of Season interface.
type MockSeason struct {
	ctrl     *gomock.Controller
	recorder *MockSeasonMockRecorder
}

// MockSeasonMockRecorder is the mock recorder for MockSeason.
type MockSeasonMockRecorder struct {
	mock *MockSeason
}

// NewMockSeason creates a new mock instance.
func NewMockSeason(ctrl *gomock.Controller) *MockSeason {
	mock := &MockSeason{ctrl: ctrl}
	mock.recorder = &MockSeasonMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSeason) EXPECT() *MockSeasonMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockSeason) List(ctx context.Context, params *repository.ListSeasonsParams) (entity.Seasons, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, params)
	ret0, _ := ret[0].(entity.Seasons)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockSeasonMockRecorder) List(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSeason)(nil).List), ctx, params)
}

// MockViewingHistory is a mock of ViewingHistory interface.
type MockViewingHistory struct {
	ctrl     *gomock.Controller
	recorder *MockViewingHistoryMockRecorder
}

// MockViewingHistoryMockRecorder is the mock recorder for MockViewingHistory.
type MockViewingHistoryMockRecorder struct {
	mock *MockViewingHistory
}

// NewMockViewingHistory creates a new mock instance.
func NewMockViewingHistory(ctrl *gomock.Controller) *MockViewingHistory {
	mock := &MockViewingHistory{ctrl: ctrl}
	mock.recorder = &MockViewingHistoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockViewingHistory) EXPECT() *MockViewingHistoryMockRecorder {
	return m.recorder
}

// BatchGet mocks base method.
func (m *MockViewingHistory) BatchGet(ctx context.Context, ids []string, userID string) (entity.ViewingHistories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchGet", ctx, ids, userID)
	ret0, _ := ret[0].(entity.ViewingHistories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGet indicates an expected call of BatchGet.
func (mr *MockViewingHistoryMockRecorder) BatchGet(ctx, ids, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGet", reflect.TypeOf((*MockViewingHistory)(nil).BatchGet), ctx, ids, userID)
}

// Create mocks base method.
func (m *MockViewingHistory) Create(ctx context.Context, viewingHistory *entity.ViewingHistory) (*entity.ViewingHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, viewingHistory)
	ret0, _ := ret[0].(*entity.ViewingHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockViewingHistoryMockRecorder) Create(ctx, viewingHistory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockViewingHistory)(nil).Create), ctx, viewingHistory)
}

// MockGenre is a mock of Genre interface.
type MockGenre struct {
	ctrl     *gomock.Controller
	recorder *MockGenreMockRecorder
}

// MockGenreMockRecorder is the mock recorder for MockGenre.
type MockGenreMockRecorder struct {
	mock *MockGenre
}

// NewMockGenre creates a new mock instance.
func NewMockGenre(ctrl *gomock.Controller) *MockGenre {
	mock := &MockGenre{ctrl: ctrl}
	mock.recorder = &MockGenreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGenre) EXPECT() *MockGenreMockRecorder {
	return m.recorder
}

// BatchGet mocks base method.
func (m *MockGenre) BatchGet(ctx context.Context, ids []string) (entity.Genres, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchGet", ctx, ids)
	ret0, _ := ret[0].(entity.Genres)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGet indicates an expected call of BatchGet.
func (mr *MockGenreMockRecorder) BatchGet(ctx, ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGet", reflect.TypeOf((*MockGenre)(nil).BatchGet), ctx, ids)
}