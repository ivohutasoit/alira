package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/ivohutasoit/alira/constant"
	"github.com/ivohutasoit/alira/model"
	"github.com/ivohutasoit/alira/model/domain"
)

func GenerateToken(length int) string {
	if length <= 0 {
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
	if length <= 0 {
		length = 16
	}
	b := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = constant.Letters[rand.Intn(len(constant.Letters))]
	}
	return string(b)
}

func GenerateNationalCode(args ...interface{}) string {
	//Example Nomor NIK : 1234567890ABCDEF
	//12 nomor merupakan kode provinsi
	//34 nomor merupakan kode kotamadya atau kabupaten kota
	//56 nomor kode kecamatan
	//78 nomor tanggal lahir
	//90 nomor bulan lahir
	//AB nomor tahun lahir
	//CDEF nomor registrasi kependudukan

	// Country code 2 characters ID=62
	// BirthDate 8 characters format yyyyMMDD
	// Sequence 4 characters
	date, _ := args[1].(time.Time)

	code := fmt.Sprintf("%d%d%02d%02d", 62, date.Year(), date.Month(), date.Day())

	var identities []domain.Identity
	model.GetDatabase().Where("code LIKE ?", code+"%").Find(&identities).Order("code DESC")

	var identity domain.Identity
	if identities != nil {
		identity = identities[0]
	}

	var seq int
	if identity.BaseModel.ID != "" {
		if n, err := strconv.Atoi(identity.Code[10:len(identity.Code)]); err == nil {
			seq = n + 1
		} else {
			fmt.Println(seq, "is not an integer.")
		}
	}

	return fmt.Sprintf("%s%06d", code, seq)
}
