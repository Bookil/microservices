package ports

import (
	"context"

	"github.com/Bookil/microservices/auth/internal/application/core/domain"
)

type UserPorts interface {
	Register(ctx context.Context, firstName, lastName, email string) (domain.UserID,error)
	GetUserIDByEmail(ctx context.Context,email string)(domain.UserID,error)
}
