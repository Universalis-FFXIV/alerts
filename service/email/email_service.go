package email

import (
	"bytes"
	"context"
	_ "embed"
	"os"
	"text/template"
	"time"

	"github.com/Universalis-FFXIV/alerts/model"
	"github.com/Universalis-FFXIV/alerts/service/common"
	"github.com/mailgun/mailgun-go/v4"
)

//go:embed email_template.html
var emailTemplate string

type emailService struct {
	client        *mailgun.MailgunImpl
	et            *template.Template
	senderAddress string
}

// New creates a new Mailgun-backed NotificationService.
func New(domain string, key string) (common.NotificationService, error) {
	client := mailgun.NewMailgun(domain, key)

	et, err := template.New("universalis_email_template").Parse(emailTemplate)
	if err != nil {
		return nil, err
	}

	email := &emailService{
		client:        client,
		et:            et,
		senderAddress: os.Getenv("UNIVERSALIS_ALERTS_EMAIL_ADDRESS"),
	}

	return email, nil
}

func (e *emailService) SendNotification(address string, notification *model.Notification) error {
	subject := "Alert triggered for " + notification.ItemName

	var body bytes.Buffer
	err := e.et.Execute(&body, notification)
	if err != nil {
		return err
	}

	message := e.client.NewMessage(e.senderAddress, subject, body.String(), address)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, _, err = e.client.Send(ctx, message)

	return err
}
