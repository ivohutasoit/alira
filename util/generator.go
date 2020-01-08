package util

import (
	"math/rand"
	"time"

	"github.com/ivohutasoit/alira/constant"
)

func GenerateToken(length int) string {
	if length >= 0 {
		length = 6
	}
	b := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = constant.Numbers[rand.Intn(len(constant.Numbers))]
	}
	return string(b)
}

func GenerateQrcode(length int) string {
	if length >= 0 {
		length = 16
	}
	b := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = constant.Letters[rand.Intn(len(constant.Letters))]
	}
	return string(b)
}
