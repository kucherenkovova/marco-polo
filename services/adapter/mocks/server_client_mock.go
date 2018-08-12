// Code generated by MockGen. DO NOT EDIT.
// Source: services.pb.go

// Package mock_proto is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	. "github.com/kucherenkovova/marco-polo/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockServerClient is a mock of ServerClient interface
type MockServerClient struct {
	ctrl     *gomock.Controller
	recorder *MockServerClientMockRecorder
}

// MockServerClientMockRecorder is the mock recorder for MockServerClient
type MockServerClientMockRecorder struct {
	mock *MockServerClient
}

// NewMockServerClient creates a new mock instance
func NewMockServerClient(ctrl *gomock.Controller) *MockServerClient {
	mock := &MockServerClient{ctrl: ctrl}
	mock.recorder = &MockServerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServerClient) EXPECT() *MockServerClientMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockServerClient) Send(ctx context.Context, in *Phrase, opts ...grpc.CallOption) (*Phrase, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Send", varargs...)
	ret0, _ := ret[0].(*Phrase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send
func (mr *MockServerClientMockRecorder) Send(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockServerClient)(nil).Send), varargs...)
}

// MockServerServer is a mock of ServerServer interface
type MockServerServer struct {
	ctrl     *gomock.Controller
	recorder *MockServerServerMockRecorder
}

// MockServerServerMockRecorder is the mock recorder for MockServerServer
type MockServerServerMockRecorder struct {
	mock *MockServerServer
}

// NewMockServerServer creates a new mock instance
func NewMockServerServer(ctrl *gomock.Controller) *MockServerServer {
	mock := &MockServerServer{ctrl: ctrl}
	mock.recorder = &MockServerServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServerServer) EXPECT() *MockServerServerMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockServerServer) Send(arg0 context.Context, arg1 *Phrase) (*Phrase, error) {
	ret := m.ctrl.Call(m, "Send", arg0, arg1)
	ret0, _ := ret[0].(*Phrase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send
func (mr *MockServerServerMockRecorder) Send(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockServerServer)(nil).Send), arg0, arg1)
}

// MockAdapterClient is a mock of AdapterClient interface
type MockAdapterClient struct {
	ctrl     *gomock.Controller
	recorder *MockAdapterClientMockRecorder
}

// MockAdapterClientMockRecorder is the mock recorder for MockAdapterClient
type MockAdapterClientMockRecorder struct {
	mock *MockAdapterClient
}

// NewMockAdapterClient creates a new mock instance
func NewMockAdapterClient(ctrl *gomock.Controller) *MockAdapterClient {
	mock := &MockAdapterClient{ctrl: ctrl}
	mock.recorder = &MockAdapterClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAdapterClient) EXPECT() *MockAdapterClientMockRecorder {
	return m.recorder
}

// Forward mocks base method
func (m *MockAdapterClient) Forward(ctx context.Context, in *Phrase, opts ...grpc.CallOption) (*Phrase, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Forward", varargs...)
	ret0, _ := ret[0].(*Phrase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Forward indicates an expected call of Forward
func (mr *MockAdapterClientMockRecorder) Forward(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Forward", reflect.TypeOf((*MockAdapterClient)(nil).Forward), varargs...)
}

// MockAdapterServer is a mock of AdapterServer interface
type MockAdapterServer struct {
	ctrl     *gomock.Controller
	recorder *MockAdapterServerMockRecorder
}

// MockAdapterServerMockRecorder is the mock recorder for MockAdapterServer
type MockAdapterServerMockRecorder struct {
	mock *MockAdapterServer
}

// NewMockAdapterServer creates a new mock instance
func NewMockAdapterServer(ctrl *gomock.Controller) *MockAdapterServer {
	mock := &MockAdapterServer{ctrl: ctrl}
	mock.recorder = &MockAdapterServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAdapterServer) EXPECT() *MockAdapterServerMockRecorder {
	return m.recorder
}

// Forward mocks base method
func (m *MockAdapterServer) Forward(arg0 context.Context, arg1 *Phrase) (*Phrase, error) {
	ret := m.ctrl.Call(m, "Forward", arg0, arg1)
	ret0, _ := ret[0].(*Phrase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Forward indicates an expected call of Forward
func (mr *MockAdapterServerMockRecorder) Forward(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Forward", reflect.TypeOf((*MockAdapterServer)(nil).Forward), arg0, arg1)
}
