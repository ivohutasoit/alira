package domain

import (
	"encoding/gob"
	"time"
	
	"github.com/ivohutasoit/alira/model"
)

func init() {
	gob.Register(&Identity{})
}

type Identity struct {
	model.BaseModel
	Class string `json:"-" bson:"-"`
	UserID string `json:"user_id" bson:"user_id"`
	User User `json:"-" gorm:"foreignkey:UserID"`
	Code string `json:"code" bson:"code"`
	Email string `json:"email" bson:"email"`
	NotBefore time.Time `json:"not_before" bson:"not_before"`
	NotAfter time.Time `json:"not_after" bson:"not_after" gorm:"default:null"`
	Valid bool `json:"valid" bson:"valid" gorm:"default:true"`
}

func (Identity) TableName() string {
	return "identities"
}
