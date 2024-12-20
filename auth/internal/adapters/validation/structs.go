package validation

import "github.com/Bookil/microservices/auth/internal/application/core/domain"

type registerInputs struct {
	FirstName domain.UserID `validate:"required"`
	LastName  string        `validate:"required"`
	Email     string        `validate:"required,email"`
	Password  string        `validate:"required,min=8"`
}

type sendVerificationCodeAgainInputs struct {
	Email string `validate:"required,email"`
}

type loginInputs struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type authenticateAndAuthorizationInputs struct {
	AccessToken string `validate:"required,jwt"`
}

type verifyEmailInputs struct {
	Email          string `validate:"required,email"`
	ValidationCode string `validate:"required,len=6"`
}

type changePasswordInputs struct {
	UserID      domain.UserID `validate:"required"`
	OldPassword string        `validate:"required"`
	NewPassword string        `validate:"required,min=8"`
}

type refreshTokenInputs struct {
	UserID       domain.UserID `validate:"required"`
	RefreshToken string        `validate:"required"`
}

type resetPasswordInputs struct {
	Email string `validate:"required,email"`
}

type submitResetPasswordInputs struct {
	ResetPasswordToken string `validate:"required"`
	Password           string `validate:"required,min=8"`
}

type deleteAccountInputs struct {
	UserID   domain.UserID `validate:"required"`
	Password string        `validate:"required"`
}
