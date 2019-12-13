package util

import (
	"bytes"
	"html/template"
)

func ParseMailTemplate(name string, data interface{}) (interface{}, error) {
	tmpl, err := template.ParseFiles(name)
	if err != nil {
		return nil, err
	}

	buff := new(bytes.Buffer)
	if err := tmpl.Execute(buff, data); err != nil {
		return nil, err
	}
	return buff.String(), nil
}
