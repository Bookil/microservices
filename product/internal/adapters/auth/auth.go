package auth

import (
	"context"
	"fmt"
	"log"
	"time"

	"product/config"
	"product/internal/adapters/grpc/interceptor"

	retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"

	authv1 "github.com/Bookil/Bookil-Proto/gen/golang/auth/v1"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	auth authv1.AuthServiceClient
}

func generateURL(url *config.AuthService) string {
	return fmt.Sprintf("%s:%d", url.Host, url.Port)
}

func NewAdapter(url *config.AuthService) (*Adapter, error) {
	var opts []grpc.DialOption

	if config.CurrentEnv == config.Production {
		cb := gobreaker.NewCircuitBreaker(
			gobreaker.Settings{
				Name:        "auth",
				MaxRequests: 5,
				Timeout:     5 * time.Second,
				ReadyToTrip: func(counts gobreaker.Counts) bool {
					failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
					return failureRatio > 0.5
				},
				OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
					log.Println("Circcuit breaker state changed:", name, from, to)
				},
			},
		)

		opts = append(opts, grpc.WithUnaryInterceptor(interceptor.CircuitBreakerInterceptor(cb)))

		opts = append(opts, grpc.WithUnaryInterceptor(
			retry.UnaryClientInterceptor(
				retry.WithCodes(codes.ResourceExhausted, codes.Unavailable),
				retry.WithMax(3),
				retry.WithPerRetryTimeout(time.Second),
				retry.WithBackoff(retry.BackoffLinear(2*time.Second)),
			),
		))
	}

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(generateURL(url), opts...)
	if err != nil {
		return nil, err
	}

	client := authv1.NewAuthServiceClient(conn)

	return &Adapter{auth: client}, nil
}

func (a *Adapter) Authenticate(ctx context.Context, accessToken string) (string, error) {
	response, err := a.auth.Authentication(ctx, &authv1.AuthenticationRequest{
		AccessToken: accessToken,
	})
	if err != nil {
		return "", err
	}

	return response.UserId, nil
}
