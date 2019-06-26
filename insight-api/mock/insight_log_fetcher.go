// Code generated by MockGen. DO NOT EDIT.
// Source: database.go

// Package mock_insight_api is a generated GoMock package.
package mock_insight_api

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	insight_api "go.nlx.io/nlx/insight-api"
	irma "go.nlx.io/nlx/insight-api/irma"
)

// MockInsightLogFetcher is a mock of InsightLogFetcher interface
type MockInsightLogFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockInsightLogFetcherMockRecorder
}

// MockInsightLogFetcherMockRecorder is the mock recorder for MockInsightLogFetcher
type MockInsightLogFetcherMockRecorder struct {
	mock *MockInsightLogFetcher
}

// NewMockInsightLogFetcher creates a new mock instance
func NewMockInsightLogFetcher(ctrl *gomock.Controller) *MockInsightLogFetcher {
	mock := &MockInsightLogFetcher{ctrl: ctrl}
	mock.recorder = &MockInsightLogFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInsightLogFetcher) EXPECT() *MockInsightLogFetcherMockRecorder {
	return m.recorder
}

// GetLogRecords mocks base method
func (m *MockInsightLogFetcher) GetLogRecords(rowsPerPage, page int, dataSubjectsByIrmaAttribute map[string][]string, claims *irma.VerificationResultClaims) (*insight_api.GetLogRecordsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogRecords", rowsPerPage, page, dataSubjectsByIrmaAttribute, claims)
	ret0, _ := ret[0].(*insight_api.GetLogRecordsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogRecords indicates an expected call of GetLogRecords
func (mr *MockInsightLogFetcherMockRecorder) GetLogRecords(rowsPerPage, page, dataSubjectsByIrmaAttribute, claims interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogRecords", reflect.TypeOf((*MockInsightLogFetcher)(nil).GetLogRecords), rowsPerPage, page, dataSubjectsByIrmaAttribute, claims)
}