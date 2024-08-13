package api

import (
	"context"
	"time"

	"github.com/Bookil/microservices/auth/internal/application/core/domain"
	"github.com/Bookil/microservices/auth/internal/ports"
	"github.com/Bookil/microservices/auth/utils/hash"
	auth_manager "github.com/tahadostifam/go-auth-manager"
)

type Application struct {
	db ports.DBPort
	email ports.EmailPort
	authManager  auth_manager.AuthManager
	hashManager *hash.HashManager
}

const (
	VerifyEmailTokenExpr       = time.Minute * 5     // 5 minutes
	ResetPasswordTokenExpr     = time.Minute * 10    // 10 minutes
	AccessTokenExpr            = time.Minute * 30    // 30 minutes
	RefreshTokenExpr           = time.Hour * 24 * 14 // 2 weeks
	LockAccountDuration        = time.Second * 5
	MaximumFailedLoginAttempts = 5
)


func (a *Application) Register(userID domain.UserID,email,password,verifyEmailRedirectUrl string) (string, error) {
	passwordHash, err := a.hashManager.HashPassword(password)
        if err != nil {
                return "", ErrHashingPassword
        }

        auth := domain.NewAuth(userID, passwordHash)

        err = a.db.Create(auth)
        if err != nil {
                return "", ErrCreateAuthStore
        }

		ctx := context.Background()

        verifyEmailToken, err := a.authManager.GenerateToken(
                ctx, auth_manager.VerifyEmail,
                &auth_manager.TokenPayload{
                        UUID:      userID,
                        TokenType: auth_manager.VerifyEmail,
                        CreatedAt: time.Now(),
                },
                VerifyEmailTokenExpr,
        )
        if err != nil {
                return "", ErrCreateEmailToken
        }

        err = a.email.SendVerificationEmail(email, verifyEmailRedirectUrl, verifyEmailToken)

		//TODO:error handling
		if err != nil {
                return "",err
        }

        return verifyEmailToken, nil
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
