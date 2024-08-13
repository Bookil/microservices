package ports

import (
	"time"

	"github.com/Bookil/microservices/auth/internal/application/core/domain"
)

type DBPort interface {
	Create(auth *domain.Auth) (error)
	GetUserAuth(userID domain.UserID) (*domain.Auth, error)
	ChangePassword(userID domain.UserID, passwordHash string) error
	VerifyEmail(userID domain.UserID) error
	IncrementFailedLoginAttempts(userID domain.UserID) error
	ClearFailedLoginAttempts(userID domain.UserID) error
	LockAccount(userID domain.UserID, lockDuration time.Duration) error
	UnlockAccount(userID domain.UserID) error
	DeleteByID(userID domain.UserID) error
}
