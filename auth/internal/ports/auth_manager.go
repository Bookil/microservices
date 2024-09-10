package ports

import (
	"context"

	"github.com/Bookil/microservices/auth/internal/application/core/domain"
)

type AuthManager interface {
	GenerateAccessToken(ctx context.Context, userID domain.UserID) (accessToken string, _ error)
	DecodeAccessToken(ctx context.Context, accessToken string) (*domain.AccessTokenClaims, error)
	GenerateRefreshToken(ctx context.Context, userID domain.UserID) (refreshToken string, _ error)
	DecodeRefreshToken(ctx context.Context, userID domain.UserID, RefreshToken string) (*domain.RefreshTokenClaims, error)
	GenerateResetPasswordToken(ctx context.Context, userID domain.UserID) (resetPasswordToken string, err error)
	DecodeResetPasswordToken(ctx context.Context, resetPasswordToken string) (*domain.ResetPasswordTokenClaims, error)
	GenerateVerificationCode(ctx context.Context, key string) (verificationCode string, _ error)
	CompareVerificationCode(ctx context.Context, userID domain.UserID, verificationCode string) (bool, error)
}
