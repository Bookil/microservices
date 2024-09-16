// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/ports/payment.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=./internal/ports/payment.go -destination=./mocks/payment_mock.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	domain "github.com/Bookil/microservices/order/internal/application/core/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockPaymentPort is a mock of PaymentPort interface.
type MockPaymentPort struct {
	ctrl     *gomock.Controller
	recorder *MockPaymentPortMockRecorder
}

// MockPaymentPortMockRecorder is the mock recorder for MockPaymentPort.
type MockPaymentPortMockRecorder struct {
	mock *MockPaymentPort
}

// NewMockPaymentPort creates a new mock instance.
func NewMockPaymentPort(ctrl *gomock.Controller) *MockPaymentPort {
	mock := &MockPaymentPort{ctrl: ctrl}
	mock.recorder = &MockPaymentPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPaymentPort) EXPECT() *MockPaymentPortMockRecorder {
	return m.recorder
}

// Charge mocks base method.
func (m *MockPaymentPort) Charge(arg0 *domain.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Charge", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Charge indicates an expected call of Charge.
func (mr *MockPaymentPortMockRecorder) Charge(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Charge", reflect.TypeOf((*MockPaymentPort)(nil).Charge), arg0)
}
