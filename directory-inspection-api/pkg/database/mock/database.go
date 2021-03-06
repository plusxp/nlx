// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/database/database.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	nlxversion "go.nlx.io/nlx/common/nlxversion"
	database "go.nlx.io/nlx/directory-inspection-api/pkg/database"
	reflect "reflect"
)

// MockDirectoryDatabase is a mock of DirectoryDatabase interface
type MockDirectoryDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDirectoryDatabaseMockRecorder
}

// MockDirectoryDatabaseMockRecorder is the mock recorder for MockDirectoryDatabase
type MockDirectoryDatabaseMockRecorder struct {
	mock *MockDirectoryDatabase
}

// NewMockDirectoryDatabase creates a new mock instance
func NewMockDirectoryDatabase(ctrl *gomock.Controller) *MockDirectoryDatabase {
	mock := &MockDirectoryDatabase{ctrl: ctrl}
	mock.recorder = &MockDirectoryDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDirectoryDatabase) EXPECT() *MockDirectoryDatabaseMockRecorder {
	return m.recorder
}

// ListServices mocks base method
func (m *MockDirectoryDatabase) ListServices(ctx context.Context, organizationName string) ([]*database.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServices", ctx, organizationName)
	ret0, _ := ret[0].([]*database.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices
func (mr *MockDirectoryDatabaseMockRecorder) ListServices(ctx, organizationName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockDirectoryDatabase)(nil).ListServices), ctx, organizationName)
}

// RegisterOutwayVersion mocks base method
func (m *MockDirectoryDatabase) RegisterOutwayVersion(ctx context.Context, version nlxversion.Version) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterOutwayVersion", ctx, version)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterOutwayVersion indicates an expected call of RegisterOutwayVersion
func (mr *MockDirectoryDatabaseMockRecorder) RegisterOutwayVersion(ctx, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterOutwayVersion", reflect.TypeOf((*MockDirectoryDatabase)(nil).RegisterOutwayVersion), ctx, version)
}

// ListOrganizations mocks base method
func (m *MockDirectoryDatabase) ListOrganizations(ctx context.Context) ([]*database.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOrganizations", ctx)
	ret0, _ := ret[0].([]*database.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrganizations indicates an expected call of ListOrganizations
func (mr *MockDirectoryDatabaseMockRecorder) ListOrganizations(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrganizations", reflect.TypeOf((*MockDirectoryDatabase)(nil).ListOrganizations), ctx)
}

// GetOrganizationInwayAddress mocks base method
func (m *MockDirectoryDatabase) GetOrganizationInwayAddress(ctx context.Context, organizationName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrganizationInwayAddress", ctx, organizationName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrganizationInwayAddress indicates an expected call of GetOrganizationInwayAddress
func (mr *MockDirectoryDatabaseMockRecorder) GetOrganizationInwayAddress(ctx, organizationName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrganizationInwayAddress", reflect.TypeOf((*MockDirectoryDatabase)(nil).GetOrganizationInwayAddress), ctx, organizationName)
}
