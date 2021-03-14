// Code generated by MockGen. DO NOT EDIT.
// Source: internal/entrypoint/handler/interfaces.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	echo "github.com/labstack/echo/v4"
)

// MockPersonHandler is a mock of PersonHandler interface.
type MockPersonHandler struct {
	ctrl     *gomock.Controller
	recorder *MockPersonHandlerMockRecorder
}

// MockPersonHandlerMockRecorder is the mock recorder for MockPersonHandler.
type MockPersonHandlerMockRecorder struct {
	mock *MockPersonHandler
}

// NewMockPersonHandler creates a new mock instance.
func NewMockPersonHandler(ctrl *gomock.Controller) *MockPersonHandler {
	mock := &MockPersonHandler{ctrl: ctrl}
	mock.recorder = &MockPersonHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPersonHandler) EXPECT() *MockPersonHandlerMockRecorder {
	return m.recorder
}

// GetById mocks base method.
func (m *MockPersonHandler) GetById(arg0 echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetById indicates an expected call of GetById.
func (mr *MockPersonHandlerMockRecorder) GetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockPersonHandler)(nil).GetById), arg0)
}
