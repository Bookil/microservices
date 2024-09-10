package ports

import "github.com/Bookil/microservices/auth/internal/application/core/domain"

type Validation interface {
	ValidateRegisterInputs(firstName, lastName, email, password string) error
	ValidateLoginInputs(email, password string) error
	ValidateSendVerificationCode(email string) error
	ValidateVerifyEmailInputs(string, string) error
	ValidateChangePasswordInputs(domain.UserID, string, string) error
	ValidateAuthenticateInputs(string) error
	ValidateRefreshTokenInputs(domain.UserID, string) error
	ValidateResetPasswordInputs(email string) error
	ValidateSubmitResetPasswordInputs(string, string) error
	ValidateDeleteAccountInputs(domain.UserID, string) error
}
