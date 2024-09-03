package validation

import "github.com/Bookil/microservices/user/internal/application/core/domain"

type registerInputs struct {
	FirstName string `validate:"required,min=1,max=50"`
	LastName  string `validate:"required,min=1,max=50"`
	Email     string `validate:"required,email"`
}

type getUserIDByEmailInputs struct {
	Email string `validate:"required,email"`
}

type changePasswordInputs struct {
	UserID      domain.UserID `validate:"required"`
	OldPassword string        `validate:"required"`
	NewPassword string        `validate:"required,min=8"`
}

type updateInputs struct {
	UserID    domain.UserID `validate:"required"`
	FirstName string        `validate:"required,min=1,max=50"`
	LastName  string        `validate:"required,min=1,max=50"`
}

type deleteAccountInputs struct {
	UserID   domain.UserID `validate:"required"`
	Password string        `validate:"required"`
}
