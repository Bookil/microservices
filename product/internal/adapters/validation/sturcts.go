package validation

type getBooksByTitle struct {
	Title string `validate:"required"`
}

type getBooksByAuthor struct {
	Author string `validate:"required"`
}

type getBooksByGenre struct {
	Genre string `validate:"required"`
}

type bookInput struct {
	Title       string  `validate:"required"`
	Description string  `validate:"required"`
	Price       float32 `validate:"required"`
	Quantity    uint    `validate:"required"`
	Year        uint    `validate:"required"`
}

type addAuthor struct {
	Name  string `validate:"required"`
	About string `validate:"required"`
}

type addGenre struct {
	Name string `validate:"required"`
}
