package domain

type Genre struct {
	ID    uint    `gorm:"primaryKey"`
	Name  string  `gorm:"unique"`
	Books []*Book `gorm:"many2many:book_genres;"`
}

func NewGenre(name string) *Genre {
	return &Genre{
		Name: name,
	}
}
