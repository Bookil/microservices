package auth

import (
	"context"
	"fmt"
	"time"

	authv1 "github.com/Bookil/Bookil-Proto/gen/golang/auth/v1"
	"github.com/Bookil/microservices/user/config"
	"github.com/Bookil/microservices/user/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	auth authv1.AuthServiceClient
}

func generateURL(url *config.Auth) string {
	return fmt.Sprintf("%s:%d", url.Host, url.Port)
}

func NewAdapter(url *config.Auth) (*Adapter, error) {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(generateURL(url), opts...)
	if err != nil {
		return nil, err
	}

	client := authv1.NewAuthServiceClient(conn)

	return &Adapter{auth: client}, nil
}

func (a *Adapter) Register(ctx context.Context, userID domain.UserID, password string) (string, error) {
	response, err := a.auth.Register(ctx, &authv1.RegisterRequest{
		UserId:   userID,
		Password: password,
	})
	if err != nil {
		return "", err
	}

	return response.VerificationCode, nil
}

func (a *Adapter) Authenticate(ctx context.Context, accessToken string) (domain.UserID, error) {
	response, err := a.auth.Authentication(ctx, &authv1.AuthenticationRequest{
		AccessToken: accessToken,
	})
	if err != nil {
		return "", err
	}

	return response.UserId, nil
}

func (a *Adapter) VerifyEmail(ctx context.Context, userID domain.UserID, code string) error {
	_, err := a.auth.VerifyEmail(ctx, &authv1.VerifyEmailRequest{
		UserId:           userID,
		VerificationCode: code,
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *Adapter) Login(ctx context.Context, userID, password string) (string, string, error) {
	response, err := a.auth.Login(ctx, &authv1.LoginRequest{
		UserId:   userID,
		Password: password,
	})
	if err != nil {
		return "", "", err
	}

	return response.AccessToken, response.RefreshToken, nil
}

func (a *Adapter) ChangePassword(ctx context.Context, userID domain.UserID, newPassword string, oldPassword string) error {
	_, err := a.auth.ChangePassword(ctx, &authv1.ChangePasswordRequest{
		UserId:      userID,
		NewPassword: newPassword,
		OldPassword: oldPassword,
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *Adapter) ResetPassword(ctx context.Context, userID string) (string, time.Duration, error) {
	response, err := a.auth.ResetPassword(ctx, &authv1.ResetPasswordRequest{
		UserId: userID,
	})
	if err != nil {
		return "", 0, err
	}

	return response.ResetPasswordToken, response.Timeout.AsDuration(), nil
}

func (a *Adapter) SubmitResetPassword(ctx context.Context, token string, newPassword string) error {
	_, err := a.auth.SubmitResetPassword(ctx, &authv1.SubmitResetPasswordRequest{
		ResetPasswordToken: token,
		NewPassword:        newPassword,
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *Adapter) DeleteAccount(ctx context.Context, userID domain.UserID, password string) error {
	_, err := a.auth.DeleteAccount(ctx, &authv1.DeleteAccountRequest{
		UserId:   userID,
		Password: password,
	})
	if err != nil {
		return err
	}

	return nil
}
