package email

import "github.com/mailgun/mailgun-go/v4"

type Email interface {
	SendNotification(address string, text string) error
}

type EmailService struct {
	client *mailgun.MailgunImpl
}

func New(domain string, key string) Email {
	client := mailgun.NewMailgun(domain, key)
	email := &EmailService{client: client}
	return email
}

func (e *EmailService) SendNotification(address string, text string) error {
	return nil
}
