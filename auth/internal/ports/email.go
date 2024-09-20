package ports

import (
	"context"
	"time"
)

type EmailPort interface {
	SendVerificationCode(ctx context.Context, email, name, code string) error
	SendResetPassword(ctx context.Context, email, name, url string, duration time.Duration) error
	SendWelcome(ctx context.Context, email, name string) error
}
