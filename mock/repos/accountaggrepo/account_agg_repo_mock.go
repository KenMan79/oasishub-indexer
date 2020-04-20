// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/figment-networks/oasishub-indexer/repos/accountaggrepo (interfaces: DbRepo)

// Package mock_accountaggrepo is a generated GoMock package.
package mock_accountaggrepo

import (
	accountagg "github.com/figment-networks/oasishub-indexer/models/accountagg"
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
func (m *MockDbRepo) Count() (*int64, errors.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count")
	ret0, _ := ret[0].(*int64)
	ret1, _ := ret[1].(errors.ApplicationError)
	return ret0, ret1
}

// Count indicates an expected call of Count
func (mr *MockDbRepoMockRecorder) Count() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockDbRepo)(nil).Count))
}

// Create mocks base method
func (m *MockDbRepo) Create(arg0 *accountagg.Model) errors.ApplicationError {
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

// Exists mocks base method
func (m *MockDbRepo) Exists(arg0 types.PublicKey) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exists indicates an expected call of Exists
func (mr *MockDbRepoMockRecorder) Exists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockDbRepo)(nil).Exists), arg0)
}

// GetByPublicKey mocks base method
func (m *MockDbRepo) GetByPublicKey(arg0 types.PublicKey) (*accountagg.Model, errors.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByPublicKey", arg0)
	ret0, _ := ret[0].(*accountagg.Model)
	ret1, _ := ret[1].(errors.ApplicationError)
	return ret0, ret1
}

// GetByPublicKey indicates an expected call of GetByPublicKey
func (mr *MockDbRepoMockRecorder) GetByPublicKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPublicKey", reflect.TypeOf((*MockDbRepo)(nil).GetByPublicKey), arg0)
}

// Save mocks base method
func (m *MockDbRepo) Save(arg0 *accountagg.Model) errors.ApplicationError {
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
