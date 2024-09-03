package ports

import (
	"context"

	"github.com/Bookil/microservices/user/internal/application/core/domain"
)

type AuthPort interface {
	Authenticate(ctx context.Context, accessToken string) (domain.UserID, error)
	ChangePassword(ctx context.Context, userID domain.UserID, oldPassword string, newPassword string) error
	DeleteAccount(ctx context.Context, userID domain.UserID, password string) error
}
