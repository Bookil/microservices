package domain

type (
	UserID = string

	User struct{
		UserID UserID
		FirstName string
		LastName string
		Email string
	}
)

