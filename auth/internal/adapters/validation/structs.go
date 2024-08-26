package validation

import "github.com/Bookil/microservices/auth/internal/application/core/domain"

type registerInputs struct {
	UserID   domain.UserID `validate:"required"`
	Password string `validate:"required,min=8"`
}

type loginInputs struct {
	UserID   domain.UserID `validate:"required"`
	Password string        `validate:"required"`
}

type authenticateInputs struct {
	AccessToken string `validate:"required"`
}

type verifyEmailInputs struct {
	UserID         domain.UserID `validate:"required"`
	ValidationCode string        `validate:"required,len=6"`
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
	UserID domain.UserID `validate:"required"`
}

type submitResetPasswordInputs struct {
	ResetPasswordToken string `validate:"required"`
	Password           string `validate:"required,min=8"`
}

type deleteAccountInputs struct {
	UserID   domain.UserID `validate:"required"`
	Password string        `validate:"required"`
}
