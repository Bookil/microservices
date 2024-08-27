package ports

import "time"

type EmailPort interface {
	SendVerificationCode(email,code string)error
	SendResetPassword(url string,duration time.Duration,email string)error
	SendWellCome(email string)error
}