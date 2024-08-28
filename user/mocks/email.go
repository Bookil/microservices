package mocks

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type MockedEmail struct {
	mock.Mock
}

func (m *MockedEmail) SendVerificationCode(email, code string) error {
	args := m.Called(email, code)
	return args.Error(0)
}

func (m *MockedEmail) SendResetPassword(url,token,email string, duration time.Duration) error {
	args := m.Called(url, token,duration, email)
	return args.Error(0)
}

func (m *MockedEmail) SendWellCome(email string) error {
	args := m.Called(email)
	return args.Error(0)
}
