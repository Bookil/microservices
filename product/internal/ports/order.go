package ports

import (
	"context"

	"product/internal/application/core/domain"
)

type OrderPort interface {
	AddBookToOrder(ctx context.Context, ID domain.BookID, quantity uint) error
	DeleteBookOfOrderByID(ctx context.Context, ID domain.BookID) error
}
