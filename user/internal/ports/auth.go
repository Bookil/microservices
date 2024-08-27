package ports

import (
	"context"
	"time"

	"github.com/Bookil/microservices/user/internal/application/core/domain"
)

type AuthPort interface {
	Register(ctx context.Context, userID domain.UserID, password string) (string, error)
	Authenticate(ctx context.Context, accessToken string) (domain.UserID, error)
	VerifyEmail(ctx context.Context, userID domain.UserID, code string) error
	Login(ctx context.Context, userID, password string) (string, string, error)
	ChangePassword(ctx context.Context, userID domain.UserID, newPassword string, oldPassword string) error
	ResetPassword(ctx context.Context, userID string) (string, time.Duration, error)
	SubmitResetPassword(ctx context.Context, token string, newPassword string) error
	DeleteAccount(ctx context.Context, userID domain.UserID, password string) error
}