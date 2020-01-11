package util

import (
	"crypto/tls"
	"errors"
	"fmt"
	"os"
	"strings"
)

func GenerateUrl(args ...interface{}) (string, error) {
	if 3 > len(args) {
		return "", errors.New("not enough parameters")
	}

	var protocol, host, uri string
	var err error
	encrypted := false
	for i, p := range args {
		switch i {
		case 0:
			param, ok := p.(*tls.ConnectionState)
			if !ok {
				return "", errors.New("connection state type required")
			}
			protocol = "http://"
			if param != nil {
				protocol = "https://"
			}
		case 1:
			param, ok := p.(string)
			if !ok {
				return "", errors.New("string type required")
			}
			host = param
		case 2:
			param, ok := p.(string)
			if !ok {
				return "", errors.New("string type required")
			}
			uri = param
		case 3:
			param, ok := p.(bool)
			if !ok {
				return "", errors.New("bool type required")
			}
			encrypted = param
		default:
			return "", errors.New("wrong parameter count")
		}
	}
	url := fmt.Sprintf("%s%s%s", protocol, host, uri)
	if encrypted == true {
		url, err = Encrypt(strings.TrimSpace(url), os.Getenv("APP.SECRET.KEY"))
		if err != nil {
			fmt.Println(err)
			return "", errors.New("could not encrypt")
		}
	}

	return url, nil
}
