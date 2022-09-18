// Code generated by MockGen. DO NOT EDIT.
// Source: domain/helper.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	http "net/http"
	reflect "reflect"
	domain "xm/domain"

	gomock "github.com/golang/mock/gomock"
)

// MockHelperHttpClientInterface is a mock of HelperHttpClientInterface interface.
type MockHelperHttpClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockHelperHttpClientInterfaceMockRecorder
}

// MockHelperHttpClientInterfaceMockRecorder is the mock recorder for MockHelperHttpClientInterface.
type MockHelperHttpClientInterfaceMockRecorder struct {
	mock *MockHelperHttpClientInterface
}

// NewMockHelperHttpClientInterface creates a new mock instance.
func NewMockHelperHttpClientInterface(ctrl *gomock.Controller) *MockHelperHttpClientInterface {
	mock := &MockHelperHttpClientInterface{ctrl: ctrl}
	mock.recorder = &MockHelperHttpClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHelperHttpClientInterface) EXPECT() *MockHelperHttpClientInterfaceMockRecorder {
	return m.recorder
}

// IpapiRequest mocks base method.
func (m *MockHelperHttpClientInterface) IpapiRequest(client *http.Client) domain.Helper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IpapiRequest", client)
	ret0, _ := ret[0].(domain.Helper)
	return ret0
}

// IpapiRequest indicates an expected call of IpapiRequest.
func (mr *MockHelperHttpClientInterfaceMockRecorder) IpapiRequest(client interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpapiRequest", reflect.TypeOf((*MockHelperHttpClientInterface)(nil).IpapiRequest), client)
}
