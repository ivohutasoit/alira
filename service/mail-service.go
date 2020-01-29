package service

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/ivohutasoit/alira/util"
)

type Mail struct {
	From     string
	To       []string
	Subject  string
	Template string
	Data     map[interface{}]interface{}
}

const (
	MIME = "MIME-VERSION: 1.0;\nContent-Type: text/html; charset:\"UTF-8\";\n\n"
)

func (m *Mail) NewMail(from string, to []string, subject string, template string, data map[interface{}]interface{}) *Mail {
	return &Mail{
		From:     from,
		To:       to,
		Subject:  subject,
		Template: template,
		Data:     data,
	}
}

type MailService struct{}

func (ms *MailService) Send(mail *Mail) (map[interface{}]interface{}, error) {
	mail.From = os.Getenv("SMTP_SENDER")
	var message string
	for _, to := range mail.To {
		body, _ := util.ParseMailTemplate(mail.Template, mail.Data)
		fullbody := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n%s\r\n%s", mail.From,
			to,
			mail.Subject,
			MIME,
			body,
		)
		mailSMTP := fmt.Sprintf("%s:%s", os.Getenv("SMTP_HOST"), os.Getenv("SMTP_PORT"))
		if err := smtp.SendMail(mailSMTP, smtp.PlainAuth("", mail.From,
			os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_HOST")), os.Getenv("SMTP_NAME"),
			mail.To, []byte(fullbody)); err != nil {
			return nil, err
		}
		message += fmt.Sprintf("Mail has been sent to %s \n", to)
	}
	return map[interface{}]interface{}{
		"message": message,
	}, nil
}
