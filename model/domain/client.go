package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/ivohutasoit/alira/model"
	"github.com/jinzhu/gorm"
)

type Client struct {
	model.BaseModel
	Code          string    `form:"code" json:"code" bson:"code" xml:"code"`
	Name          string    `form:"name" json:"name" bson:"name" xml:"name"`
	SecretKey     string    `form:"secret_key" json:"secret_key" bson:"secret_key" xml:"secret_key"`
	CreatedBy     string    `form:"created_by" json:"created_by" bson:"created_by" xml:"created_by"`
	Domain        string    `form:"domain" json:"domain" bson:"domain" xml:"domain"`
	Namespace     string    `form:"namespace" json:"namespace" bson:"namespace" xml:"namespace"`
	PrivacyPolicy string    `form:"privacy_policy" json:"privacy_policy" bson:"privacy_policy" xml:"privacy_policy"`
	TermOfService string    `form:"term_of_service" json:"term_of_service" bson:"term_of_service" xml:"term_of_service"`
	Icon          string    `form:"icon" json:"icon" bson:"icon" xml:"icon"`
	Email         string    `form:"email" json:"email" bson:"email" xml:"email"`
	NotBefore     time.Time `form:"not_before" json:"not_before" bson:"not_before" xml:"not_before" gorm:"default:null"`
	NotAfter      time.Time `form:"not_after" json:"not_after" bson:"not_after" xml:"not_after" gorm:"default:null"`
}

func (client *Client) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	scope.SetColumn("NotBefore", time.Now())
	return nil
}

func (Client) TableName() string {
	return "Client"
}
