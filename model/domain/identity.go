package domain

import (
	"encoding/gob"
	"time"

	"github.com/ivohutasoit/alira/model"
	"github.com/jinzhu/gorm"
)

func init() {
	gob.Register(&Identity{})
}

type Identity struct {
	model.BaseModel
	Class      string    `json:"-" bson:"-"`
	UserID     string    `json:"user_id" bson:"user_id"`
	User       User      `json:"-" gorm:"foreignkey:UserID"`
	NationID   string    `json:"nation_id" bson:"nation_id"`
	Fullname   string    `json:"fullname" bson:"fullname"`
	Code       string    `json:"code" bson:"code"`
	Email      string    `json:"email" bson:"email" gorm:"default:null"`
	Address    string    `json:"address" bson:"address" gorm:"default:null"`
	City       string    `json:"city" bson:"city" gorm:"default:null"`
	State      string    `json:"state" bson:"state" gorm:"default:null"`
	Province   string    `json:"province" bson:"province" gorm:"default:null"`
	Country    string    `json:"country" bson:"country" gorm:"default:null"`
	PostalCode string    `json:"postal_code" bson:"postal_code" gorm:"default:null"`
	NotBefore  time.Time `json:"not_before" bson:"not_before"`
	NotAfter   time.Time `json:"not_after" bson:"not_after" gorm:"default:null"`
	Valid      bool      `json:"valid" bson:"valid" gorm:"default:true"`
}

func (identity *Identity) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("NotBefore", time.Now())
	return nil
}

func (Identity) TableName() string {
	return "identities"
}
