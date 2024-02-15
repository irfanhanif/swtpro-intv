// Code generated by MockGen. DO NOT EDIT.
// Source: user_authentication.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/irfanhanif/swtpro-intv/entity"
)

// MockINewUserAuthentication is a mock of INewUserAuthentication interface.
type MockINewUserAuthentication struct {
	ctrl     *gomock.Controller
	recorder *MockINewUserAuthenticationMockRecorder
}

// MockINewUserAuthenticationMockRecorder is the mock recorder for MockINewUserAuthentication.
type MockINewUserAuthenticationMockRecorder struct {
	mock *MockINewUserAuthentication
}

// NewMockINewUserAuthentication creates a new mock instance.
func NewMockINewUserAuthentication(ctrl *gomock.Controller) *MockINewUserAuthentication {
	mock := &MockINewUserAuthentication{ctrl: ctrl}
	mock.recorder = &MockINewUserAuthenticationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockINewUserAuthentication) EXPECT() *MockINewUserAuthenticationMockRecorder {
	return m.recorder
}

// NewUserAuthentication mocks base method.
func (m *MockINewUserAuthentication) NewUserAuthentication(phoneNumber, password string) entity.IUserAuthentication {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUserAuthentication", phoneNumber, password)
	ret0, _ := ret[0].(entity.IUserAuthentication)
	return ret0
}

// NewUserAuthentication indicates an expected call of NewUserAuthentication.
func (mr *MockINewUserAuthenticationMockRecorder) NewUserAuthentication(phoneNumber, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUserAuthentication", reflect.TypeOf((*MockINewUserAuthentication)(nil).NewUserAuthentication), phoneNumber, password)
}

// MockIUserAuthentication is a mock of IUserAuthentication interface.
type MockIUserAuthentication struct {
	ctrl     *gomock.Controller
	recorder *MockIUserAuthenticationMockRecorder
}

// MockIUserAuthenticationMockRecorder is the mock recorder for MockIUserAuthentication.
type MockIUserAuthenticationMockRecorder struct {
	mock *MockIUserAuthentication
}

// NewMockIUserAuthentication creates a new mock instance.
func NewMockIUserAuthentication(ctrl *gomock.Controller) *MockIUserAuthentication {
	mock := &MockIUserAuthentication{ctrl: ctrl}
	mock.recorder = &MockIUserAuthenticationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserAuthentication) EXPECT() *MockIUserAuthenticationMockRecorder {
	return m.recorder
}

// Password mocks base method.
func (m *MockIUserAuthentication) Password() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Password")
	ret0, _ := ret[0].(string)
	return ret0
}

// Password indicates an expected call of Password.
func (mr *MockIUserAuthenticationMockRecorder) Password() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Password", reflect.TypeOf((*MockIUserAuthentication)(nil).Password))
}

// PhoneNumber mocks base method.
func (m *MockIUserAuthentication) PhoneNumber() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PhoneNumber")
	ret0, _ := ret[0].(string)
	return ret0
}

// PhoneNumber indicates an expected call of PhoneNumber.
func (mr *MockIUserAuthenticationMockRecorder) PhoneNumber() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PhoneNumber", reflect.TypeOf((*MockIUserAuthentication)(nil).PhoneNumber))
}

// Validate mocks base method.
func (m *MockIUserAuthentication) Validate() []error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].([]error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockIUserAuthenticationMockRecorder) Validate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockIUserAuthentication)(nil).Validate))
}
