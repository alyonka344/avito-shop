// Code generated by MockGen. DO NOT EDIT.
// Source: C:/Users/User/Desktop/lab/avito-shop/internal/auth/authService.go

// Package mocks is a generated GoMock package.
package mocks

import (
	model "avito-shop/internal/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthService is a mock of AuthService interface.
type MockAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceMockRecorder
}

// MockAuthServiceMockRecorder is the mock recorder for MockAuthService.
type MockAuthServiceMockRecorder struct {
	mock *MockAuthService
}

// NewMockAuthService creates a new mock instance.
func NewMockAuthService(ctrl *gomock.Controller) *MockAuthService {
	mock := &MockAuthService{ctrl: ctrl}
	mock.recorder = &MockAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthService) EXPECT() *MockAuthServiceMockRecorder {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockAuthService) GenerateToken(user model.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthServiceMockRecorder) GenerateToken(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthService)(nil).GenerateToken), user)
}

// ValidateToken mocks base method.
func (m *MockAuthService) ValidateToken(token string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateToken", token)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateToken indicates an expected call of ValidateToken.
func (mr *MockAuthServiceMockRecorder) ValidateToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToken", reflect.TypeOf((*MockAuthService)(nil).ValidateToken), token)
}
