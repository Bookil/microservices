package ports

import "github.com/Bookil/microservices/user/internal/application/core/domain"

type Validation interface {
	ValidateRegisterInputs(firstName, lastName, email string) error
	ValidateChangePasswordInputs(userID domain.UserID, oldPassword string, newPassword string) error
	ValidationUpdateInputs(userID domain.UserID, firstName, lastName string) error
	ValidateDeleteAccountInputs(userID domain.UserID, password string) error
}
