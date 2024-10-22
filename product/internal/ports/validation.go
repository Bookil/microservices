package ports

type Validation interface {
	ValidateGetBooksByTitle(title string) error
	ValidateGetBooksByGenre(genre string) error
	ValidateGetBooksByAuthor(author string) error
	ValidateAddBook(title string, description string, price float32, quantity uint, year uint) error
	ValidateModifyBookByID(title string, description string, price float32, quantity uint, year uint) error
	ValidateAddAuthor(name, about string) error
	ValidateAddGenre(name string) error
}
