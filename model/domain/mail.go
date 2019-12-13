package domain

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
