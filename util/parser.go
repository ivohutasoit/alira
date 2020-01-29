package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"html/template"

	alira "github.com/ivohutasoit/alira"
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

func ParseResponse(data []byte, code int, response alira.Response, out interface{}) error {
	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}

	if response.Code != code {
		return errors.New("unexpected code return")
	}

	if err := json.Unmarshal([]byte(response.Data), &out); err != nil {
		return err
	}
	return nil
}
