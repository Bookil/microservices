// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/ports/email.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=./internal/ports/email.go -destination=./mocks/email_mock.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockSMTPPort is a mock of SMTPPort interface.
type MockSMTPPort struct {
	ctrl     *gomock.Controller
	recorder *MockSMTPPortMockRecorder
}

// MockSMTPPortMockRecorder is the mock recorder for MockSMTPPort.
type MockSMTPPortMockRecorder struct {
	mock *MockSMTPPort
}

// NewMockSMTPPort creates a new mock instance.
func NewMockSMTPPort(ctrl *gomock.Controller) *MockSMTPPort {
	mock := &MockSMTPPort{ctrl: ctrl}
	mock.recorder = &MockSMTPPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSMTPPort) EXPECT() *MockSMTPPortMockRecorder {
	return m.recorder
}

// SendResetPassword mocks base method.
func (m *MockSMTPPort) SendResetPassword(email, name, url, expiry string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendResetPassword", email, name, url, expiry)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendResetPassword indicates an expected call of SendResetPassword.
func (mr *MockSMTPPortMockRecorder) SendResetPassword(email, name, url, expiry any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendResetPassword", reflect.TypeOf((*MockSMTPPort)(nil).SendResetPassword), email, name, url, expiry)
}

// SendVerificationCode mocks base method.
func (m *MockSMTPPort) SendVerificationCode(email, name, code string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendVerificationCode", email, name, code)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendVerificationCode indicates an expected call of SendVerificationCode.
func (mr *MockSMTPPortMockRecorder) SendVerificationCode(email, name, code any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendVerificationCode", reflect.TypeOf((*MockSMTPPort)(nil).SendVerificationCode), email, name, code)
}

// SendWelcome mocks base method.
func (m *MockSMTPPort) SendWelcome(fullName, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendWelcome", fullName, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendWelcome indicates an expected call of SendWelcome.
func (mr *MockSMTPPortMockRecorder) SendWelcome(fullName, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendWelcome", reflect.TypeOf((*MockSMTPPort)(nil).SendWelcome), fullName, email)
}
