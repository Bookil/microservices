// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/ports/db.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=./internal/ports/db.go -destination=./mocks/db_mock.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	domain "github.com/Bookil/microservices/order/internal/application/core/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockDBPort is a mock of DBPort interface.
type MockDBPort struct {
	ctrl     *gomock.Controller
	recorder *MockDBPortMockRecorder
}

// MockDBPortMockRecorder is the mock recorder for MockDBPort.
type MockDBPortMockRecorder struct {
	mock *MockDBPort
}

// NewMockDBPort creates a new mock instance.
func NewMockDBPort(ctrl *gomock.Controller) *MockDBPort {
	mock := &MockDBPort{ctrl: ctrl}
	mock.recorder = &MockDBPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBPort) EXPECT() *MockDBPortMockRecorder {
	return m.recorder
}

// DeleteOrder mocks base method.
func (m *MockDBPort) DeleteOrder(orderItem domain.OrderID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrder", orderItem)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrder indicates an expected call of DeleteOrder.
func (mr *MockDBPortMockRecorder) DeleteOrder(orderItem any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrder", reflect.TypeOf((*MockDBPort)(nil).DeleteOrder), orderItem)
}

// GetOrder mocks base method.
func (m *MockDBPort) GetOrder(id domain.OrderID) (*domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", id)
	ret0, _ := ret[0].(*domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockDBPortMockRecorder) GetOrder(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockDBPort)(nil).GetOrder), id)
}

// SaveOrder mocks base method.
func (m *MockDBPort) SaveOrder(arg0 *domain.Order) (*domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveOrder", arg0)
	ret0, _ := ret[0].(*domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveOrder indicates an expected call of SaveOrder.
func (mr *MockDBPortMockRecorder) SaveOrder(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveOrder", reflect.TypeOf((*MockDBPort)(nil).SaveOrder), arg0)
}

// UpdateOrder mocks base method.
func (m *MockDBPort) UpdateOrder(orderID domain.OrderID, OrderItems *[]domain.OrderItem) (*domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrder", orderID, OrderItems)
	ret0, _ := ret[0].(*domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrder indicates an expected call of UpdateOrder.
func (mr *MockDBPortMockRecorder) UpdateOrder(orderID, OrderItems any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrder", reflect.TypeOf((*MockDBPort)(nil).UpdateOrder), orderID, OrderItems)
}
