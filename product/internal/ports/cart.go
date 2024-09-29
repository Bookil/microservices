package ports

import (
	"context"

	"product/internal/application/core/domain"
)

type CartPort interface {
	AddBookToCart(ctx context.Context, ID domain.BookID, quantity uint) error
	DeleteBookFromCartByID(ctx context.Context, ID domain.BookID) error
}
