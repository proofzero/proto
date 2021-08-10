// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/v1alpha1/kubelt_grpc.pb.go

// Package v1alpha1 is a generated GoMock package.
package v1alpha1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/proofzero/proto/pkg/v1alpha1"
	grpc "google.golang.org/grpc"
)

// MockKubeltClient is a mock of KubeltClient interface.
type MockKubeltClient struct {
	ctrl     *gomock.Controller
	recorder *MockKubeltClientMockRecorder
}

// MockKubeltClientMockRecorder is the mock recorder for MockKubeltClient.
type MockKubeltClientMockRecorder struct {
	mock *MockKubeltClient
}

// NewMockKubeltClient creates a new mock instance.
func NewMockKubeltClient(ctrl *gomock.Controller) *MockKubeltClient {
	mock := &MockKubeltClient{ctrl: ctrl}
	mock.recorder = &MockKubeltClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubeltClient) EXPECT() *MockKubeltClientMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockKubeltClient) Apply(ctx context.Context, in *v1alpha1.ApplyRequest, opts ...grpc.CallOption) (*v1alpha1.ApplyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Apply", varargs...)
	ret0, _ := ret[0].(*v1alpha1.ApplyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockKubeltClientMockRecorder) Apply(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockKubeltClient)(nil).Apply), varargs...)
}

// Commit mocks base method.
func (m *MockKubeltClient) Commit(ctx context.Context, in *v1alpha1.CommitRequest, opts ...grpc.CallOption) (*v1alpha1.CommitResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Commit", varargs...)
	ret0, _ := ret[0].(*v1alpha1.CommitResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Commit indicates an expected call of Commit.
func (mr *MockKubeltClientMockRecorder) Commit(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockKubeltClient)(nil).Commit), varargs...)
}

// HealthCheck mocks base method.
func (m *MockKubeltClient) HealthCheck(ctx context.Context, in *v1alpha1.HealthCheckRequest, opts ...grpc.CallOption) (*v1alpha1.HealthCheckResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HealthCheck", varargs...)
	ret0, _ := ret[0].(*v1alpha1.HealthCheckResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockKubeltClientMockRecorder) HealthCheck(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockKubeltClient)(nil).HealthCheck), varargs...)
}

// MockKubeltServer is a mock of KubeltServer interface.
type MockKubeltServer struct {
	ctrl     *gomock.Controller
	recorder *MockKubeltServerMockRecorder
}

// MockKubeltServerMockRecorder is the mock recorder for MockKubeltServer.
type MockKubeltServerMockRecorder struct {
	mock *MockKubeltServer
}

// NewMockKubeltServer creates a new mock instance.
func NewMockKubeltServer(ctrl *gomock.Controller) *MockKubeltServer {
	mock := &MockKubeltServer{ctrl: ctrl}
	mock.recorder = &MockKubeltServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubeltServer) EXPECT() *MockKubeltServerMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockKubeltServer) Apply(arg0 context.Context, arg1 *v1alpha1.ApplyRequest) (*v1alpha1.ApplyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha1.ApplyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockKubeltServerMockRecorder) Apply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockKubeltServer)(nil).Apply), arg0, arg1)
}

// Commit mocks base method.
func (m *MockKubeltServer) Commit(arg0 context.Context, arg1 *v1alpha1.CommitRequest) (*v1alpha1.CommitResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha1.CommitResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Commit indicates an expected call of Commit.
func (mr *MockKubeltServerMockRecorder) Commit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockKubeltServer)(nil).Commit), arg0, arg1)
}

// HealthCheck mocks base method.
func (m *MockKubeltServer) HealthCheck(arg0 context.Context, arg1 *v1alpha1.HealthCheckRequest) (*v1alpha1.HealthCheckResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha1.HealthCheckResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockKubeltServerMockRecorder) HealthCheck(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockKubeltServer)(nil).HealthCheck), arg0, arg1)
}

// mustEmbedUnimplementedKubeltServer mocks base method.
func (m *MockKubeltServer) mustEmbedUnimplementedKubeltServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedKubeltServer")
}

// mustEmbedUnimplementedKubeltServer indicates an expected call of mustEmbedUnimplementedKubeltServer.
func (mr *MockKubeltServerMockRecorder) mustEmbedUnimplementedKubeltServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedKubeltServer", reflect.TypeOf((*MockKubeltServer)(nil).mustEmbedUnimplementedKubeltServer))
}

// MockUnsafeKubeltServer is a mock of UnsafeKubeltServer interface.
type MockUnsafeKubeltServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeKubeltServerMockRecorder
}

// MockUnsafeKubeltServerMockRecorder is the mock recorder for MockUnsafeKubeltServer.
type MockUnsafeKubeltServerMockRecorder struct {
	mock *MockUnsafeKubeltServer
}

// NewMockUnsafeKubeltServer creates a new mock instance.
func NewMockUnsafeKubeltServer(ctrl *gomock.Controller) *MockUnsafeKubeltServer {
	mock := &MockUnsafeKubeltServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeKubeltServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeKubeltServer) EXPECT() *MockUnsafeKubeltServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedKubeltServer mocks base method.
func (m *MockUnsafeKubeltServer) mustEmbedUnimplementedKubeltServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedKubeltServer")
}

// mustEmbedUnimplementedKubeltServer indicates an expected call of mustEmbedUnimplementedKubeltServer.
func (mr *MockUnsafeKubeltServerMockRecorder) mustEmbedUnimplementedKubeltServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedKubeltServer", reflect.TypeOf((*MockUnsafeKubeltServer)(nil).mustEmbedUnimplementedKubeltServer))
}
