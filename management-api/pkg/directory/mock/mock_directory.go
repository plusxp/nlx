// Code generated by MockGen. DO NOT EDIT.
// Source: directory.pb.go

// Package mock_directory is a generated GoMock package.
package mock_directory

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	directory "go.nlx.io/nlx/management-api/pkg/directory"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockDirectoryClient is a mock of DirectoryClient interface
type MockDirectoryClient struct {
	ctrl     *gomock.Controller
	recorder *MockDirectoryClientMockRecorder
}

// MockDirectoryClientMockRecorder is the mock recorder for MockDirectoryClient
type MockDirectoryClientMockRecorder struct {
	mock *MockDirectoryClient
}

// NewMockDirectoryClient creates a new mock instance
func NewMockDirectoryClient(ctrl *gomock.Controller) *MockDirectoryClient {
	mock := &MockDirectoryClient{ctrl: ctrl}
	mock.recorder = &MockDirectoryClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDirectoryClient) EXPECT() *MockDirectoryClientMockRecorder {
	return m.recorder
}

// ListServices mocks base method
func (m *MockDirectoryClient) ListServices(ctx context.Context, in *directory.Empty, opts ...grpc.CallOption) (*directory.ListServicesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListServices", varargs...)
	ret0, _ := ret[0].(*directory.ListServicesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices
func (mr *MockDirectoryClientMockRecorder) ListServices(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockDirectoryClient)(nil).ListServices), varargs...)
}

// GetOrganizationService mocks base method
func (m *MockDirectoryClient) GetOrganizationService(ctx context.Context, in *directory.GetOrganizationServiceRequest, opts ...grpc.CallOption) (*directory.DirectoryService, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOrganizationService", varargs...)
	ret0, _ := ret[0].(*directory.DirectoryService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrganizationService indicates an expected call of GetOrganizationService
func (mr *MockDirectoryClientMockRecorder) GetOrganizationService(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrganizationService", reflect.TypeOf((*MockDirectoryClient)(nil).GetOrganizationService), varargs...)
}

// RequestAccessToService mocks base method
func (m *MockDirectoryClient) RequestAccessToService(ctx context.Context, in *directory.RequestAccessToServiceRequest, opts ...grpc.CallOption) (*directory.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RequestAccessToService", varargs...)
	ret0, _ := ret[0].(*directory.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestAccessToService indicates an expected call of RequestAccessToService
func (mr *MockDirectoryClientMockRecorder) RequestAccessToService(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestAccessToService", reflect.TypeOf((*MockDirectoryClient)(nil).RequestAccessToService), varargs...)
}

// MockDirectoryServer is a mock of DirectoryServer interface
type MockDirectoryServer struct {
	ctrl     *gomock.Controller
	recorder *MockDirectoryServerMockRecorder
}

// MockDirectoryServerMockRecorder is the mock recorder for MockDirectoryServer
type MockDirectoryServerMockRecorder struct {
	mock *MockDirectoryServer
}

// NewMockDirectoryServer creates a new mock instance
func NewMockDirectoryServer(ctrl *gomock.Controller) *MockDirectoryServer {
	mock := &MockDirectoryServer{ctrl: ctrl}
	mock.recorder = &MockDirectoryServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDirectoryServer) EXPECT() *MockDirectoryServerMockRecorder {
	return m.recorder
}

// ListServices mocks base method
func (m *MockDirectoryServer) ListServices(arg0 context.Context, arg1 *directory.Empty) (*directory.ListServicesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServices", arg0, arg1)
	ret0, _ := ret[0].(*directory.ListServicesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices
func (mr *MockDirectoryServerMockRecorder) ListServices(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockDirectoryServer)(nil).ListServices), arg0, arg1)
}

// GetOrganizationService mocks base method
func (m *MockDirectoryServer) GetOrganizationService(arg0 context.Context, arg1 *directory.GetOrganizationServiceRequest) (*directory.DirectoryService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrganizationService", arg0, arg1)
	ret0, _ := ret[0].(*directory.DirectoryService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrganizationService indicates an expected call of GetOrganizationService
func (mr *MockDirectoryServerMockRecorder) GetOrganizationService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrganizationService", reflect.TypeOf((*MockDirectoryServer)(nil).GetOrganizationService), arg0, arg1)
}

// RequestAccessToService mocks base method
func (m *MockDirectoryServer) RequestAccessToService(arg0 context.Context, arg1 *directory.RequestAccessToServiceRequest) (*directory.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestAccessToService", arg0, arg1)
	ret0, _ := ret[0].(*directory.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestAccessToService indicates an expected call of RequestAccessToService
func (mr *MockDirectoryServerMockRecorder) RequestAccessToService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestAccessToService", reflect.TypeOf((*MockDirectoryServer)(nil).RequestAccessToService), arg0, arg1)
}
