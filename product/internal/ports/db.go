package ports

import (
	"context"
	"product/internal/application/core/domain"
)

type DBPorts interface {
	GetAllBooks(ctx context.Context) ([]*domain.Book, error)
	GetBookByID(ctx context.Context,ID domain.BookID) (*domain.Book, error)
	GetBooksByTitle(ctx context.Context,title string) ([]*domain.Book, error)
	GetBooksByAuthor(ctx context.Context,author string) ([]*domain.Book, error)
	GetBooksByGenre(ctx context.Context,genreName string) ([]*domain.Book, error)
	CreateBook(ctx context.Context,book *domain.Book) (*domain.Book, error)
	ModifyBookByID(ctx context.Context,book *domain.Book) (*domain.Book, error)
	DeleteBookByID(ctx context.Context,ID domain.BookID) error
}
