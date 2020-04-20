// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/figment-networks/oasishub-indexer/repos/syncablerepo (interfaces: DbRepo)

// Package mock_syncablerepo is a generated GoMock package.
package mock_syncablerepo

import (
	syncable "github.com/figment-networks/oasishub-indexer/models/syncable"
	types "github.com/figment-networks/oasishub-indexer/types"
	errors "github.com/figment-networks/oasishub-indexer/utils/errors"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDbRepo is a mock of DbRepo interface
type MockDbRepo struct {
	ctrl     *gomock.Controller
	recorder *MockDbRepoMockRecorder
}

// MockDbRepoMockRecorder is the mock recorder for MockDbRepo
type MockDbRepoMockRecorder struct {
	mock *MockDbRepo
}

// NewMockDbRepo creates a new mock instance
func NewMockDbRepo(ctrl *gomock.Controller) *MockDbRepo {
	mock := &MockDbRepo{ctrl: ctrl}
	mock.recorder = &MockDbRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDbRepo) EXPECT() *MockDbRepoMockRecorder {
	return m.recorder
}

// Count mocks base method
func (m *MockDbRepo) Count(arg0 syncable.Type) (*int64, errors.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", arg0)
	ret0, _ := ret[0].(*int64)
	ret1, _ := ret[1].(errors.ApplicationError)
	return ret0, ret1
}

// Count indicates an expected call of Count
func (mr *MockDbRepoMockRecorder) Count(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockDbRepo)(nil).Count), arg0)
}

// Create mocks base method
func (m *MockDbRepo) Create(arg0 *syncable.Model) errors.ApplicationError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(errors.ApplicationError)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockDbRepoMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDbRepo)(nil).Create), arg0)
}

// DeletePrevByHeight mocks base method
func (m *MockDbRepo) DeletePrevByHeight(arg0 types.Height) errors.ApplicationError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePrevByHeight", arg0)
	ret0, _ := ret[0].(errors.ApplicationError)
	return ret0
}

// DeletePrevByHeight indicates an expected call of DeletePrevByHeight
func (mr *MockDbRepoMockRecorder) DeletePrevByHeight(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePrevByHeight", reflect.TypeOf((*MockDbRepo)(nil).DeletePrevByHeight), arg0)
}

// Exists mocks base method
func (m *MockDbRepo) Exists(arg0 syncable.Type, arg1 types.Height) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exists indicates an expected call of Exists
func (mr *MockDbRepoMockRecorder) Exists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockDbRepo)(nil).Exists), arg0, arg1)
}

// GetByHeight mocks base method
func (m *MockDbRepo) GetByHeight(arg0 syncable.Type, arg1 types.Height) (*syncable.Model, errors.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByHeight", arg0, arg1)
	ret0, _ := ret[0].(*syncable.Model)
	ret1, _ := ret[1].(errors.ApplicationError)
	return ret0, ret1
}

// GetByHeight indicates an expected call of GetByHeight
func (mr *MockDbRepoMockRecorder) GetByHeight(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByHeight", reflect.TypeOf((*MockDbRepo)(nil).GetByHeight), arg0, arg1)
}

// GetMostRecent mocks base method
func (m *MockDbRepo) GetMostRecent(arg0 syncable.Type) (*syncable.Model, errors.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMostRecent", arg0)
	ret0, _ := ret[0].(*syncable.Model)
	ret1, _ := ret[1].(errors.ApplicationError)
	return ret0, ret1
}

// GetMostRecent indicates an expected call of GetMostRecent
func (mr *MockDbRepoMockRecorder) GetMostRecent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMostRecent", reflect.TypeOf((*MockDbRepo)(nil).GetMostRecent), arg0)
}

// GetMostRecentCommonHeight mocks base method
func (m *MockDbRepo) GetMostRecentCommonHeight() (*types.Height, errors.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMostRecentCommonHeight")
	ret0, _ := ret[0].(*types.Height)
	ret1, _ := ret[1].(errors.ApplicationError)
	return ret0, ret1
}

// GetMostRecentCommonHeight indicates an expected call of GetMostRecentCommonHeight
func (mr *MockDbRepoMockRecorder) GetMostRecentCommonHeight() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMostRecentCommonHeight", reflect.TypeOf((*MockDbRepo)(nil).GetMostRecentCommonHeight))
}

// Save mocks base method
func (m *MockDbRepo) Save(arg0 *syncable.Model) errors.ApplicationError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(errors.ApplicationError)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockDbRepoMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockDbRepo)(nil).Save), arg0)
}
