package ports

import (
	"context"
	"time"

	"github.com/Bookil/microservices/auth/internal/application/core/domain"
)

type DBPort interface {
	Create(ctx context.Context, auth *domain.Auth) (*domain.Auth, error)
	GetByID(ctx context.Context, userID domain.UserID) (*domain.Auth, error)
	ChangePassword(ctx context.Context, userID domain.UserID, hashedPassword string) (*domain.Auth, error)
	VerifyEmail(ctx context.Context, userID domain.UserID) (*domain.Auth, error)
	IncrementFailedLoginAttempts(ctx context.Context, userID domain.UserID) (*domain.Auth, error)
	ClearFailedLoginAttempts(ctx context.Context, userID domain.UserID) (*domain.Auth, error)
	LockAccount(ctx context.Context, userID domain.UserID, lockDuration time.Duration) (*domain.Auth, error)
	UnlockAccount(ctx context.Context, userID domain.UserID) (*domain.Auth, error)
	DeleteByID(ctx context.Context, userID domain.UserID) error
	Save(ctx context.Context, auth *domain.Auth) (*domain.Auth, error)
}
