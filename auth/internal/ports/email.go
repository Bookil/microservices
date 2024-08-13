package ports

type EmailPort interface {
	SendVerificationEmail(email, verifyEmailRedirectUrl, verifyEmailToken string) error
}
