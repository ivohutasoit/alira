package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
)

// Encrypt function can be used to encrypt plain test using secret key
// First parameter is mandatory
func Encrypt(args ...interface{}) (string, error) {
	var plain []byte
	var key []byte = []byte(os.Getenv("APP.SECRET.KEY"))
	if 1 > len(args) {
		return "", errors.New("not enough parameters")
	}

	for i, p := range args {
		switch i {
		case 0: //plain text
			param, ok := p.(string)
			if !ok {
				return "", errors.New("plain text parameter not type string")
			}
			plain = []byte(param)
		case 1:
			param, ok := p.(string)
			if !ok {
				return "", errors.New("screet key parameter not type string")
			}
			key = []byte(param)
		default:
			return "", errors.New("wrong parameter count")
		}
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(string(plain)))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plain)

	return base64.URLEncoding.EncodeToString(cipherText), nil
}

// Decrypt function can be used to decrupt encrypted test using secret key
// First parameter is mandatory
func Decrypt(args ...interface{}) (string, error) {
	var cipherText []byte
	var err error
	var key []byte = []byte(os.Getenv("APP.SECRET.KEY"))
	if 1 > len(args) {
		return "", errors.New("not enough parameters")
	}

	for i, p := range args {
		switch i {
		case 0: //plain text
			param, ok := p.(string)
			if !ok {
				return "", errors.New("plain text parameter not type string")
			}
			cipherText, err = base64.URLEncoding.DecodeString(param)
			if err != nil {
				return "", err
			}
		case 1:
			param, ok := p.(string)
			if !ok {
				return "", errors.New("screet key parameter not type string")
			}
			key = []byte(param)
		default:
			return "", errors.New("wrong parameter count")
		}
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", errors.New("invalid cipher text")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return fmt.Sprintf("%s", cipherText), nil
}
