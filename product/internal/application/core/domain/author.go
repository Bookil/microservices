package domain

type Author struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	About string
	Books []*Book `gorm:"many2many:book_authors;"`
}

func NewAuthor(name, about string) *Author {
	return &Author{
		Name:  name,
		About: about,
	}
}
