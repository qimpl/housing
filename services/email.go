package services

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

var mg *mailgun.MailgunImpl

// TemplateValues refers to Mailgun email templates variables
type TemplateValues struct {
	IsEng     bool
	Variables map[string]string
}

func init() {
	mg = mailgun.NewMailgun(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_API_KEY"))
}

// SendEmail to a given email address using a given subject and template
func SendEmail(subject, body, recipientEmail, templateName string, templateValues *TemplateValues) {
	message := mg.NewMessage("noreply@qimpl.io", subject, body, recipientEmail)
	message.SetTemplate(templateName)
	message.AddTemplateVariable("is_eng", templateValues.IsEng)

	for name, value := range templateValues.Variables {
		message.AddTemplateVariable(name, value)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Email ID: %s Resp: %s\n", id, resp)
}
