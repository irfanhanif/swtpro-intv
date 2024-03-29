// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	entity "github.com/irfanhanif/swtpro-intv/entity"
)

// MockINewUser is a mock of INewUser interface.
type MockINewUser struct {
	ctrl     *gomock.Controller
	recorder *MockINewUserMockRecorder
}

// MockINewUserMockRecorder is the mock recorder for MockINewUser.
type MockINewUserMockRecorder struct {
	mock *MockINewUser
}

// NewMockINewUser creates a new mock instance.
func NewMockINewUser(ctrl *gomock.Controller) *MockINewUser {
	mock := &MockINewUser{ctrl: ctrl}
	mock.recorder = &MockINewUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockINewUser) EXPECT() *MockINewUserMockRecorder {
	return m.recorder
}

// NewUser mocks base method.
func (m *MockINewUser) NewUser(phoneNumber, password, fullName string) entity.IUser {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUser", phoneNumber, password, fullName)
	ret0, _ := ret[0].(entity.IUser)
	return ret0
}

// NewUser indicates an expected call of NewUser.
func (mr *MockINewUserMockRecorder) NewUser(phoneNumber, password, fullName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUser", reflect.TypeOf((*MockINewUser)(nil).NewUser), phoneNumber, password, fullName)
}

// NewUserWithID mocks base method.
func (m *MockINewUser) NewUserWithID(id uuid.UUID, phoneNumber, password, fullName string) entity.IUser {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUserWithID", id, phoneNumber, password, fullName)
	ret0, _ := ret[0].(entity.IUser)
	return ret0
}

// NewUserWithID indicates an expected call of NewUserWithID.
func (mr *MockINewUserMockRecorder) NewUserWithID(id, phoneNumber, password, fullName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUserWithID", reflect.TypeOf((*MockINewUser)(nil).NewUserWithID), id, phoneNumber, password, fullName)
}

// MockIUser is a mock of IUser interface.
type MockIUser struct {
	ctrl     *gomock.Controller
	recorder *MockIUserMockRecorder
}

// MockIUserMockRecorder is the mock recorder for MockIUser.
type MockIUserMockRecorder struct {
	mock *MockIUser
}

// NewMockIUser creates a new mock instance.
func NewMockIUser(ctrl *gomock.Controller) *MockIUser {
	mock := &MockIUser{ctrl: ctrl}
	mock.recorder = &MockIUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUser) EXPECT() *MockIUserMockRecorder {
	return m.recorder
}

// FullName mocks base method.
func (m *MockIUser) FullName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FullName")
	ret0, _ := ret[0].(string)
	return ret0
}

// FullName indicates an expected call of FullName.
func (mr *MockIUserMockRecorder) FullName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FullName", reflect.TypeOf((*MockIUser)(nil).FullName))
}

// HashedPassword mocks base method.
func (m *MockIUser) HashedPassword() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashedPassword")
	ret0, _ := ret[0].(string)
	return ret0
}

// HashedPassword indicates an expected call of HashedPassword.
func (mr *MockIUserMockRecorder) HashedPassword() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashedPassword", reflect.TypeOf((*MockIUser)(nil).HashedPassword))
}

// ID mocks base method.
func (m *MockIUser) ID() uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(uuid.UUID)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockIUserMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockIUser)(nil).ID))
}

// Password mocks base method.
func (m *MockIUser) Password() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Password")
	ret0, _ := ret[0].(string)
	return ret0
}

// Password indicates an expected call of Password.
func (mr *MockIUserMockRecorder) Password() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Password", reflect.TypeOf((*MockIUser)(nil).Password))
}

// PhoneNumber mocks base method.
func (m *MockIUser) PhoneNumber() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PhoneNumber")
	ret0, _ := ret[0].(string)
	return ret0
}

// PhoneNumber indicates an expected call of PhoneNumber.
func (mr *MockIUserMockRecorder) PhoneNumber() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PhoneNumber", reflect.TypeOf((*MockIUser)(nil).PhoneNumber))
}

// Validate mocks base method.
func (m *MockIUser) Validate() []error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate")
	ret0, _ := ret[0].([]error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockIUserMockRecorder) Validate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockIUser)(nil).Validate))
}
