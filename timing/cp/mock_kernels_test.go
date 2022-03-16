// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.com/akita/mgpusim/v3/kernels (interfaces: GridBuilder)

package cp

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	kernels "gitlab.com/akita/mgpusim/v3/kernels"
)

// MockGridBuilder is a mock of GridBuilder interface.
type MockGridBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockGridBuilderMockRecorder
}

// MockGridBuilderMockRecorder is the mock recorder for MockGridBuilder.
type MockGridBuilderMockRecorder struct {
	mock *MockGridBuilder
}

// NewMockGridBuilder creates a new mock instance.
func NewMockGridBuilder(ctrl *gomock.Controller) *MockGridBuilder {
	mock := &MockGridBuilder{ctrl: ctrl}
	mock.recorder = &MockGridBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGridBuilder) EXPECT() *MockGridBuilderMockRecorder {
	return m.recorder
}

// NextWG mocks base method.
func (m *MockGridBuilder) NextWG() *kernels.WorkGroup {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextWG")
	ret0, _ := ret[0].(*kernels.WorkGroup)
	return ret0
}

// NextWG indicates an expected call of NextWG.
func (mr *MockGridBuilderMockRecorder) NextWG() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextWG", reflect.TypeOf((*MockGridBuilder)(nil).NextWG))
}

// NumWG mocks base method.
func (m *MockGridBuilder) NumWG() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NumWG")
	ret0, _ := ret[0].(int)
	return ret0
}

// NumWG indicates an expected call of NumWG.
func (mr *MockGridBuilderMockRecorder) NumWG() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NumWG", reflect.TypeOf((*MockGridBuilder)(nil).NumWG))
}

// SetKernel mocks base method.
func (m *MockGridBuilder) SetKernel(arg0 kernels.KernelLaunchInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetKernel", arg0)
}

// SetKernel indicates an expected call of SetKernel.
func (mr *MockGridBuilderMockRecorder) SetKernel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetKernel", reflect.TypeOf((*MockGridBuilder)(nil).SetKernel), arg0)
}

// Skip mocks base method.
func (m *MockGridBuilder) Skip(arg0 int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Skip", arg0)
}

// Skip indicates an expected call of Skip.
func (mr *MockGridBuilderMockRecorder) Skip(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Skip", reflect.TypeOf((*MockGridBuilder)(nil).Skip), arg0)
}
