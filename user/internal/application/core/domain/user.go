package domain

import (
	"fmt"

	"github.com/Bookil/microservices/user/utils/random"
)

type (
	UserID = string

	User struct {
		UserID    UserID
		FirstName string
		LastName  string
		Email     string
	}
)

func NewUser(firstName, lastName, email string) *User {
	return &User{
		UserID:    fmt.Sprintf("%d", random.GenerateUserID()),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
}
