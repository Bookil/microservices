package port

type SMTPPort interface {
	SendVerificationCode(email,name,code string) error
	SendResetPassword(email,name,url, expiry string) error
	SendWelcome(fullName,email string) error
}
