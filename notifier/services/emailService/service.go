package emailService

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type MailService interface {
	SendEmail(target, text string) error
}

func NewMailService() MailService {
	return &SendGridMailService{}
}

type SendGridMailService struct {
}

func (s *SendGridMailService) SendEmail(target, text string) error {
	from := mail.NewEmail("DoTenX", "sender")
	subject := "ao email"
	to := mail.NewEmail("for User", target)
	plainTextContent := text
	htmlContent := "<strong>" + text + "</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient("apiKey")
	_, err := client.Send(message)
	return err
}
