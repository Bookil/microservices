package domain

type EmailMessage struct {
	Template string
	Receiver []string
	Args     map[string]interface{}
	Subject  string
}

func NewEmailMessage(template, subject string, args map[string]interface{}, receiver []string) *EmailMessage {
	return &EmailMessage{
		Template: template,
		Subject:  subject,
		Args:     args,
		Receiver: receiver,
	}
}
