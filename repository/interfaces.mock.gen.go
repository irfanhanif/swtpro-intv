// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	entity "github.com/irfanhanif/swtpro-intv/entity"
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
func (m *MockRepositoryInterface) GetTestById(ctx context.Context, input GetTestByIdInput) (GetTestByIdOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestById", ctx, input)
	ret0, _ := ret[0].(GetTestByIdOutput)
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

// MockIGetUserByPhoneNumber is a mock of IGetUserByPhoneNumber interface.
type MockIGetUserByPhoneNumber struct {
	ctrl     *gomock.Controller
	recorder *MockIGetUserByPhoneNumberMockRecorder
}

// MockIGetUserByPhoneNumberMockRecorder is the mock recorder for MockIGetUserByPhoneNumber.
type MockIGetUserByPhoneNumberMockRecorder struct {
	mock *MockIGetUserByPhoneNumber
}

// NewMockIGetUserByPhoneNumber creates a new mock instance.
func NewMockIGetUserByPhoneNumber(ctrl *gomock.Controller) *MockIGetUserByPhoneNumber {
	mock := &MockIGetUserByPhoneNumber{ctrl: ctrl}
	mock.recorder = &MockIGetUserByPhoneNumberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIGetUserByPhoneNumber) EXPECT() *MockIGetUserByPhoneNumberMockRecorder {
	return m.recorder
}

// GetUserByPhoneNumber mocks base method.
func (m *MockIGetUserByPhoneNumber) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (entity.IUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByPhoneNumber", ctx, phoneNumber)
	ret0, _ := ret[0].(entity.IUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByPhoneNumber indicates an expected call of GetUserByPhoneNumber.
func (mr *MockIGetUserByPhoneNumberMockRecorder) GetUserByPhoneNumber(ctx, phoneNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByPhoneNumber", reflect.TypeOf((*MockIGetUserByPhoneNumber)(nil).GetUserByPhoneNumber), ctx, phoneNumber)
}

// MockIIncrementLoginCount is a mock of IIncrementLoginCount interface.
type MockIIncrementLoginCount struct {
	ctrl     *gomock.Controller
	recorder *MockIIncrementLoginCountMockRecorder
}

// MockIIncrementLoginCountMockRecorder is the mock recorder for MockIIncrementLoginCount.
type MockIIncrementLoginCountMockRecorder struct {
	mock *MockIIncrementLoginCount
}

// NewMockIIncrementLoginCount creates a new mock instance.
func NewMockIIncrementLoginCount(ctrl *gomock.Controller) *MockIIncrementLoginCount {
	mock := &MockIIncrementLoginCount{ctrl: ctrl}
	mock.recorder = &MockIIncrementLoginCountMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIIncrementLoginCount) EXPECT() *MockIIncrementLoginCountMockRecorder {
	return m.recorder
}

// IncrementLoginCount mocks base method.
func (m *MockIIncrementLoginCount) IncrementLoginCount(ctx context.Context, userID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrementLoginCount", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrementLoginCount indicates an expected call of IncrementLoginCount.
func (mr *MockIIncrementLoginCountMockRecorder) IncrementLoginCount(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementLoginCount", reflect.TypeOf((*MockIIncrementLoginCount)(nil).IncrementLoginCount), ctx, userID)
}

// MockIGetUserByID is a mock of IGetUserByID interface.
type MockIGetUserByID struct {
	ctrl     *gomock.Controller
	recorder *MockIGetUserByIDMockRecorder
}

// MockIGetUserByIDMockRecorder is the mock recorder for MockIGetUserByID.
type MockIGetUserByIDMockRecorder struct {
	mock *MockIGetUserByID
}

// NewMockIGetUserByID creates a new mock instance.
func NewMockIGetUserByID(ctrl *gomock.Controller) *MockIGetUserByID {
	mock := &MockIGetUserByID{ctrl: ctrl}
	mock.recorder = &MockIGetUserByIDMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIGetUserByID) EXPECT() *MockIGetUserByIDMockRecorder {
	return m.recorder
}

// GetUserByID mocks base method.
func (m *MockIGetUserByID) GetUserByID(ctx context.Context, id uuid.UUID) (entity.IUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, id)
	ret0, _ := ret[0].(entity.IUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockIGetUserByIDMockRecorder) GetUserByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockIGetUserByID)(nil).GetUserByID), ctx, id)
}
