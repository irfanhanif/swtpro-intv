// Code generated by MockGen. DO NOT EDIT.
// Source: context.go

// Package mock is a generated GoMock package.
package mock

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIContext is a mock of IContext interface.
type MockIContext struct {
	ctrl     *gomock.Controller
	recorder *MockIContextMockRecorder
}

// MockIContextMockRecorder is the mock recorder for MockIContext.
type MockIContextMockRecorder struct {
	mock *MockIContext
}

// NewMockIContext creates a new mock instance.
func NewMockIContext(ctrl *gomock.Controller) *MockIContext {
	mock := &MockIContext{ctrl: ctrl}
	mock.recorder = &MockIContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIContext) EXPECT() *MockIContextMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockIContext) Get(key string) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockIContextMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIContext)(nil).Get), key)
}

// JSON mocks base method.
func (m *MockIContext) JSON(code int, i interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JSON", code, i)
	ret0, _ := ret[0].(error)
	return ret0
}

// JSON indicates an expected call of JSON.
func (mr *MockIContextMockRecorder) JSON(code, i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSON", reflect.TypeOf((*MockIContext)(nil).JSON), code, i)
}

// NoContent mocks base method.
func (m *MockIContext) NoContent(code int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NoContent", code)
	ret0, _ := ret[0].(error)
	return ret0
}

// NoContent indicates an expected call of NoContent.
func (mr *MockIContextMockRecorder) NoContent(code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NoContent", reflect.TypeOf((*MockIContext)(nil).NoContent), code)
}

// Param mocks base method.
func (m *MockIContext) Param(name string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Param", name)
	ret0, _ := ret[0].(string)
	return ret0
}

// Param indicates an expected call of Param.
func (mr *MockIContextMockRecorder) Param(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Param", reflect.TypeOf((*MockIContext)(nil).Param), name)
}

// Request mocks base method.
func (m *MockIContext) Request() *http.Request {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Request")
	ret0, _ := ret[0].(*http.Request)
	return ret0
}

// Request indicates an expected call of Request.
func (mr *MockIContextMockRecorder) Request() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockIContext)(nil).Request))
}

// Set mocks base method.
func (m *MockIContext) Set(key string, val interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", key, val)
}

// Set indicates an expected call of Set.
func (mr *MockIContextMockRecorder) Set(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockIContext)(nil).Set), key, val)
}
