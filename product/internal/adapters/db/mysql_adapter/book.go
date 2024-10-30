package mysql_adapter

import (
	"context"

	"product/internal/application/core/domain"
)

func (a *Adapter) AddBook(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	for _, author := range book.Authors {
		_, err := a.GetAuthorByID(ctx, author.ID)
		if err != nil {
			return nil, err
		}
	}
	for _, genre := range book.Genres {
		_, err := a.GetGenreByID(ctx, genre.ID)
		if err != nil {
			return nil, err
		}
	}

	err := a.db.Create(book).Error
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (a *Adapter) GetAllBooks(ctx context.Context) ([]domain.Book, error) {
	var books []domain.Book

	err := a.db.WithContext(ctx).Preload("Genres").Preload("Authors").Find(&books).Error

	return books, err
}

func (a *Adapter) GetBookByID(ctx context.Context, ID domain.BookID) (*domain.Book, error) {
	book := &domain.Book{}

	err := a.db.WithContext(ctx).Preload("Genres").Preload("Authors").Find(book, ID).First(book).Error
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (a *Adapter) GetBooksByTitle(ctx context.Context, title string) ([]domain.Book, error) {
	var books []domain.Book

	err := a.db.WithContext(ctx).Preload("Genres").Preload("Authors").Where("title = ?", title).Find(&books).Error

	if len(books) == 0 {
		return nil, ErrBookNotFound
	}

	return books, err
}

func (a *Adapter) GetBooksByAuthor(ctx context.Context, authorName string) ([]domain.Book, error) {
	var books []domain.Book

	err := a.db.WithContext(ctx).
		Joins("JOIN book_authors ON book_authors.book_id = books.id").
		Joins("JOIN authors ON authors.id = book_authors.author_id").
		Where("authors.name = ?", authorName).
		Preload("Authors").
		Preload("Genres").
		Find(&books).Error

	if len(books) == 0 {
		return nil, ErrAuthorNotFound
	}

	return books, err
}

func (a *Adapter) GetBooksByGenre(ctx context.Context, genreName string) ([]domain.Book, error) {
	var books []domain.Book

	err := a.db.WithContext(ctx).
		Joins("JOIN book_genres ON book_genres.book_id = books.id").
		Joins("JOIN genres ON genres.id = book_genres.genre_id").
		Where("genres.name = ?", genreName).
		Preload("Genres").Preload("Authors").
		Find(&books).Error

	if len(books) == 0 {
		return nil, ErrGenreNotFound
	}

	return books, err
}

func (a *Adapter) ModifyBookByID(ctx context.Context, ID domain.BookID, book *domain.Book) (*domain.Book, error) {
	savedBook, err := a.GetBookByID(ctx, ID)
	if err != nil {
		return nil, err
	}

	savedBook.Title = book.Title
	savedBook.Description = book.Description
	savedBook.Year = book.Year
	savedBook.Quantity = book.Quantity
	savedBook.Price = book.Price

	err = a.db.Model(savedBook).Association("Authors").Replace(book.Authors)
	if err != nil {
		return nil, err
	}

	err = a.db.Model(savedBook).Association("Genres").Replace(book.Genres)
	if err != nil {
		return nil, err
	}

	err = a.db.Save(savedBook).Error
	if err != nil {
		return nil, err
	}

	return savedBook, nil
}

func (a *Adapter) DeleteBookByID(ctx context.Context, ID domain.BookID) error {
	book, err := a.GetBookByID(ctx, ID)
	if err != nil {
		return err
	}

	err = a.db.Model(&book).Association("Authors").Clear()
	if err != nil {
		return err
	}

	err = a.db.Model(&book).Association("Genres").Clear()
	if err != nil {
		return err
	}

	err = a.db.Delete(&book).Error
	if err != nil {
		return err
	}

	return nil
}
