package ports

import "time"

type EmailPort interface {
	SendVerificationCode(email, code string) error
	SendResetPassword(url,token,email string, duration time.Duration) error
	SendWellCome(email string) error
}
