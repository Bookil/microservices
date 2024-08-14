package ports

import (
	"context"
	"time"

	"github.com/Bookil/microservices/auth/internal/application/core/domain"
)

type DBPort interface {
	Create(ctx context.Context, auth *domain.Auth) error
	GetByID(ctx context.Context, userID domain.UserID) (*domain.Auth, error)
	ChangePassword(ctx context.Context, userID domain.UserID, hashedPassword string) error
	VerifyEmail(ctx context.Context, userID domain.UserID) error
	IncrementFailedLoginAttempts(ctx context.Context, userID domain.UserID) error
	ClearFailedLoginAttempts(ctx context.Context, userID domain.UserID) error
	LockAccount(ctx context.Context, userID domain.UserID, lockDuration time.Duration) error
	UnlockAccount(ctx context.Context, userID domain.UserID) error
	DeleteByID(ctx context.Context, userID domain.UserID) error
	Save(ctx context.Context, auth *domain.Auth) error
}
