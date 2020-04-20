// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/figment-networks/oasishub-indexer/repos/debondingdelegationseqrepo (interfaces: DbRepo)

// Package mock_debondingdelegationseqrepo is a generated GoMock package.
package mock_debondingdelegationseqrepo

import (
	debondingdelegationseq "github.com/figment-networks/oasishub-indexer/models/debondingdelegationseq"
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

// Create mocks base method
func (m *MockDbRepo) Create(arg0 *debondingdelegationseq.Model) errors.ApplicationError {
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
func (m *MockDbRepo) Exists(arg0 types.Height) bool {
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

// GetByHeight mocks base method
func (m *MockDbRepo) GetByHeight(arg0 types.Height) ([]debondingdelegationseq.Model, errors.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByHeight", arg0)
	ret0, _ := ret[0].([]debondingdelegationseq.Model)
	ret1, _ := ret[1].(errors.ApplicationError)
	return ret0, ret1
}

// GetByHeight indicates an expected call of GetByHeight
func (mr *MockDbRepoMockRecorder) GetByHeight(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByHeight", reflect.TypeOf((*MockDbRepo)(nil).GetByHeight), arg0)
}

// GetRecentByDelegatorUID mocks base method
func (m *MockDbRepo) GetRecentByDelegatorUID(arg0 types.PublicKey, arg1 int64) ([]debondingdelegationseq.Model, errors.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecentByDelegatorUID", arg0, arg1)
	ret0, _ := ret[0].([]debondingdelegationseq.Model)
	ret1, _ := ret[1].(errors.ApplicationError)
	return ret0, ret1
}

// GetRecentByDelegatorUID indicates an expected call of GetRecentByDelegatorUID
func (mr *MockDbRepoMockRecorder) GetRecentByDelegatorUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecentByDelegatorUID", reflect.TypeOf((*MockDbRepo)(nil).GetRecentByDelegatorUID), arg0, arg1)
}

// GetRecentByValidatorUID mocks base method
func (m *MockDbRepo) GetRecentByValidatorUID(arg0 types.PublicKey, arg1 int64) ([]debondingdelegationseq.Model, errors.ApplicationError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecentByValidatorUID", arg0, arg1)
	ret0, _ := ret[0].([]debondingdelegationseq.Model)
	ret1, _ := ret[1].(errors.ApplicationError)
	return ret0, ret1
}

// GetRecentByValidatorUID indicates an expected call of GetRecentByValidatorUID
func (mr *MockDbRepoMockRecorder) GetRecentByValidatorUID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecentByValidatorUID", reflect.TypeOf((*MockDbRepo)(nil).GetRecentByValidatorUID), arg0, arg1)
}
