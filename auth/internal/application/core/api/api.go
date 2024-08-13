package api

import (
	"time"

	"github.com/Bookil/microservices/auth/internal/application/core/domain"
	"github.com/Bookil/microservices/auth/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func (a *Application) Register(userID domain.UserID, password string) (string, error) {
	panic("unimplemented")
}

func (a *Application) Authenticate(accessToken string) domain.UserID {
	panic("unimplemented")

}

func (a *Application) VerifyEmail(verifyEmailToken string) error {
	panic("unimplemented")
}

func (a *Application) Login(email, password string) (string, string, error) {
	panic("unimplemented")

}
func (a *Application) ChangePassword(userID domain.UserID, newPassword string, oldPassword string) error {
	panic("unimplemented")

}
func (a *Application) RefreshToken(userID domain.UserID, refreshToken string) (string, error) {
	panic("unimplemented")

}
func (a *Application) SendResetPassword(email string, resetPasswordRedirectUrl string) (string, time.Duration, error) {
	panic("unimplemented")

}
func (a *Application) SubmitResetPassword(token string, newPassword string) error {
	panic("unimplemented")

}
func (a *Application) DeleteAccount(userID domain.UserID, password string) error {
	panic("unimplemented")

}
