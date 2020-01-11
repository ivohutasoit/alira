package service

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/ivohutasoit/alira/model/domain"
	"github.com/ivohutasoit/alira/util"
)

type MailService struct{}

func (ms *MailService) Send(mail *domain.Mail) (map[interface{}]interface{}, error) {
	main.From = os.Getenv("SMTP_SENDER")
	var message string
	for _, to := range mail.To {
		body, _ := util.ParseMailTemplate(mail.Template, mail.Data)
		fullbody := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n%s\r\n%s", mail.From,
			to,
			mail.Subject,
			domain.MIME,
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
