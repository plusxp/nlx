// Code generated by MockGen. DO NOT EDIT.
// Source: irma/verification.go

// Package mock_insight_api is a generated GoMock package.
package mock_insight_api

import (
	rsa "crypto/rsa"
	reflect "reflect"

	jwt_go "github.com/dgrijalva/jwt-go"
	gomock "github.com/golang/mock/gomock"
	irma "go.nlx.io/nlx/insight-api/irma"
)

// MockJWTHandler is a mock of JWTHandler interface
type MockJWTHandler struct {
	ctrl     *gomock.Controller
	recorder *MockJWTHandlerMockRecorder
}

// MockJWTHandlerMockRecorder is the mock recorder for MockJWTHandler
type MockJWTHandlerMockRecorder struct {
	mock *MockJWTHandler
}

// NewMockJWTHandler creates a new mock instance
func NewMockJWTHandler(ctrl *gomock.Controller) *MockJWTHandler {
	mock := &MockJWTHandler{ctrl: ctrl}
	mock.recorder = &MockJWTHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockJWTHandler) EXPECT() *MockJWTHandlerMockRecorder {
	return m.recorder
}

// GenerateAndSignJWT mocks base method
func (m *MockJWTHandler) GenerateAndSignJWT(request *irma.DiscloseRequest, serviceProviderName string, rsaSignPrivateKey *rsa.PrivateKey) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAndSignJWT", request, serviceProviderName, rsaSignPrivateKey)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAndSignJWT indicates an expected call of GenerateAndSignJWT
func (mr *MockJWTHandlerMockRecorder) GenerateAndSignJWT(request, serviceProviderName, rsaSignPrivateKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAndSignJWT", reflect.TypeOf((*MockJWTHandler)(nil).GenerateAndSignJWT), request, serviceProviderName, rsaSignPrivateKey)
}

// VerifyIRMAVerificationResult mocks base method
func (m *MockJWTHandler) VerifyIRMAVerificationResult(jwtBytes []byte, rsaVerifyPublicKey *rsa.PublicKey) (*jwt_go.Token, *irma.VerificationResultClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyIRMAVerificationResult", jwtBytes, rsaVerifyPublicKey)
	ret0, _ := ret[0].(*jwt_go.Token)
	ret1, _ := ret[1].(*irma.VerificationResultClaims)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// VerifyIRMAVerificationResult indicates an expected call of VerifyIRMAVerificationResult
func (mr *MockJWTHandlerMockRecorder) VerifyIRMAVerificationResult(jwtBytes, rsaVerifyPublicKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyIRMAVerificationResult", reflect.TypeOf((*MockJWTHandler)(nil).VerifyIRMAVerificationResult), jwtBytes, rsaVerifyPublicKey)
}