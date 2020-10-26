// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/database/database.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	database "go.nlx.io/nlx/management-api/pkg/database"
	reflect "reflect"
	time "time"
)

// MockConfigDatabase is a mock of ConfigDatabase interface
type MockConfigDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockConfigDatabaseMockRecorder
}

// MockConfigDatabaseMockRecorder is the mock recorder for MockConfigDatabase
type MockConfigDatabaseMockRecorder struct {
	mock *MockConfigDatabase
}

// NewMockConfigDatabase creates a new mock instance
func NewMockConfigDatabase(ctrl *gomock.Controller) *MockConfigDatabase {
	mock := &MockConfigDatabase{ctrl: ctrl}
	mock.recorder = &MockConfigDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigDatabase) EXPECT() *MockConfigDatabaseMockRecorder {
	return m.recorder
}

// ListServices mocks base method
func (m *MockConfigDatabase) ListServices(ctx context.Context) ([]*database.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServices", ctx)
	ret0, _ := ret[0].([]*database.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices
func (mr *MockConfigDatabaseMockRecorder) ListServices(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockConfigDatabase)(nil).ListServices), ctx)
}

// GetService mocks base method
func (m *MockConfigDatabase) GetService(ctx context.Context, name string) (*database.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", ctx, name)
	ret0, _ := ret[0].(*database.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService
func (mr *MockConfigDatabaseMockRecorder) GetService(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockConfigDatabase)(nil).GetService), ctx, name)
}

// CreateService mocks base method
func (m *MockConfigDatabase) CreateService(ctx context.Context, service *database.Service) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateService", ctx, service)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateService indicates an expected call of CreateService
func (mr *MockConfigDatabaseMockRecorder) CreateService(ctx, service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateService", reflect.TypeOf((*MockConfigDatabase)(nil).CreateService), ctx, service)
}

// UpdateService mocks base method
func (m *MockConfigDatabase) UpdateService(ctx context.Context, name string, service *database.Service) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateService", ctx, name, service)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateService indicates an expected call of UpdateService
func (mr *MockConfigDatabaseMockRecorder) UpdateService(ctx, name, service interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateService", reflect.TypeOf((*MockConfigDatabase)(nil).UpdateService), ctx, name, service)
}

// DeleteService mocks base method
func (m *MockConfigDatabase) DeleteService(ctx context.Context, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteService", ctx, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteService indicates an expected call of DeleteService
func (mr *MockConfigDatabaseMockRecorder) DeleteService(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockConfigDatabase)(nil).DeleteService), ctx, name)
}

// ListInways mocks base method
func (m *MockConfigDatabase) ListInways(ctx context.Context) ([]*database.Inway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInways", ctx)
	ret0, _ := ret[0].([]*database.Inway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInways indicates an expected call of ListInways
func (mr *MockConfigDatabaseMockRecorder) ListInways(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInways", reflect.TypeOf((*MockConfigDatabase)(nil).ListInways), ctx)
}

// GetInway mocks base method
func (m *MockConfigDatabase) GetInway(ctx context.Context, name string) (*database.Inway, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInway", ctx, name)
	ret0, _ := ret[0].(*database.Inway)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInway indicates an expected call of GetInway
func (mr *MockConfigDatabaseMockRecorder) GetInway(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInway", reflect.TypeOf((*MockConfigDatabase)(nil).GetInway), ctx, name)
}

// CreateInway mocks base method
func (m *MockConfigDatabase) CreateInway(ctx context.Context, inway *database.Inway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInway", ctx, inway)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateInway indicates an expected call of CreateInway
func (mr *MockConfigDatabaseMockRecorder) CreateInway(ctx, inway interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInway", reflect.TypeOf((*MockConfigDatabase)(nil).CreateInway), ctx, inway)
}

// UpdateInway mocks base method
func (m *MockConfigDatabase) UpdateInway(ctx context.Context, name string, inway *database.Inway) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInway", ctx, name, inway)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInway indicates an expected call of UpdateInway
func (mr *MockConfigDatabaseMockRecorder) UpdateInway(ctx, name, inway interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInway", reflect.TypeOf((*MockConfigDatabase)(nil).UpdateInway), ctx, name, inway)
}

// DeleteInway mocks base method
func (m *MockConfigDatabase) DeleteInway(ctx context.Context, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInway", ctx, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInway indicates an expected call of DeleteInway
func (mr *MockConfigDatabaseMockRecorder) DeleteInway(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInway", reflect.TypeOf((*MockConfigDatabase)(nil).DeleteInway), ctx, name)
}

// PutInsightConfiguration mocks base method
func (m *MockConfigDatabase) PutInsightConfiguration(ctx context.Context, configuration *database.InsightConfiguration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutInsightConfiguration", ctx, configuration)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutInsightConfiguration indicates an expected call of PutInsightConfiguration
func (mr *MockConfigDatabaseMockRecorder) PutInsightConfiguration(ctx, configuration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutInsightConfiguration", reflect.TypeOf((*MockConfigDatabase)(nil).PutInsightConfiguration), ctx, configuration)
}

// GetInsightConfiguration mocks base method
func (m *MockConfigDatabase) GetInsightConfiguration(ctx context.Context) (*database.InsightConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInsightConfiguration", ctx)
	ret0, _ := ret[0].(*database.InsightConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInsightConfiguration indicates an expected call of GetInsightConfiguration
func (mr *MockConfigDatabaseMockRecorder) GetInsightConfiguration(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInsightConfiguration", reflect.TypeOf((*MockConfigDatabase)(nil).GetInsightConfiguration), ctx)
}

// ListAllOutgoingAccessRequests mocks base method
func (m *MockConfigDatabase) ListAllOutgoingAccessRequests(ctx context.Context) ([]*database.OutgoingAccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllOutgoingAccessRequests", ctx)
	ret0, _ := ret[0].([]*database.OutgoingAccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllOutgoingAccessRequests indicates an expected call of ListAllOutgoingAccessRequests
func (mr *MockConfigDatabaseMockRecorder) ListAllOutgoingAccessRequests(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllOutgoingAccessRequests", reflect.TypeOf((*MockConfigDatabase)(nil).ListAllOutgoingAccessRequests), ctx)
}

// ListOutgoingAccessRequests mocks base method
func (m *MockConfigDatabase) ListOutgoingAccessRequests(ctx context.Context, organizationName, serviceName string) ([]*database.OutgoingAccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOutgoingAccessRequests", ctx, organizationName, serviceName)
	ret0, _ := ret[0].([]*database.OutgoingAccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOutgoingAccessRequests indicates an expected call of ListOutgoingAccessRequests
func (mr *MockConfigDatabaseMockRecorder) ListOutgoingAccessRequests(ctx, organizationName, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOutgoingAccessRequests", reflect.TypeOf((*MockConfigDatabase)(nil).ListOutgoingAccessRequests), ctx, organizationName, serviceName)
}

// GetOutgoingAccessRequest mocks base method
func (m *MockConfigDatabase) GetOutgoingAccessRequest(ctx context.Context, id string) (*database.OutgoingAccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOutgoingAccessRequest", ctx, id)
	ret0, _ := ret[0].(*database.OutgoingAccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOutgoingAccessRequest indicates an expected call of GetOutgoingAccessRequest
func (mr *MockConfigDatabaseMockRecorder) GetOutgoingAccessRequest(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutgoingAccessRequest", reflect.TypeOf((*MockConfigDatabase)(nil).GetOutgoingAccessRequest), ctx, id)
}

// GetLatestOutgoingAccessRequest mocks base method
func (m *MockConfigDatabase) GetLatestOutgoingAccessRequest(ctx context.Context, organizationName, serviceName string) (*database.OutgoingAccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestOutgoingAccessRequest", ctx, organizationName, serviceName)
	ret0, _ := ret[0].(*database.OutgoingAccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestOutgoingAccessRequest indicates an expected call of GetLatestOutgoingAccessRequest
func (mr *MockConfigDatabaseMockRecorder) GetLatestOutgoingAccessRequest(ctx, organizationName, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestOutgoingAccessRequest", reflect.TypeOf((*MockConfigDatabase)(nil).GetLatestOutgoingAccessRequest), ctx, organizationName, serviceName)
}

// ListAllLatestOutgoingAccessRequests mocks base method
func (m *MockConfigDatabase) ListAllLatestOutgoingAccessRequests(ctx context.Context) (map[string]*database.OutgoingAccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllLatestOutgoingAccessRequests", ctx)
	ret0, _ := ret[0].(map[string]*database.OutgoingAccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllLatestOutgoingAccessRequests indicates an expected call of ListAllLatestOutgoingAccessRequests
func (mr *MockConfigDatabaseMockRecorder) ListAllLatestOutgoingAccessRequests(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllLatestOutgoingAccessRequests", reflect.TypeOf((*MockConfigDatabase)(nil).ListAllLatestOutgoingAccessRequests), ctx)
}

// LockOutgoingAccessRequest mocks base method
func (m *MockConfigDatabase) LockOutgoingAccessRequest(ctx context.Context, accessRequest *database.OutgoingAccessRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockOutgoingAccessRequest", ctx, accessRequest)
	ret0, _ := ret[0].(error)
	return ret0
}

// LockOutgoingAccessRequest indicates an expected call of LockOutgoingAccessRequest
func (mr *MockConfigDatabaseMockRecorder) LockOutgoingAccessRequest(ctx, accessRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockOutgoingAccessRequest", reflect.TypeOf((*MockConfigDatabase)(nil).LockOutgoingAccessRequest), ctx, accessRequest)
}

// UnlockOutgoingAccessRequest mocks base method
func (m *MockConfigDatabase) UnlockOutgoingAccessRequest(ctx context.Context, accessRequest *database.OutgoingAccessRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnlockOutgoingAccessRequest", ctx, accessRequest)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnlockOutgoingAccessRequest indicates an expected call of UnlockOutgoingAccessRequest
func (mr *MockConfigDatabaseMockRecorder) UnlockOutgoingAccessRequest(ctx, accessRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlockOutgoingAccessRequest", reflect.TypeOf((*MockConfigDatabase)(nil).UnlockOutgoingAccessRequest), ctx, accessRequest)
}

// CreateOutgoingAccessRequest mocks base method
func (m *MockConfigDatabase) CreateOutgoingAccessRequest(ctx context.Context, accessRequest *database.OutgoingAccessRequest) (*database.OutgoingAccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOutgoingAccessRequest", ctx, accessRequest)
	ret0, _ := ret[0].(*database.OutgoingAccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOutgoingAccessRequest indicates an expected call of CreateOutgoingAccessRequest
func (mr *MockConfigDatabaseMockRecorder) CreateOutgoingAccessRequest(ctx, accessRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOutgoingAccessRequest", reflect.TypeOf((*MockConfigDatabase)(nil).CreateOutgoingAccessRequest), ctx, accessRequest)
}

// UpdateOutgoingAccessRequestState mocks base method
func (m *MockConfigDatabase) UpdateOutgoingAccessRequestState(ctx context.Context, accessRequest *database.OutgoingAccessRequest, state database.AccessRequestState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOutgoingAccessRequestState", ctx, accessRequest, state)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOutgoingAccessRequestState indicates an expected call of UpdateOutgoingAccessRequestState
func (mr *MockConfigDatabaseMockRecorder) UpdateOutgoingAccessRequestState(ctx, accessRequest, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOutgoingAccessRequestState", reflect.TypeOf((*MockConfigDatabase)(nil).UpdateOutgoingAccessRequestState), ctx, accessRequest, state)
}

// WatchOutgoingAccessRequests mocks base method
func (m *MockConfigDatabase) WatchOutgoingAccessRequests(ctx context.Context, output chan *database.OutgoingAccessRequest) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WatchOutgoingAccessRequests", ctx, output)
}

// WatchOutgoingAccessRequests indicates an expected call of WatchOutgoingAccessRequests
func (mr *MockConfigDatabaseMockRecorder) WatchOutgoingAccessRequests(ctx, output interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchOutgoingAccessRequests", reflect.TypeOf((*MockConfigDatabase)(nil).WatchOutgoingAccessRequests), ctx, output)
}

// ListAllIncomingAccessRequests mocks base method
func (m *MockConfigDatabase) ListAllIncomingAccessRequests(ctx context.Context) ([]*database.IncomingAccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllIncomingAccessRequests", ctx)
	ret0, _ := ret[0].([]*database.IncomingAccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllIncomingAccessRequests indicates an expected call of ListAllIncomingAccessRequests
func (mr *MockConfigDatabaseMockRecorder) ListAllIncomingAccessRequests(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllIncomingAccessRequests", reflect.TypeOf((*MockConfigDatabase)(nil).ListAllIncomingAccessRequests), ctx)
}

// ListIncomingAccessRequests mocks base method
func (m *MockConfigDatabase) ListIncomingAccessRequests(ctx context.Context, organizationName, serviceName string) ([]*database.IncomingAccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListIncomingAccessRequests", ctx, organizationName, serviceName)
	ret0, _ := ret[0].([]*database.IncomingAccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIncomingAccessRequests indicates an expected call of ListIncomingAccessRequests
func (mr *MockConfigDatabaseMockRecorder) ListIncomingAccessRequests(ctx, organizationName, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIncomingAccessRequests", reflect.TypeOf((*MockConfigDatabase)(nil).ListIncomingAccessRequests), ctx, organizationName, serviceName)
}

// GetLatestIncomingAccessRequest mocks base method
func (m *MockConfigDatabase) GetLatestIncomingAccessRequest(ctx context.Context, organizationName, serviceName string) (*database.IncomingAccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestIncomingAccessRequest", ctx, organizationName, serviceName)
	ret0, _ := ret[0].(*database.IncomingAccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestIncomingAccessRequest indicates an expected call of GetLatestIncomingAccessRequest
func (mr *MockConfigDatabaseMockRecorder) GetLatestIncomingAccessRequest(ctx, organizationName, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestIncomingAccessRequest", reflect.TypeOf((*MockConfigDatabase)(nil).GetLatestIncomingAccessRequest), ctx, organizationName, serviceName)
}

// ListAllLatestIncomingAccessRequests mocks base method
func (m *MockConfigDatabase) ListAllLatestIncomingAccessRequests(ctx context.Context) (map[string]*database.IncomingAccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllLatestIncomingAccessRequests", ctx)
	ret0, _ := ret[0].(map[string]*database.IncomingAccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllLatestIncomingAccessRequests indicates an expected call of ListAllLatestIncomingAccessRequests
func (mr *MockConfigDatabaseMockRecorder) ListAllLatestIncomingAccessRequests(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllLatestIncomingAccessRequests", reflect.TypeOf((*MockConfigDatabase)(nil).ListAllLatestIncomingAccessRequests), ctx)
}

// GetIncomingAccessRequest mocks base method
func (m *MockConfigDatabase) GetIncomingAccessRequest(ctx context.Context, id string) (*database.IncomingAccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIncomingAccessRequest", ctx, id)
	ret0, _ := ret[0].(*database.IncomingAccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIncomingAccessRequest indicates an expected call of GetIncomingAccessRequest
func (mr *MockConfigDatabaseMockRecorder) GetIncomingAccessRequest(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIncomingAccessRequest", reflect.TypeOf((*MockConfigDatabase)(nil).GetIncomingAccessRequest), ctx, id)
}

// CreateIncomingAccessRequest mocks base method
func (m *MockConfigDatabase) CreateIncomingAccessRequest(ctx context.Context, accessRequest *database.IncomingAccessRequest) (*database.IncomingAccessRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIncomingAccessRequest", ctx, accessRequest)
	ret0, _ := ret[0].(*database.IncomingAccessRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIncomingAccessRequest indicates an expected call of CreateIncomingAccessRequest
func (mr *MockConfigDatabaseMockRecorder) CreateIncomingAccessRequest(ctx, accessRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIncomingAccessRequest", reflect.TypeOf((*MockConfigDatabase)(nil).CreateIncomingAccessRequest), ctx, accessRequest)
}

// UpdateIncomingAccessRequestState mocks base method
func (m *MockConfigDatabase) UpdateIncomingAccessRequestState(ctx context.Context, accessRequest *database.IncomingAccessRequest, state database.AccessRequestState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIncomingAccessRequestState", ctx, accessRequest, state)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateIncomingAccessRequestState indicates an expected call of UpdateIncomingAccessRequestState
func (mr *MockConfigDatabaseMockRecorder) UpdateIncomingAccessRequestState(ctx, accessRequest, state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIncomingAccessRequestState", reflect.TypeOf((*MockConfigDatabase)(nil).UpdateIncomingAccessRequestState), ctx, accessRequest, state)
}

// CreateAccessGrant mocks base method
func (m *MockConfigDatabase) CreateAccessGrant(ctx context.Context, accessRequest *database.IncomingAccessRequest) (*database.AccessGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessGrant", ctx, accessRequest)
	ret0, _ := ret[0].(*database.AccessGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccessGrant indicates an expected call of CreateAccessGrant
func (mr *MockConfigDatabaseMockRecorder) CreateAccessGrant(ctx, accessRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessGrant", reflect.TypeOf((*MockConfigDatabase)(nil).CreateAccessGrant), ctx, accessRequest)
}

// RevokeAccessGrant mocks base method
func (m *MockConfigDatabase) RevokeAccessGrant(ctx context.Context, serviceName, organizationName, id string) (*database.AccessGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeAccessGrant", ctx, serviceName, organizationName, id)
	ret0, _ := ret[0].(*database.AccessGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeAccessGrant indicates an expected call of RevokeAccessGrant
func (mr *MockConfigDatabaseMockRecorder) RevokeAccessGrant(ctx, serviceName, organizationName, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeAccessGrant", reflect.TypeOf((*MockConfigDatabase)(nil).RevokeAccessGrant), ctx, serviceName, organizationName, id)
}

// ListAccessGrantsForService mocks base method
func (m *MockConfigDatabase) ListAccessGrantsForService(ctx context.Context, serviceName string) ([]*database.AccessGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccessGrantsForService", ctx, serviceName)
	ret0, _ := ret[0].([]*database.AccessGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccessGrantsForService indicates an expected call of ListAccessGrantsForService
func (mr *MockConfigDatabaseMockRecorder) ListAccessGrantsForService(ctx, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccessGrantsForService", reflect.TypeOf((*MockConfigDatabase)(nil).ListAccessGrantsForService), ctx, serviceName)
}

// GetLatestAccessGrantForService mocks base method
func (m *MockConfigDatabase) GetLatestAccessGrantForService(ctx context.Context, organizationName, serviceName string) (*database.AccessGrant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestAccessGrantForService", ctx, organizationName, serviceName)
	ret0, _ := ret[0].(*database.AccessGrant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestAccessGrantForService indicates an expected call of GetLatestAccessGrantForService
func (mr *MockConfigDatabaseMockRecorder) GetLatestAccessGrantForService(ctx, organizationName, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestAccessGrantForService", reflect.TypeOf((*MockConfigDatabase)(nil).GetLatestAccessGrantForService), ctx, organizationName, serviceName)
}

// CreateAccessProof mocks base method
func (m *MockConfigDatabase) CreateAccessProof(ctx context.Context, accessProof *database.AccessProof) (*database.AccessProof, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessProof", ctx, accessProof)
	ret0, _ := ret[0].(*database.AccessProof)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccessProof indicates an expected call of CreateAccessProof
func (mr *MockConfigDatabaseMockRecorder) CreateAccessProof(ctx, accessProof interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessProof", reflect.TypeOf((*MockConfigDatabase)(nil).CreateAccessProof), ctx, accessProof)
}

// RevokeAccessProof mocks base method
func (m *MockConfigDatabase) RevokeAccessProof(ctx context.Context, organizationName, serviceName, id string, revokedAt time.Time) (*database.AccessProof, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeAccessProof", ctx, organizationName, serviceName, id, revokedAt)
	ret0, _ := ret[0].(*database.AccessProof)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeAccessProof indicates an expected call of RevokeAccessProof
func (mr *MockConfigDatabaseMockRecorder) RevokeAccessProof(ctx, organizationName, serviceName, id, revokedAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeAccessProof", reflect.TypeOf((*MockConfigDatabase)(nil).RevokeAccessProof), ctx, organizationName, serviceName, id, revokedAt)
}

// GetLatestAccessProofForService mocks base method
func (m *MockConfigDatabase) GetLatestAccessProofForService(ctx context.Context, organizationName, serviceName string) (*database.AccessProof, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestAccessProofForService", ctx, organizationName, serviceName)
	ret0, _ := ret[0].(*database.AccessProof)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestAccessProofForService indicates an expected call of GetLatestAccessProofForService
func (mr *MockConfigDatabaseMockRecorder) GetLatestAccessProofForService(ctx, organizationName, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestAccessProofForService", reflect.TypeOf((*MockConfigDatabase)(nil).GetLatestAccessProofForService), ctx, organizationName, serviceName)
}

// GetSettings mocks base method
func (m *MockConfigDatabase) GetSettings(ctx context.Context) (*database.Settings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSettings", ctx)
	ret0, _ := ret[0].(*database.Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSettings indicates an expected call of GetSettings
func (mr *MockConfigDatabaseMockRecorder) GetSettings(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSettings", reflect.TypeOf((*MockConfigDatabase)(nil).GetSettings), ctx)
}

// UpdateSettings mocks base method
func (m *MockConfigDatabase) UpdateSettings(ctx context.Context, settings *database.Settings) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSettings", ctx, settings)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSettings indicates an expected call of UpdateSettings
func (mr *MockConfigDatabaseMockRecorder) UpdateSettings(ctx, settings interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSettings", reflect.TypeOf((*MockConfigDatabase)(nil).UpdateSettings), ctx, settings)
}
