package user

import (
	"context"
	"fmt"
	"time"

	userv1 "github.com/Bookil/Bookil-Proto/gen/golang/user/v1"
	"github.com/Bookil/microservices/auth/config"
	"github.com/Bookil/microservices/auth/internal/application/core/domain"

	retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	user userv1.UserServiceClient
}

func generateURL(url *config.UserService) string {
	return fmt.Sprintf("%s:%d", url.Host, url.Port)
}

func NewAdapter(url *config.UserService) (*Adapter, error) {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithUnaryInterceptor(
		retry.UnaryClientInterceptor(
			retry.WithCodes(codes.ResourceExhausted, codes.Unavailable),
			retry.WithMax(3),
			retry.WithPerRetryTimeout(time.Second),
			retry.WithBackoff(retry.BackoffLinear(2*time.Second)),
		),
	))

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(generateURL(url), opts...)
	if err != nil {
		return nil, err
	}

	client := userv1.NewUserServiceClient(conn)

	return &Adapter{user: client}, nil
}

func (a *Adapter) Register(ctx context.Context, firstName, lastName, email string) (domain.UserID, error) {
	response, err := a.user.Register(ctx, &userv1.RegisterRequest{
		FisrtName: firstName,
		LastName:  lastName,
		Email:     email,
	})

	if err != nil {
		return "", err
	}

	return response.UserId, nil
}

func (a *Adapter) GetUserIDByEmail(ctx context.Context, email string) (domain.UserID, error) {
	response, err := a.user.GetUserIDByEmail(ctx, &userv1.GetUserIDByEmailRequest{
		Email: email,
	})

	if err != nil {
		return "", err
	}

	return response.UserId, nil
}
