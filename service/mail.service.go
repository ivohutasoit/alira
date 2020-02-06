package service

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/ivohutasoit/alira/message"
	"github.com/ivohutasoit/alira/util"
)

const (
	MIME = "MIME-VERSION: 1.0;\nContent-Type: text/html; charset:\"UTF-8\";\n\n"
)

type Mail struct{}

func (s *Mail) Send(mail *message.Mail) (map[interface{}]interface{}, error) {
	mail.From = os.Getenv("SMTP_SENDER")
	var message string
	parser := &util.Parser{}
	for _, to := range mail.To {
		body, _ := parser.MailTemplate(mail.Template, mail.Data)
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
