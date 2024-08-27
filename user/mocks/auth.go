package mocks

import (
	"context"
	"time"

	"github.com/Bookil/microservices/user/internal/application/core/domain"
	"github.com/stretchr/testify/mock"
)

type MockedAuth struct {
	mock.Mock
}

func (m *MockedAuth) Register(ctx context.Context, userID domain.UserID, password string) (string, error) {
	args := m.Called(userID, password)
	return args.String(0), args.Error(1)
}

func (m *MockedAuth) Authenticate(ctx context.Context, accessToken string) (domain.UserID, error) {
	args := m.Called(accessToken)
	return args.Get(0).(domain.UserID), args.Error(1)
}

func (m *MockedAuth) VerifyEmail(ctx context.Context, userID domain.UserID, code string) error {
	args := m.Called(userID, code)
	return args.Error(0)
}

func (m *MockedAuth) Update(ctx context.Context, userID domain.UserID, firstName, lastName string) (*domain.User, error) {
	args := m.Called(userID, firstName, lastName)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockedAuth) Login(ctx context.Context, userID, password string) (string, string, error) {
	args := m.Called(userID, userID, password)
	return args.String(0), args.String(1), args.Error(1)
}

func (m *MockedAuth) ChangePassword(ctx context.Context, userID domain.UserID, newPassword string, oldPassword string) error {
	args := m.Called(userID, oldPassword, newPassword)
	return args.Error(0)
}

func (m *MockedAuth) ResetPassword(ctx context.Context, userID string) (string, time.Duration, error) {
	args := m.Called(userID)
	return args.String(0),time.Duration(args.Int(1)),args.Error(2)
}

func(m *MockedAuth)SubmitResetPassword(ctx context.Context, token string, newPassword string) error{
	args := m.Called(token,newPassword)
	return args.Error(0)
}

func(m *MockedAuth)DeleteAccount(ctx context.Context, userID domain.UserID, password string) error{
	args := m.Called(userID,password)
	return args.Error(0)
}