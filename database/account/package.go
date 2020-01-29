package account

import (
	"time"

	"github.com/google/uuid"
	alira "github.com/ivohutasoit/alira"
	"github.com/jinzhu/gorm"
)

type Subscription struct {
	alira.Model
	Class       string `gorm:"default:USERTYPE"`
	Code        string
	Name        string
	Description string
	Applicator  string `gorm:"defaulr:null"`
	Subscriber  string `sql:"index"`
	Signature   string
	NotBefore   time.Time
	NotAfter    time.Time
	AgreedAt    time.Time
}

func (model *Subscription) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	scope.SetColumn("NotBefore", time.Now())
	scope.SetColumn("AgreedAt", time.Now())
	return nil
}

func (Subscription) TableName() string {
	return "subscriptions"
}
