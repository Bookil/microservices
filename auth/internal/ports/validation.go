package ports

import "github.com/Bookil/microservices/auth/internal/application/core/domain"

type Validation interface {
	ValidateRegisterInputs(domain.UserID,string) error
	ValidateLoginInputs(domain.UserID, string) error
	ValidateVerifyEmailInputs(domain.UserID,string) error
	ValidateChangePasswordInputs(domain.UserID,string, string) error
	ValidateAuthenticateInputs(string) error
	ValidateRefreshTokenInputs(domain.UserID,string) error
	ValidateResetPasswordInputs(domain.UserID) error
	ValidateSubmitResetPasswordInputs(string, string) error
	ValidateDeleteAccountInputs(domain.UserID, string) error
}
