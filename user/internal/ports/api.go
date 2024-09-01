package ports

import (
	"context"

	"github.com/Bookil/microservices/user/internal/application/core/domain"
)

type APIPort interface {
	Register(ctx context.Context, firstName, lastName, email string) (string, error)
	GetUserIDByEmail(ctx context.Context, email string) (domain.UserID, error)
	ChangePassword(ctx context.Context, userID, oldPassword, newPassword string) error
	Update(ctx context.Context, userID, firstName, lastName string) error
	DeleteAccount(ctx context.Context, userID, password string) error
}
