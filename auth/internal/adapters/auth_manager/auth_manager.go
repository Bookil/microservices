package auth_manager

import (
	"context"
	"time"

	"github.com/Bookil/microservices/auth/config"
	"github.com/Bookil/microservices/auth/internal/application/core/domain"
	"github.com/go-redis/redis/v8"
	auth_manager "github.com/tahadostifam/go-auth-manager"
)

const (
	ResetPasswordTokenExpr = time.Second * 360   // 5 minutes
	AccessTokenExpr        = time.Minute * 30    // 30 minutes
	RefreshTokenExpr       = time.Hour * 24 * 14 // 2 weeks
	VerificationCodeExpr   = time.Minute * 2     // 2 minutes
	VerificationCodeLength = 6
)

type Adapter struct {
	authManger auth_manager.AuthManager
}

func NewAdapter(redisClient *redis.Client, jwtConfigs config.JWT) *Adapter {
	authManger := auth_manager.NewAuthManager(redisClient, auth_manager.AuthManagerOpts{
		PrivateKey: jwtConfigs.SecretKey,
	})

	return &Adapter{
		authManger: authManger,
	}
}

func (a *Adapter) GenerateAccessToken(ctx context.Context, userID domain.UserID,role domain.Role) (accessToken string, _ error) {
	accessToken, err := a.authManger.GenerateAccessToken(ctx, userID,role, AccessTokenExpr)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (a *Adapter) DecodeAccessToken(ctx context.Context, accessToken string) (*domain.AccessTokenClaims, error) {
	accessTokenPayload, err := a.authManger.DecodeAccessToken(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	accessTokenClaims := &domain.AccessTokenClaims{
		UserID: accessTokenPayload.Payload.UUID,
		Role: accessTokenPayload.Payload.Role,
	}

	return accessTokenClaims, nil
}

func (a *Adapter) GenerateRefreshToken(ctx context.Context, userID domain.UserID) (refreshToken string, _ error) {
	refreshTokenClaims, err := a.authManger.GenerateRefreshToken(ctx, userID, &auth_manager.RefreshTokenPayload{
		IPAddress:  "not implemented yet",
		UserAgent:  "not implemented yet",
		LoggedInAt: time.Duration(time.Now().UnixMilli()),
	},
		RefreshTokenExpr)
	if err != nil {
		return "", err
	}

	return refreshTokenClaims, nil
}

func (a *Adapter) DecodeRefreshToken(ctx context.Context, userID domain.UserID, RefreshToken string) (*domain.RefreshTokenClaims, error) {
	refreshTokenPayload, err := a.authManger.DecodeRefreshToken(ctx, userID, RefreshToken)
	if err != nil {
		return nil, err
	}

	refreshTokenClaims := &domain.RefreshTokenClaims{
		IPAddress:  refreshTokenPayload.IPAddress,
		UserAgent:  refreshTokenPayload.UserAgent,
		LoggedInAt: refreshTokenPayload.LoggedInAt,
	}

	return refreshTokenClaims, err
}

func (a *Adapter) GenerateResetPasswordToken(ctx context.Context, userID domain.UserID) (resetPasswordToken string, _ error) {
	resetPasswordToken, err := a.authManger.GenerateToken(ctx, auth_manager.ResetPassword, &auth_manager.TokenPayload{UUID: userID}, ResetPasswordTokenExpr)
	if err != nil {
		return "", err
	}

	return resetPasswordToken, nil
}

func (a *Adapter) DecodeResetPasswordToken(ctx context.Context, resetPasswordToken string) (*domain.ResetPasswordTokenClaims, error) {
	resetPasswordTokePayload, err := a.authManger.DecodeToken(ctx, resetPasswordToken, auth_manager.ResetPassword)
	if err != nil {
		return nil, err
	}
	resetPasswordTokeClaims := &domain.ResetPasswordTokenClaims{
		UserID:   resetPasswordTokePayload.UUID,
		CreateAt: resetPasswordTokePayload.CreatedAt,
	}

	return resetPasswordTokeClaims, nil
}

func (a *Adapter) GenerateVerificationCode(ctx context.Context, key string) (verificationCode string, _ error) {
	verificationCode, err := a.authManger.GenerateVerificationCode(ctx, key, VerificationCodeLength, VerificationCodeExpr)
	if err != nil {
		return "", err
	}

	return verificationCode, nil
}

func (a *Adapter) CompareVerificationCode(ctx context.Context, userID domain.UserID, verificationCode string) (bool, error) {
	isEqual, err := a.authManger.CompareVerificationCode(ctx, userID, verificationCode)
	if err != nil {
		return false, err
	}

	return isEqual, nil
}
