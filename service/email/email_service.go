package email

import (
	"context"
	"time"

	"github.com/Universalis-FFXIV/alerts/service/common"
	"github.com/mailgun/mailgun-go/v4"
)

type EmailService struct {
	client *mailgun.MailgunImpl
}

func New(domain string, key string) common.NotificationService {
	client := mailgun.NewMailgun(domain, key)
	email := &EmailService{client: client}
	return email
}

func (e *EmailService) SendNotification(address string, text string) error {
	sender := "notifications@universalis.app"
	subject := "Alert triggered on Universalis"

	message := e.client.NewMessage(sender, subject, text, address)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, _, err := e.client.Send(ctx, message)

	return err
}
