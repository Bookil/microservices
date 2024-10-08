// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/ports/cart.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=./internal/ports/cart.go -destination=./mocks/cart_mock.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	domain "product/internal/application/core/domain"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCartPort is a mock of CartPort interface.
type MockCartPort struct {
	ctrl     *gomock.Controller
	recorder *MockCartPortMockRecorder
}

// MockCartPortMockRecorder is the mock recorder for MockCartPort.
type MockCartPortMockRecorder struct {
	mock *MockCartPort
}

// NewMockCartPort creates a new mock instance.
func NewMockCartPort(ctrl *gomock.Controller) *MockCartPort {
	mock := &MockCartPort{ctrl: ctrl}
	mock.recorder = &MockCartPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCartPort) EXPECT() *MockCartPortMockRecorder {
	return m.recorder
}

// AddBookToCart mocks base method.
func (m *MockCartPort) AddBookToCart(ctx context.Context, bookID domain.BookID, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBookToCart", ctx, bookID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddBookToCart indicates an expected call of AddBookToCart.
func (mr *MockCartPortMockRecorder) AddBookToCart(ctx, bookID, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBookToCart", reflect.TypeOf((*MockCartPort)(nil).AddBookToCart), ctx, bookID, userID)
}

// DeleteBookFromCartByID mocks base method.
func (m *MockCartPort) DeleteBookFromCartByID(ctx context.Context, ID domain.BookID, cartID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBookFromCartByID", ctx, ID, cartID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBookFromCartByID indicates an expected call of DeleteBookFromCartByID.
func (mr *MockCartPortMockRecorder) DeleteBookFromCartByID(ctx, ID, cartID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBookFromCartByID", reflect.TypeOf((*MockCartPort)(nil).DeleteBookFromCartByID), ctx, ID, cartID)
}
