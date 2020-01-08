package domain

import (
	"time"

	"github.com/ivohutasoit/alira/model"
)

type Subscribe struct {
	model.BaseModel
	Code      string    `json:"code" bson:"code"`
	UserID    string    `json:"user_id" bson:"user_id"`
	User      User      `json:"-" gorm:"foreignkey:UserID"`
	Purpose   string    `json:"purpose" bson:"purpose"`
	Signature string    `json:"signature" bson:"signature"`
	NotBefore time.Time `json:"not_before" bson:"not_before"`
	NotAfter  time.Time `json:"not_after" bson:"not_after"`
	AgreedAt  time.Time `json:"agreed_at" bson:"agreed_at"`
}

func (Subscribe) TableName() string {
	return "subscribes"
}
