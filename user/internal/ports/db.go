package ports

import (
	"context"

	"github.com/Bookil/microservices/user/internal/application/core/domain"
)

type DBPort interface {
	Create(ctx context.Context, firstName, lastName, email string) (*domain.User, error)
	Update(ctx context.Context, firstName, LastName string) (*domain.User, error)
	Delete(ctx context.Context, userID domain.UserID) error
}
