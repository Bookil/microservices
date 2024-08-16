package mocks

import (
	"context"
	"time"

	"github.com/Bookil/microservices/auth/internal/application/core/domain"
	"github.com/stretchr/testify/mock"
)

type MockedDB struct {
	mock.Mock
}

func (d *MockedDB) Create(ctx context.Context, auth *domain.Auth) error {
	args := d.Called(auth)
	return args.Error(1)
}

func (d *MockedDB) GetByID(ctx context.Context, userID domain.UserID) (*domain.Auth, error) {
	args := d.Called(userID)
	return args.Get(0).(*domain.Auth), args.Error(1)
}

func (d *MockedDB) ChangePassword(ctx context.Context, userID domain.UserID, hashedPassword string) (*domain.Auth, error) {
	args := d.Called(userID, hashedPassword)
	return args.Get(0).(*domain.Auth), args.Error(1)
}

func (d *MockedDB) VerifyEmail(ctx context.Context, userID domain.UserID) (*domain.Auth, error) {
	args := d.Called(userID)
	return args.Get(0).(*domain.Auth), args.Error(1)
}

func (d *MockedDB) IncrementFailedLoginAttempts(ctx context.Context, userID domain.UserID) (*domain.Auth, error) {
	args := d.Called(userID)
	return args.Get(0).(*domain.Auth), args.Error(1)
}

func (d *MockedDB) ClearFailedLoginAttempts(ctx context.Context, userID domain.UserID) (*domain.Auth, error) {
	args := d.Called(userID)
	return args.Get(0).(*domain.Auth), args.Error(1)
}

func (d *MockedDB) LockAccount(ctx context.Context, userID domain.UserID, lockDuration time.Duration) (*domain.Auth, error) {
	args := d.Called(userID,lockDuration)
	return args.Get(0).(*domain.Auth), args.Error(1)
}

func (d *MockedDB) UnlockAccount(ctx context.Context, userID domain.UserID) (*domain.Auth, error){
	args := d.Called(userID)
	return args.Get(0).(*domain.Auth), args.Error(1)
}

func (d *MockedDB) DeleteByID(ctx context.Context, userID domain.UserID) error {
	args := d.Called(userID)
	return args.Error(1)
}

func (d *MockedDB) Save(ctx context.Context, auth *domain.Auth) (*domain.Auth, error){
	args := d.Called(auth)
	return args.Get(0).(*domain.Auth), args.Error(1)
}
