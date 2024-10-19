package api

import (
	"context"

	"product/internal/application/core/domain"
	"product/internal/ports"
)

type Application struct {
	cart ports.CartPort
	DB   ports.DBPort
}

func NewApplication(cart ports.CartPort, db ports.DBPort) *Application {
	return &Application{
		cart: cart,
		DB:   db,
	}
}

func (a *Application) GetAllBooks(ctx context.Context) ([]domain.Book, error) {
	books, err := a.DB.GetAllBooks(ctx)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (a *Application) GetBookByID(ctx context.Context, ID domain.BookID) (*domain.Book, error) {
	book, err := a.DB.GetBookByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (a *Application) GetBooksByTitle(ctx context.Context, title string) ([]domain.Book, error) {
	books, err := a.DB.GetBooksByTitle(ctx, title)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (a *Application) GetBooksByAuthor(ctx context.Context, authorName string) ([]domain.Book, error) {
	books, err := a.DB.GetBooksByAuthor(ctx, authorName)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (a *Application) GetBooksByGenre(ctx context.Context, genreName string) ([]domain.Book, error) {
	books, err := a.DB.GetBooksByGenre(ctx, genreName)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (a *Application) AddBook(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	book, err := a.DB.AddBook(ctx, book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (a *Application) ModifyBookByID(ctx context.Context, ID domain.BookID, book *domain.Book) (*domain.Book, error) {
	book, err := a.DB.ModifyBookByID(ctx, ID, book)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (a *Application) DeleteBookByID(ctx context.Context, ID domain.BookID) error {
	err := a.DB.DeleteBookByID(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}

func (a *Application) AddBookToCart(ctx context.Context, ID domain.BookID, userID string) error {
	book, err := a.DB.GetBookByID(ctx, ID)
	if err != nil {
		return err
	}

	err = a.cart.AddBookToCart(ctx, book.ID, userID)

	return err
}

func (a *Application) DeleteBookFromCartByID(ctx context.Context, bookID domain.BookID, cartID uint) error {
	err := a.cart.DeleteBookFromCartByID(ctx, bookID, cartID)

	return err
}

func (a *Application) AddAuthor(ctx context.Context, author *domain.Author) (*domain.Author, error) {
	author, err := a.DB.AddAuthor(ctx, author)
	if err != nil {
		return nil, err
	}

	return author, nil
}

func (a *Application) GetAuthorByID(ctx context.Context, ID uint) (*domain.Author, error) {
	author, err := a.DB.GetAuthorByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return author, nil
}

func (a *Application) DeleteAuthorByID(ctx context.Context, ID uint) error {
	err := a.DB.DeleteAuthorByID(ctx, ID)
	
	return err
}

func (a *Application) GetAllAuthors(ctx context.Context) ([]domain.Author, error) {
	authors, err := a.DB.GetAllAuthors(ctx)
	if err != nil {
		return nil, err
	}

	return authors, nil
}

func (a *Application) AddGenre(ctx context.Context, genre *domain.Genre) (*domain.Genre, error) {
	genre, err := a.DB.AddGenre(ctx, genre)
	if err != nil {
		return nil, err
	}

	return genre, nil
}

func (a *Application) GetGenreByID(ctx context.Context, ID uint) (*domain.Genre, error) {
	genre, err := a.DB.GetGenreByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	return genre, nil
}

func (a *Application) DeleteGenreByID(ctx context.Context, ID uint) error {
	err := a.DB.DeleteGenreByID(ctx, ID)
	
	return err
}

func (a *Application) GetAllGenres(ctx context.Context) ([]domain.Genre, error) {
	genres, err := a.DB.GetAllGenres(ctx)
	if err != nil {
		return nil, err
	}

	return genres, nil
}
