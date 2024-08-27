package ports

import (
	"context"
)

type APIPort interface {
	Register(ctx context.Context, firstName, lastName, email, password string) (string, error)
	Login(ctx context.Context, email, password string) (string, string, error)
	ChangePassword(ctx context.Context, accessToken, oldPassword, newPassword string) error
	ResetPassword(ctx context.Context, email string) error
	Update(ctx context.Context, accessToken, firstName, lastName string) error
	DeleteAccount(ctx context.Context, accessToken, password string) error
}
