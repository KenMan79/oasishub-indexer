// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/figment-networks/oasishub-indexer/indexer (interfaces: TargetsReader,SystemEventCreatorStore)

// Package mock_indexer is a generated GoMock package.
package mock_indexer

import (
	pipeline "github.com/figment-networks/indexing-engine/pipeline"
	model "github.com/figment-networks/oasishub-indexer/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTargetsReader is a mock of TargetsReader interface
type MockTargetsReader struct {
	ctrl     *gomock.Controller
	recorder *MockTargetsReaderMockRecorder
}

// MockTargetsReaderMockRecorder is the mock recorder for MockTargetsReader
type MockTargetsReaderMockRecorder struct {
	mock *MockTargetsReader
}

// NewMockTargetsReader creates a new mock instance
func NewMockTargetsReader(ctrl *gomock.Controller) *MockTargetsReader {
	mock := &MockTargetsReader{ctrl: ctrl}
	mock.recorder = &MockTargetsReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTargetsReader) EXPECT() *MockTargetsReaderMockRecorder {
	return m.recorder
}

// GetAllAvailableTasks mocks base method
func (m *MockTargetsReader) GetAllAvailableTasks() []pipeline.TaskName {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAvailableTasks")
	ret0, _ := ret[0].([]pipeline.TaskName)
	return ret0
}

// GetAllAvailableTasks indicates an expected call of GetAllAvailableTasks
func (mr *MockTargetsReaderMockRecorder) GetAllAvailableTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAvailableTasks", reflect.TypeOf((*MockTargetsReader)(nil).GetAllAvailableTasks))
}

// GetAllVersionedTasks mocks base method
func (m *MockTargetsReader) GetAllVersionedTasks() ([]pipeline.TaskName, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllVersionedTasks")
	ret0, _ := ret[0].([]pipeline.TaskName)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllVersionedTasks indicates an expected call of GetAllVersionedTasks
func (mr *MockTargetsReaderMockRecorder) GetAllVersionedTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllVersionedTasks", reflect.TypeOf((*MockTargetsReader)(nil).GetAllVersionedTasks))
}

// GetAllVersionedVersionIds mocks base method
func (m *MockTargetsReader) GetAllVersionedVersionIds() []int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllVersionedVersionIds")
	ret0, _ := ret[0].([]int64)
	return ret0
}

// GetAllVersionedVersionIds indicates an expected call of GetAllVersionedVersionIds
func (mr *MockTargetsReaderMockRecorder) GetAllVersionedVersionIds() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllVersionedVersionIds", reflect.TypeOf((*MockTargetsReader)(nil).GetAllVersionedVersionIds))
}

// GetCurrentVersionID mocks base method
func (m *MockTargetsReader) GetCurrentVersionID() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentVersionID")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetCurrentVersionID indicates an expected call of GetCurrentVersionID
func (mr *MockTargetsReaderMockRecorder) GetCurrentVersionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentVersionID", reflect.TypeOf((*MockTargetsReader)(nil).GetCurrentVersionID))
}

// GetTasksByTargetIds mocks base method
func (m *MockTargetsReader) GetTasksByTargetIds(arg0 []int64) ([]pipeline.TaskName, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTasksByTargetIds", arg0)
	ret0, _ := ret[0].([]pipeline.TaskName)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTasksByTargetIds indicates an expected call of GetTasksByTargetIds
func (mr *MockTargetsReaderMockRecorder) GetTasksByTargetIds(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasksByTargetIds", reflect.TypeOf((*MockTargetsReader)(nil).GetTasksByTargetIds), arg0)
}

// GetTasksByVersionIds mocks base method
func (m *MockTargetsReader) GetTasksByVersionIds(arg0 []int64) ([]pipeline.TaskName, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTasksByVersionIds", arg0)
	ret0, _ := ret[0].([]pipeline.TaskName)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTasksByVersionIds indicates an expected call of GetTasksByVersionIds
func (mr *MockTargetsReaderMockRecorder) GetTasksByVersionIds(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasksByVersionIds", reflect.TypeOf((*MockTargetsReader)(nil).GetTasksByVersionIds), arg0)
}

// MockSystemEventCreatorStore is a mock of SystemEventCreatorStore interface
type MockSystemEventCreatorStore struct {
	ctrl     *gomock.Controller
	recorder *MockSystemEventCreatorStoreMockRecorder
}

// MockSystemEventCreatorStoreMockRecorder is the mock recorder for MockSystemEventCreatorStore
type MockSystemEventCreatorStoreMockRecorder struct {
	mock *MockSystemEventCreatorStore
}

// NewMockSystemEventCreatorStore creates a new mock instance
func NewMockSystemEventCreatorStore(ctrl *gomock.Controller) *MockSystemEventCreatorStore {
	mock := &MockSystemEventCreatorStore{ctrl: ctrl}
	mock.recorder = &MockSystemEventCreatorStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSystemEventCreatorStore) EXPECT() *MockSystemEventCreatorStoreMockRecorder {
	return m.recorder
}

// FindByHeight mocks base method
func (m *MockSystemEventCreatorStore) FindByHeight(arg0 int64) ([]model.ValidatorSeq, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByHeight", arg0)
	ret0, _ := ret[0].([]model.ValidatorSeq)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByHeight indicates an expected call of FindByHeight
func (mr *MockSystemEventCreatorStoreMockRecorder) FindByHeight(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByHeight", reflect.TypeOf((*MockSystemEventCreatorStore)(nil).FindByHeight), arg0)
}

// FindLastByAddress mocks base method
func (m *MockSystemEventCreatorStore) FindLastByAddress(arg0 string, arg1 int64) ([]model.ValidatorSeq, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindLastByAddress", arg0, arg1)
	ret0, _ := ret[0].([]model.ValidatorSeq)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindLastByAddress indicates an expected call of FindLastByAddress
func (mr *MockSystemEventCreatorStoreMockRecorder) FindLastByAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindLastByAddress", reflect.TypeOf((*MockSystemEventCreatorStore)(nil).FindLastByAddress), arg0, arg1)
}
