package message

type Mail struct {
	From     string
	To       []string
	Subject  string
	Template string
	Data     map[interface{}]interface{}
}

func (m *Mail) NewMail(from string, to []string, subject string, template string, data map[interface{}]interface{}) *Mail {
	return &Mail{
		From:     from,
		To:       to,
		Subject:  subject,
		Template: template,
		Data:     data,
	}
}
