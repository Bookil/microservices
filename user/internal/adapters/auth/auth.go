package auth

import (
	"context"
	"fmt"

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

func (a *Adapter) Authenticate(ctx context.Context, accessToken string) (domain.UserID, error) {
	response, err := a.auth.Authentication(ctx, &authv1.AuthenticationRequest{
		AccessToken: accessToken,
	})
	if err != nil {
		return "", err
	}

	return response.UserId, nil
}

func (a *Adapter) ChangePassword(ctx context.Context, userID domain.UserID, oldPassword string, newPassword string) error {
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
