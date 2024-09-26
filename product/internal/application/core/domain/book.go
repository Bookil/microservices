package domain

import (
	"time"
)

type (
	BookID = uint

)

type Book struct {
	ID          BookID `gorm:"primaryKey"`
	Quantity    uint
	Author      string
	Title       string
	Year        uint
	Description string
	Price       float64
	Genres      []*Genre `gorm:"many2many:book_genres;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewBook(title, author, description string,quantity uint,year uint, Genre []*Genre, price float64) *Book {
	return &Book{
		Author: author,
		Title:       title,
		Genres:       Genre,
		Description: description,
		Quantity: quantity,
		Year: year,
		Price: price,
	}
}
