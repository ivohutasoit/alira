package acount

import (
	"time"

	"github.com/google/uuid"
	alira "github.com/ivohutasoit/alira"
	"github.com/jinzhu/gorm"
)

type Token struct {
	alira.Model
	Class        string
	Referer      string
	UserID       string
	Agent        string
	AccessToken  string
	RefreshToken string
	IPAddress    string
	NotBefore    time.Time
	NotAfter     time.Time `gorm:"default:null"`
	Valid        bool      `gorm:"default:true"`
}

func (model *Token) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	scope.SetColumn("NotBefore", time.Now())
	return nil
}

func (Token) TableName() string {
	return "tokens"
}
