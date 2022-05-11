// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ava-labs/avalanchego/snow/networking/tracker (interfaces: CPUTargeter)

// Package tracker is a generated GoMock package.
package tracker

import (
	ids "github.com/ava-labs/avalanchego/ids"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCPUTargeter is a mock of CPUTargeter interface
type MockCPUTargeter struct {
	ctrl     *gomock.Controller
	recorder *MockCPUTargeterMockRecorder
}

// MockCPUTargeterMockRecorder is the mock recorder for MockCPUTargeter
type MockCPUTargeterMockRecorder struct {
	mock *MockCPUTargeter
}

// NewMockCPUTargeter creates a new mock instance
func NewMockCPUTargeter(ctrl *gomock.Controller) *MockCPUTargeter {
	mock := &MockCPUTargeter{ctrl: ctrl}
	mock.recorder = &MockCPUTargeterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCPUTargeter) EXPECT() *MockCPUTargeterMockRecorder {
	return m.recorder
}

// TargetCPUUsage mocks base method
func (m *MockCPUTargeter) TargetCPUUsage(arg0 ids.NodeID) (float64, float64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TargetCPUUsage", arg0)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(float64)
	return ret0, ret1
}

// TargetCPUUsage indicates an expected call of TargetCPUUsage
func (mr *MockCPUTargeterMockRecorder) TargetCPUUsage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TargetCPUUsage", reflect.TypeOf((*MockCPUTargeter)(nil).TargetCPUUsage), arg0)
}
