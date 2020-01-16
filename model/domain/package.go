package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/ivohutasoit/alira/model"
	"github.com/jinzhu/gorm"
)

type Subscribe struct {
	model.BaseModel
	Code         string    `json:"code" bson:"code"`
	Class        string    `json:"-" bson:"-" gorm:"default:USERSUBSCRIPTION"`
	SubscriberID string    `json:"subscriber_id" bson:"subscriber_id" sql:"index"`
	Purpose      string    `json:"purpose" bson:"purpose"`
	Signature    string    `json:"signature" bson:"signature"`
	NotBefore    time.Time `json:"not_before" bson:"not_before"`
	NotAfter     time.Time `json:"not_after" bson:"not_after"`
	AgreedAt     time.Time `json:"agreed_at" bson:"agreed_at"`
}

func (subscribe *Subscribe) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	scope.SetColumn("NotBefore", time.Now())
	scope.SetColumn("AgreedAt", time.Now())
	return nil
}

func (Subscribe) TableName() string {
	return "subscribes"
}
