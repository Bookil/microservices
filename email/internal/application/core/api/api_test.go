package api_test

import (
	"email/internal/application/core/api"
	"email/mocks"
	"fmt"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestSendResetPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSMTPPort := mocks.NewMockSMTPPort(ctrl)
	application := api.NewApplication(mockSMTPPort)

	email := "test@example.com"
	name := "John Doe"
	url := "https://example.com/reset-password"
	expiry := "120" //seconds

	mockSMTPPort.EXPECT().SendResetPassword(email, name, url, expiry).Return(nil)


	
	err := application.SendResetPassword(email,name,url,expiry)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

func TestSendVerificationCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSMTPPort := mocks.NewMockSMTPPort(ctrl)
	application := api.NewApplication(mockSMTPPort)

	
	email := "test@example.com"
	name := "John Doe"
	code := "123456"

	mockSMTPPort.EXPECT().SendVerificationCode(email, name, code).Return(nil)

	err := application.SendVerificationCode(email, name, code)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

func TestSendWelcome(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSMTPPort := mocks.NewMockSMTPPort(ctrl)
	application := api.NewApplication(mockSMTPPort)
	
	fullName := "John Doe"
	email := "test@example.com"

	mockSMTPPort.EXPECT().SendWelcome(fullName, email).Return(nil)

	err := application.SendWelcome(fullName, email)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

func TestSendResetPasswordError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSMTPPort := mocks.NewMockSMTPPort(ctrl)
	application := api.NewApplication(mockSMTPPort)

	email := "test@example.com"
	name := "John Doe"
	url := "https://example.com/reset-password"
	expiry := "24 hours"

	expectedError := fmt.Errorf("failed to send reset password email")
	mockSMTPPort.EXPECT().SendResetPassword(email, name, url, expiry).Return(expectedError)

	err := application.SendResetPassword(email, name, url, expiry)
	if err != expectedError {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}
}

func TestSendVerificationCodeError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSMTPPort := mocks.NewMockSMTPPort(ctrl)
	application := api.NewApplication(mockSMTPPort)

	email := "test@example.com"
	name := "John Doe"
	code := "123456"

	expectedError := fmt.Errorf("failed to send verification code email")
	mockSMTPPort.EXPECT().SendVerificationCode(email, name, code).Return(expectedError)

	err := application.SendVerificationCode(email, name, code)
	if err != expectedError {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}
}

func TestSendWelcomeError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSMTPPort := mocks.NewMockSMTPPort(ctrl)
	application := api.NewApplication(mockSMTPPort)

	fullName := "John Doe"
	email := "test@example.com"

	expectedError := fmt.Errorf("failed to send welcome email")
	mockSMTPPort.EXPECT().SendWelcome(fullName, email).Return(expectedError)

	err := application.SendWelcome(fullName, email)
	if err != expectedError {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}
}
