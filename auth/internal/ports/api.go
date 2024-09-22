package ports

import (
	"context"

	"github.com/Bookil/microservices/auth/internal/application/core/domain"
)

type APIPort interface {
	Register(ctx context.Context, firstName, lastName, email, password string) (string, error)
	VerifyEmail(ctx context.Context, email, verificationCode string) error
	SendVerificationCodeAgain(ctx context.Context, email string) error
	SendVerificationCode(ctx context.Context, email, name string) error
	Login(ctx context.Context, email, password string) (accessToken string, refreshToken string, err error)
	ResetPassword(ctx context.Context, email string) error
	SubmitResetPassword(ctx context.Context, SubmitResetPasswordToken string, newPassword string) error
	Authenticate(ctx context.Context, accessToken string) (domain.UserID, error)
	ChangePassword(ctx context.Context, userID domain.UserID, newPassword string, oldPassword string) error
	RefreshToken(ctx context.Context, userID domain.UserID, refreshToken string) (accessToken string, err error)
	DeleteAccount(ctx context.Context, userID, password string) error
}
