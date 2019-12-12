package domain

type Mail struct {
	From    string
	To      []string
	Subject string
	Body    string
}

const (
	MIME = "MIME-VERSION: 1.0;\nContent-Type: text/html; charset:\"UTF-8\";\n\n"
)

func (m *Mail) NewMailer(to []string, subject string, body string) *Mail {
	return &Mail{
		To:      to,
		Subject: subject,
		Body:    body,
	}
}
