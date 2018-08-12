// Code generated by MockGen. DO NOT EDIT.
// Source: adapter.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	proto "github.com/kucherenkovova/marco-polo/proto"
	context "golang.org/x/net/context"
	reflect "reflect"
)

// MockForwarder is a mock of Forwarder interface
type MockForwarder struct {
	ctrl     *gomock.Controller
	recorder *MockForwarderMockRecorder
}

// MockForwarderMockRecorder is the mock recorder for MockForwarder
type MockForwarderMockRecorder struct {
	mock *MockForwarder
}

// NewMockForwarder creates a new mock instance
func NewMockForwarder(ctrl *gomock.Controller) *MockForwarder {
	mock := &MockForwarder{ctrl: ctrl}
	mock.recorder = &MockForwarderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockForwarder) EXPECT() *MockForwarderMockRecorder {
	return m.recorder
}

// Forward mocks base method
func (m *MockForwarder) Forward(arg0 context.Context, arg1 *proto.Phrase) (*proto.Phrase, error) {
	ret := m.ctrl.Call(m, "Forward", arg0, arg1)
	ret0, _ := ret[0].(*proto.Phrase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Forward indicates an expected call of Forward
func (mr *MockForwarderMockRecorder) Forward(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Forward", reflect.TypeOf((*MockForwarder)(nil).Forward), arg0, arg1)
}