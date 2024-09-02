package domain

import (
	"fmt"
	"time"

	"github.com/Bookil/microservices/user/utils/random"
)

type (
	UserID = string

	User struct {
		UserID    UserID `gorm:"unique"`
		FirstName string
		LastName  string
		Email     string `gorm:"unique"`
		CreatedAt time.Time
		UpdatedAt time.Time
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
