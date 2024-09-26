package ports

import (
	"context"
	"product/internal/application/core/domain"
)

type APIPort interface {
	GetAllBooks(ctx context.Context) ([]*domain.Book, error)
	AddBookToOrder(ctx context.Context,ID domain.BookID, quantity uint) error
	DeleteBookOfOrder(ctx context.Context,ID domain.BookID) error
	GetBook(ctx context.Context,ID domain.BookID) (*domain.Book, error)
	AddBook(ctx context.Context,name, writeName, description string, genre domain.Genre, tags []string, price float64) (*domain.Book, error)
	ModifyBook(ctx context.Context,ID domain.BookID, name, writeName, description string, genre domain.Genre, tags []string, price float64) (*domain.Book, error)
	DeleteBook(ctx context.Context,ID domain.BookID) error
}
