package ports

import (
	"time"

	"github.com/Bookil/microservices/auth/internal/application/core/domain"
)

type APIPort interface {
	Register(userID domain.UserID, email, password, verifyEmailRedirectUrl string) (string, error)
	Authenticate(accessToken string) domain.UserID
	VerifyEmail(verifyEmailToken string) error
	Login(email, password string) (string, string, error)
	ChangePassword(userID domain.UserID, newPassword string, oldPassword string) error
	RefreshToken(userID domain.UserID, refreshToken string) (string, error)
	SendResetPasswordEmail(email string, resetPasswordRedirectUrl string) (string, time.Duration, error)
	SubmitResetPassword(token string, newPassword string) error
	DeleteAccount(userID domain.UserID, password string) error
}
