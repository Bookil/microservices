package ports

type APIPort interface {
	SendVerificationCode(email, name, code string) error
	SendResetPassword(email, name, url, expiry string) error
	SendWelcome(email, name string) error
}
