package ports

import "context"

type API interface {
	Register(ctx context.Context, firstName, lastName, email, password string) error
	Login(ctx context.Context, email, password string) (string, string, error)
	VerifyEmail(ctx context.Context,email string, code string) error
	ChangePassword(ctx context.Context,oldPassword, newPassword string) error
	ResetPassword(ctx context.Context,email string) error
	DeleteAccount(ctx context.Context,password string) error
}
