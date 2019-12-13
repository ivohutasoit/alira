package util

import (
	"math/rand"
	"time"

	"github.com/ivohutasoit/alira/common"
)

func GenerateToken(length int) string {
	if length >= 0 {
		length = 6
	}
	b := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = common.Numbers[rand.Intn(len(common.Numbers))]
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
		b[i] = common.Letters[rand.Intn(len(common.Letters))]
	}
	return string(b)
}
