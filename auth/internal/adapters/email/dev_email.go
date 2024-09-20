package email

import (
	"context"
	"log"
	"time"
)

type DevAdapter struct{}

func NewDevEmailAdapter() *DevAdapter {
	return &DevAdapter{}
}

func (a *DevAdapter) SendResetPassword(ctx context.Context, email, name, url string, duration time.Duration) error {
	log.Println("email sent")
	return nil
}

func (a *DevAdapter) SendVerificationCode(ctx context.Context, email, name, code string) error {
	log.Println("email sent")
	return nil
}

func (a *DevAdapter) SendWelcome(ctx context.Context, email, name string) error {
	log.Println("email sent")
	return nil
}
