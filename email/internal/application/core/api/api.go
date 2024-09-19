package api

import "email/internal/ports"

// TODO:Handle Errors
type Application struct {
	smtp ports.SMTPPort
}

func NewApplication(smtp ports.SMTPPort) *Application {
	return &Application{
		smtp: smtp,
	}
}

func (a *Application) SendVerificationCode(email, name, code string) error {
	err := a.smtp.SendVerificationCode(email, name, code)

	return err
}

func (a *Application) SendResetPassword(email, name, url, expiry string) error {
	err := a.smtp.SendResetPassword(email, name, url, expiry)

	return err
}

func (a *Application) SendWelcome(email, name string) error {
	err := a.smtp.SendWelcome(email, name)

	return err
}
