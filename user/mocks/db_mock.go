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
	context "context"
	reflect "reflect"

	domain "github.com/Bookil/microservices/user/internal/application/core/domain"
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

// Create mocks base method.
func (m *MockDBPort) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, user)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockDBPortMockRecorder) Create(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDBPort)(nil).Create), ctx, user)
}

// DeleteByID mocks base method.
func (m *MockDBPort) DeleteByID(ctx context.Context, userID domain.UserID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockDBPortMockRecorder) DeleteByID(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockDBPort)(nil).DeleteByID), ctx, userID)
}

// GetUserByEmail mocks base method.
func (m *MockDBPort) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockDBPortMockRecorder) GetUserByEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockDBPort)(nil).GetUserByEmail), ctx, email)
}

// GetUserByID mocks base method.
func (m *MockDBPort) GetUserByID(ctx context.Context, id domain.UserID) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, id)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockDBPortMockRecorder) GetUserByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockDBPort)(nil).GetUserByID), ctx, id)
}

// Update mocks base method.
func (m *MockDBPort) Update(ctx context.Context, userID domain.UserID, firstName, lastName string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, userID, firstName, lastName)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockDBPortMockRecorder) Update(ctx, userID, firstName, lastName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDBPort)(nil).Update), ctx, userID, firstName, lastName)
}
