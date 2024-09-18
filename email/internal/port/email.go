package port

import "time"

type SMTPPort interface {
	SendVerificationCode(email, code string) error
	SendResetPassword(url, token, email string, duration time.Duration) error
	SendWelcome(email string) error
}
