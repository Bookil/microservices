package ports

import "product/internal/application/core/domain"

type Validation interface {
	ValidateGetBooksByTitle(title string) error
	ValidateGetBooksByGenre(genre string) error
	ValidateGetBooksByAuthor(author string) error
	ValidateAddBook(book *domain.Book) error
	ValidateModifyBookByID(book *domain.Book) error
	ValidateAddAuthor(name, about string) error
	ValidateAddGenre(name string) error
}
