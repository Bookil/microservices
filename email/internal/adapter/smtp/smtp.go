package smtp

import (
	"fmt"
	"net/smtp"
	"path/filepath"
	"strings"
	"text/template"

	"email/config"
	"email/internal/application/core/domain"
)

var templatesPath = fmt.Sprintf("%s/%s/%s/%s/%s", config.ProjectRootPath, "internal", "adapter", "smtp", "templates")

type EmailOtp struct {
	configs *config.SMTP
}

func NewSMTPAdapter(configs *config.SMTP) *EmailOtp {
	return &EmailOtp{configs}
}

func (s *EmailOtp) sendEmail(msg *domain.EmailMessage) error {
	template, err := renderTemplate(msg.Template, msg.Args)
	if err != nil {
		return err
	}

	emailMessage := fmt.Sprintf("Subject: %s\r\n"+
		"Content-Type: text/html; charset=UTF-8\r\n"+
		"\r\n"+template, msg.Subject)

	url := fmt.Sprintf("%s:%d", s.configs.Host, s.configs.Port)
	auth := smtp.PlainAuth("", s.configs.SenderEmail, s.configs.AppPassword, s.configs.Host)
	err = smtp.SendMail(url, auth, s.configs.SenderEmail, msg.Receiver, []byte(emailMessage))
	if err != nil {
		return err
	}

	return nil
}

func (s *EmailOtp) SendVerificationCode(recipientEmail, name, code string) error {
	msg := domain.NewEmailMessage(
		"verify_email.html",
		"Verify Email",
		map[string]interface{}{"name": name, "code": code},
		[]string{recipientEmail},
	)

	err := s.sendEmail(msg)
	if err != nil {
		return err
	}

	return nil
}

func (s *EmailOtp) SendResetPassword(recipientEmail, name, url, expiry string) error {
	msg := domain.NewEmailMessage(
		"reset_password.html",
		"Reset Password",
		map[string]interface{}{"name": name, "url": url, "expiry": expiry},
		[]string{recipientEmail},
	)

	err := s.sendEmail(msg)
	if err != nil {
		return err
	}
	return nil
}

func (s *EmailOtp) SendWelcome(email, fullName string) error {
	msg := domain.NewEmailMessage(
		"welcome.html",
		"Welcome",
		map[string]interface{}{"fullName": fullName},
		[]string{email},
	)

	err := s.sendEmail(msg)
	if err != nil {
		return err
	}
	return nil
}

func renderTemplate(tmpl string, data interface{}) (string, error) {
	tmplPath := filepath.Join(templatesPath, tmpl)
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		return "", err
	}

	var buf strings.Builder
	err = t.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
