package account

import (
	"time"

	"github.com/google/uuid"
	alira "github.com/ivohutasoit/alira"
	"github.com/jinzhu/gorm"
)

type Client struct {
	alira.Model
	Code          string
	Name          string
	SecretKey     string
	CreatedBy     string
	Domain        string
	Namespace     string
	PrivacyPolicy string
	TermOfService string
	Icon          string
	Email         string
	NotBefore     time.Time `gorm:"default:null"`
	NotAfter      time.Time `gorm:"default:null"`
}

func (client *Client) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	scope.SetColumn("NotBefore", time.Now())
	return nil
}

func (Client) TableName() string {
	return "clients"
}
