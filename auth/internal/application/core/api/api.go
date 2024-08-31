package api

import (
	"context"
	"strings"
	"time"

	"github.com/Bookil/microservices/auth/internal/adapters/auth_manager"
	"github.com/Bookil/microservices/auth/internal/application/core/domain"
	"github.com/Bookil/microservices/auth/internal/ports"
)

type Application struct {
	db          ports.DBPort
	user        ports.UserPorts
	email       ports.EmailPort
	authManager ports.AuthManager
	hashManager ports.HashManager
}

const (
	LockAccountDuration        = time.Minute * 2
	MaximumFailedLoginAttempts = 3
)

func NewApplication(db ports.DBPort, user ports.UserPorts, email ports.EmailPort, authManager ports.AuthManager, hashManager ports.HashManager) *Application {
	return &Application{
		db:          db,
		user:        user,
		email:       email,
		authManager: authManager,
		hashManager: hashManager,
	}
}

func (a *Application) Register(ctx context.Context, firstName, lastName, email, password string) (userID domain.UserID, _ error) {
	userID, err := a.user.Register(ctx, firstName, lastName, email)
	if err != nil {
		return "", err
	}

	hashedPassword, err := a.hashManager.HashPassword(password)
	if err != nil {
		return "", ErrHashingPassword
	}

	auth := domain.NewAuth(userID, hashedPassword)

	_, err = a.db.Create(ctx, auth)
	if err != nil {
		return "", ErrCreateAuthStore
	}

	verificationCode, err := a.authManager.GenerateVerificationCode(ctx, userID)
	if err != nil {
		return "", ErrGenerateVerificationCode
	}

	err = a.email.SendVerificationCode(email, verificationCode)
	if err != nil {
		return "", err
	}

	return userID, nil
}

func (a *Application) VerifyEmail(ctx context.Context, userID domain.UserID, verificationCode string) error {
	isValid, err := a.authManager.CompareVerificationCode(ctx, userID, verificationCode)
	if !isValid || err != nil {
		return ErrVerifyEmail
	}

	_, err = a.db.VerifyEmail(ctx, userID)
	if err != nil {
		return ErrVerifyEmail
	}

	return nil
}

func (a *Application) Login(ctx context.Context, email, password string) (accessToken string, refreshToken string, _ error) {
	userID, err := a.user.GetUserIDByEmail(ctx, email)
	if err != nil {
		return "", "", err
	}

	auth, err := a.db.GetByID(ctx, userID)
	if err != nil {
		return "", "", ErrNotFound
	}

	if !auth.IsEmailVerified {
		return "", "", ErrEmailNotVerified
	}

	if auth.AccountLockedUntil > time.Now().Unix() {
		return "", "", ErrAccountLocked
	}

	if auth.FailedLoginAttempts >= MaximumFailedLoginAttempts {
		_, err := a.db.LockAccount(ctx, userID, LockAccountDuration)
		if err != nil {
			return "", "", ErrLockAccount
		}
		return "", "", ErrAccountLocked
	}

	isPasswordValid := a.hashManager.CheckPasswordHash(password, auth.HashedPassword)
	if !isPasswordValid {
		_, err := a.db.IncrementFailedLoginAttempts(ctx, userID)
		if err != nil {
			return "", "", ErrIncrementFailedLoginAttempts
		}
		return "", "", ErrInvalidEmailPassword
	}

	accessToken, err = a.authManager.GenerateAccessToken(ctx, auth.UserID)
	if err != nil {
		return "", "", ErrGenerateToken
	}

	refreshToken, err = a.authManager.GenerateRefreshToken(ctx, auth.UserID)
	if err != nil {
		return "", "", ErrGenerateToken
	}

	_, err = a.db.ClearFailedLoginAttempts(ctx, auth.UserID)
	if err != nil {
		return "", "", ErrClearFailedLoginAttempts
	}

	err = a.email.SendWelcome(email)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (a *Application) Authenticate(ctx context.Context, accessToken string) (domain.UserID, error) {
	tokenClaims, err := a.authManager.DecodeAccessToken(ctx, accessToken)
	if err != nil {
		return "", ErrAccessDenied
	}

	if len(strings.TrimSpace(tokenClaims.UserID)) == 0 {
		return "", ErrAccessDenied
	}

	return tokenClaims.UserID, nil
}

func (a *Application) ChangePassword(ctx context.Context, userID domain.UserID, oldPassword string, newPassword string) error {
	auth, err := a.db.GetByID(ctx, userID)
	if err != nil {
		return ErrNotFound
	}

	isPasswordValid := a.hashManager.CheckPasswordHash(oldPassword, auth.HashedPassword)
	if !isPasswordValid {
		return ErrInvalidPassword
	}

	newHashedPassword, err := a.hashManager.HashPassword(newPassword)
	if err != nil {
		return ErrHashingPassword
	}

	_, err = a.db.ChangePassword(ctx, userID, newHashedPassword)
	if err != nil {
		return ErrChangePassword
	}

	return nil
}

func (a *Application) RefreshToken(ctx context.Context, userID domain.UserID, refreshToken string) (string, error) {
	_, err := a.authManager.DecodeRefreshToken(ctx, userID, refreshToken)
	if err != nil {
		return "", ErrAccessDenied
	}

	_, err = a.db.GetByID(ctx, userID)
	if err != nil {
		return "", ErrAccessDenied
	}

	newAccessToken, err := a.authManager.GenerateAccessToken(ctx, userID)
	if err != nil {
		return "", ErrGenerateToken
	}

	return newAccessToken, nil
}

func (a *Application) ResetPassword(ctx context.Context, email string) error {
	userID, err := a.user.GetUserIDByEmail(ctx, email)
	if err != nil {
		return err
	}

	auth, err := a.db.GetByID(ctx, userID)
	if err != nil {
		return ErrNotFound
	}

	if !auth.IsEmailVerified {
		return ErrEmailNotVerified
	}

	if auth.FailedLoginAttempts >= MaximumFailedLoginAttempts {
		return ErrAccountLocked
	}

	resetPasswordToken, err := a.authManager.GenerateResetPasswordToken(ctx, userID)
	if err != nil {
		return ErrGenerateToken
	}

	err = a.email.SendResetPassword("example.com", resetPasswordToken, email, auth_manager.ResetPasswordTokenExpr)
	if err != nil {
		return err
	}

	return nil
}

func (a *Application) SubmitResetPassword(ctx context.Context, resetPasswordToken string, newPassword string) error {
	resetPasswordTokenCliams, err := a.authManager.DecodeResetPasswordToken(ctx, resetPasswordToken)
	if err != nil {
		return ErrAccessDenied
	}

	hashedPassword, err := a.hashManager.HashPassword(newPassword)
	if err != nil {
		return ErrHashingPassword
	}

	_, err = a.db.ChangePassword(ctx, resetPasswordTokenCliams.UserID, hashedPassword)
	if err != nil {
		return ErrChangePassword
	}

	return nil
}

func (a *Application) DeleteAccount(ctx context.Context, userID domain.UserID, password string) error {
	auth, err := a.db.GetByID(ctx, userID)
	if err != nil {
		return ErrNotFound
	}

	validPassword := a.hashManager.CheckPasswordHash(password, auth.HashedPassword)
	if !validPassword {
		return ErrInvalidPassword
	}

	err = a.db.DeleteByID(ctx, userID)
	if err != nil {
		return ErrDelete
	}

	return nil
}
