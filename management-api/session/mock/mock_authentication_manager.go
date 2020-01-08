// Code generated by MockGen. DO NOT EDIT.
// Source: authenticationManager.go

// Package mock_session is a generated GoMock package.
package mock_session

import (
	http "net/http"
	reflect "reflect"

	chi "github.com/go-chi/chi"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthenticationManager is a mock of AuthenticationManager interface
type MockAuthenticationManager struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticationManagerMockRecorder
}

// MockAuthenticationManagerMockRecorder is the mock recorder for MockAuthenticationManager
type MockAuthenticationManagerMockRecorder struct {
	mock *MockAuthenticationManager
}

// NewMockAuthenticationManager creates a new mock instance
func NewMockAuthenticationManager(ctrl *gomock.Controller) *MockAuthenticationManager {
	mock := &MockAuthenticationManager{ctrl: ctrl}
	mock.recorder = &MockAuthenticationManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthenticationManager) EXPECT() *MockAuthenticationManagerMockRecorder {
	return m.recorder
}

// Middleware mocks base method
func (m *MockAuthenticationManager) Middleware(next http.Handler) http.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Middleware", next)
	ret0, _ := ret[0].(http.Handler)
	return ret0
}

// Middleware indicates an expected call of Middleware
func (mr *MockAuthenticationManagerMockRecorder) Middleware(next interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Middleware", reflect.TypeOf((*MockAuthenticationManager)(nil).Middleware), next)
}

// Routes mocks base method
func (m *MockAuthenticationManager) Routes() chi.Router {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Routes")
	ret0, _ := ret[0].(chi.Router)
	return ret0
}

// Routes indicates an expected call of Routes
func (mr *MockAuthenticationManagerMockRecorder) Routes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Routes", reflect.TypeOf((*MockAuthenticationManager)(nil).Routes))
}
