// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/irfanhanif/swtpro-intv/entity"
	repository "github.com/irfanhanif/swtpro-intv/repository"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// GetTestById mocks base method.
func (m *MockRepositoryInterface) GetTestById(ctx context.Context, input repository.GetTestByIdInput) (repository.GetTestByIdOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestById", ctx, input)
	ret0, _ := ret[0].(repository.GetTestByIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestById indicates an expected call of GetTestById.
func (mr *MockRepositoryInterfaceMockRecorder) GetTestById(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestById", reflect.TypeOf((*MockRepositoryInterface)(nil).GetTestById), ctx, input)
}

// MockICreateNewUser is a mock of ICreateNewUser interface.
type MockICreateNewUser struct {
	ctrl     *gomock.Controller
	recorder *MockICreateNewUserMockRecorder
}

// MockICreateNewUserMockRecorder is the mock recorder for MockICreateNewUser.
type MockICreateNewUserMockRecorder struct {
	mock *MockICreateNewUser
}

// NewMockICreateNewUser creates a new mock instance.
func NewMockICreateNewUser(ctrl *gomock.Controller) *MockICreateNewUser {
	mock := &MockICreateNewUser{ctrl: ctrl}
	mock.recorder = &MockICreateNewUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICreateNewUser) EXPECT() *MockICreateNewUserMockRecorder {
	return m.recorder
}

// CreateNewUser mocks base method.
func (m *MockICreateNewUser) CreateNewUser(ctx context.Context, user entity.IUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNewUser indicates an expected call of CreateNewUser.
func (mr *MockICreateNewUserMockRecorder) CreateNewUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewUser", reflect.TypeOf((*MockICreateNewUser)(nil).CreateNewUser), ctx, user)
}