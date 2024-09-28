package ports

import (
	"context"
	"product/internal/application/core/domain"
)

type DBPorts interface {
	AddAuthor(ctx context.Context, author *domain.Author) (*domain.Author, error)
	GetAllAuthors(ctx context.Context) ([]domain.Author, error)
	AddBook(ctx context.Context, book *domain.Book) (*domain.Book, error)
	GetAllBooks(ctx context.Context) ([]domain.Book, error)
	GetBookByID(ctx context.Context, ID domain.BookID) (*domain.Book, error)
	GetBooksByTitle(ctx context.Context, title string) ([]domain.Book, error)
	GetBooksByAuthor(ctx context.Context, authorName string) ([]domain.Book, error)
	GetBooksByGenre(ctx context.Context, genreName string) ([]domain.Book, error)
	ModifyBookByID(ctx context.Context, ID domain.BookID, book *domain.Book) (*domain.Book, error)
	DeleteBookByID(ctx context.Context, ID domain.BookID) error
	AddGenre(ctx context.Context, genre *domain.Genre) (*domain.Genre, error)
	GetAllGenres(ctx context.Context) ([]domain.Genre, error)
}
