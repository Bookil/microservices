package api

import (
	"context"
	"strings"
	"time"

	"github.com/Bookil/microservices/auth/internal/application/core/domain"
	"github.com/Bookil/microservices/auth/internal/ports"
	"github.com/Bookil/microservices/auth/utils/hash"
	auth_manager "github.com/tahadostifam/go-auth-manager"
)

type Application struct {
	db          ports.DBPort
	authManager auth_manager.AuthManager
	hashManager *hash.HashManager
}

const (
	ResetPasswordTokenExpr     = time.Minute * 10    // 10 minutes
	AccessTokenExpr            = time.Minute * 30    // 30 minutes
	RefreshTokenExpr           = time.Hour * 24 * 14 // 2 weeks
	VerificationCodeExpr = time.Minute * 2 // 2 minutes
	LockAccountDuration        = time.Minute * 2
	MaximumFailedLoginAttempts = 3
	VerificationCodeLength = 6
)

func NewApplication(db ports.DBPort, authManager auth_manager.AuthManager, hashManager *hash.HashManager) ports.APIPort {
	return &Application{
		db:          db,
		authManager: authManager,
		hashManager: hashManager,
	}
}

func (a *Application) Register(ctx context.Context, userID domain.UserID, password string) (string, error) {
	hashedPassword, err := a.hashManager.HashPassword(password)
	if err != nil {
		return "", ErrHashingPassword
	}

	auth := domain.NewAuth(userID, hashedPassword)

	_,err = a.db.Create(ctx, auth)
	if err != nil {
		return "", ErrCreateAuthStore
	}

	verificationCode, err := a.authManager.GenerateVerificationCode(ctx,userID,VerificationCodeLength,VerificationCodeExpr)

	if err != nil {
		return "", ErrCreateEmailToken
	}

	return verificationCode, nil
}

func (a *Application) Authenticate(ctx context.Context, accessToken string) (domain.UserID, error) {
	tokenClaims, err := a.authManager.DecodeAccessToken(ctx, accessToken)
	if err != nil {
		return "", ErrAccessDenied
	}

	if len(strings.TrimSpace(tokenClaims.Payload.UUID)) == 0 {
		return "", ErrAccessDenied
	}

	return tokenClaims.Payload.UUID, nil
}

func (a *Application) VerifyEmail(ctx context.Context, userID domain.UserID,verificationCode string) error {
	isValid, err := a.authManager.CompareVerificationCode(ctx, userID,verificationCode)
	if !isValid {
		return err
	}

	_, err = a.db.VerifyEmail(ctx,userID)
	if err != nil {
		return ErrVerifyEmail
	}

	return nil
}

func (a *Application) Login(ctx context.Context, userId domain.UserID, password string) (string, string, error) {
	auth, err := a.db.GetByID(ctx, userId)
	if err != nil {
		return "", "", err
	}

	if !auth.IsEmailVerified {
		return "", "", ErrEmailNotVerified
	}

	if auth.AccountLockedUntil > time.Now().Unix() {
		return "", "", ErrAccountLocked
	}

	if auth.FailedLoginAttempts >= MaximumFailedLoginAttempts {
		_, err := a.db.LockAccount(ctx, userId, LockAccountDuration)
		if err != nil {
			return "", "", ErrLockAccount
		}
		return "", "", ErrAccountLocked
	}

	isPasswordValid := a.hashManager.CheckPasswordHash(password, auth.HashedPassword)
	if !isPasswordValid {
		_, err := a.db.IncrementFailedLoginAttempts(ctx, userId)
		if err != nil {
			return "", "", ErrIncrementFailedLoginAttempts
		}
		return "", "", ErrInvalidEmailPassword
	}

	accessToken, err := a.authManager.GenerateAccessToken(ctx, auth.UserID, AccessTokenExpr)
	if err != nil {
		return "", "", ErrGenerateToken
	}

	refreshToken, err := a.authManager.GenerateRefreshToken(ctx, auth.UserID, &auth_manager.RefreshTokenPayload{
		IPAddress:  "not implemented yet",
		UserAgent:  "not implemented yet",
		LoggedInAt: time.Duration(time.Now().UnixMilli()),
	}, RefreshTokenExpr)
	if err != nil {
		return "", "", ErrGenerateToken
	}

	_, err = a.db.ClearFailedLoginAttempts(ctx, auth.UserID)
	if err != nil {
		return "", "", ErrClearFailedLoginAttempts
	}
	return accessToken, refreshToken, nil
}

func (a *Application) ChangePassword(ctx context.Context, userID domain.UserID, newPassword string, oldPassword string) error {
	auth, err := a.db.GetByID(ctx, userID)
	if err != nil {
		return err
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

	newAccessToken, err := a.authManager.GenerateAccessToken(ctx, userID, AccessTokenExpr)
	if err != nil {
		return "", ErrGenerateToken
	}

	return newAccessToken, nil
}

func (a *Application) ResetPassword(ctx context.Context, userID string) (string, time.Duration, error) {
	auth, err := a.db.GetByID(ctx, userID)
	if err != nil {
		return "", 0, ErrNotFound
	}

	if !auth.IsEmailVerified {
		return "", 0, ErrEmailNotVerified
	}

	if auth.FailedLoginAttempts >= MaximumFailedLoginAttempts {
		return "", 0, ErrAccountLocked
	}

	resetPasswordToken, err := a.authManager.GenerateToken(ctx, auth_manager.ResetPassword, &auth_manager.TokenPayload{
		UUID:      auth.UserID,
		TokenType: auth_manager.ResetPassword,
		CreatedAt: time.Now(),
	}, ResetPasswordTokenExpr)
	if err != nil {
		return "", 0, ErrGenerateToken
	}

	return resetPasswordToken, ResetPasswordTokenExpr, nil
}

func (a *Application) SubmitResetPassword(ctx context.Context, resetPasswordToken string, newPassword string) error {
	payload, err := a.authManager.DecodeToken(ctx, resetPasswordToken, auth_manager.ResetPassword)
	if err != nil {
		return ErrAccessDenied
	}

	hashedPassword, err := a.hashManager.HashPassword(newPassword)
	if err != nil {
		return ErrHashingPassword
	}

	_, err = a.db.ChangePassword(ctx, payload.UUID, hashedPassword)
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
