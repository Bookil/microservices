package mocks

import (
	"context"

	"github.com/Bookil/microservices/user/internal/application/core/domain"
	"github.com/stretchr/testify/mock"
)

type MockedDB struct {
	mock.Mock
}

func (m *MockedDB) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockedDB) Update(ctx context.Context, userID domain.UserID, firstName, lastName string) (*domain.User, error) {
	args := m.Called(userID, firstName, lastName)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockedDB) GetUserByID(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	args := m.Called(userID)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockedDB) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	args := m.Called(email)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockedDB) Delete(ctx context.Context, userID domain.UserID) error {
	args := m.Called(userID)
	return args.Error(0)
}
