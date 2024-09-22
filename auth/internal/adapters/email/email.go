package email

import (
	"context"
	"fmt"
	"log"
	"time"

	emailv1 "github.com/Bookil/Bookil-Proto/gen/golang/email/v1"
	"github.com/Bookil/microservices/auth/config"
	"github.com/Bookil/microservices/auth/internal/adapters/grpc/interceptor"
	retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	email emailv1.EmailServiceClient
}


func generateURL(url *config.EmailService) string {
	return fmt.Sprintf("%s:%d", url.Host, url.Port)
}

func NewAdapter(url *config.EmailService) (*Adapter, error) {
	var opts []grpc.DialOption

	if config.CurrentEnv == config.Production{
		cb := gobreaker.NewCircuitBreaker(
			gobreaker.Settings{
				Name:        "email",
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

	client := emailv1.NewEmailServiceClient(conn)

	return &Adapter{email: client}, nil
}

func (a *Adapter) SendResetPassword(ctx context.Context, email, name, url string, expiry time.Duration) error {
	req := &emailv1.SendResetPasswordRequest{
		Email:  email,
		Name:   name,
		Url:    url,
		Expiry: expiry.String(),
	}

	_, err := a.email.SendResetPassword(ctx, req)

	return err
}

func (a *Adapter) SendVerificationCode(ctx context.Context, email, name, code string) error {
	req := &emailv1.SendVerificationCodeRequest{
		Email:            email,
		Name:             name,
		VerificationCode: code,
	}

	_, err := a.email.SendVerificationCode(ctx, req)

	return err
}

func (a *Adapter) SendWelcome(ctx context.Context, email, name string) error {
	req := &emailv1.SendWelcomeRequest{
		Email: email,
		Name:  name,
	}

	_, err := a.email.SendWelcome(ctx, req)

	return err
}
