package ports

import (
	"context"

	"product/internal/application/core/domain"
)

type CartPort interface {
	AddBookToCart(ctx context.Context, bookID domain.BookID, userID string) error
	DeleteBookFromCartByID(ctx context.Context, ID domain.BookID, cartID uint) error
}
