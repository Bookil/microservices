package ports

import (
	"context"

	"github.com/Bookil/microservices/user/internal/application/core/domain"
)

type DBPort interface {
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	Update(ctx context.Context, userID domain.UserID, firstName, lastName string) (*domain.User, error)
	GetUserByID(ctx context.Context, id domain.UserID) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	DeleteByID(ctx context.Context, userID domain.UserID) error
}
