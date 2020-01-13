package domain

import (
	"encoding/gob"
	"time"

	"github.com/ivohutasoit/alira/model"
	"github.com/jinzhu/gorm"
)

func init() {
	gob.Register(&Identity{})
	gob.Register(&NationalIdentity{})
}

type Identity struct {
	model.BaseModel
	Class     string    `json:"-" bson:"-"`
	UserID    string    `json:"user_id" bson:"user_id"`
	User      User      `json:"-" gorm:"foreignkey:UserID"`
	Code      string    `json:"code" bson:"code"`
	Email     string    `json:"email" bson:"email" gorm:"default:null"`
	NotBefore time.Time `json:"not_before" bson:"not_before"`
	NotAfter  time.Time `json:"not_after" bson:"not_after" gorm:"default:null"`
	Valid     bool      `json:"valid" bson:"valid" gorm:"default:true"`
}

func (identity *Identity) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("NotBefore", time.Now())
	return nil
}

func (Identity) TableName() string {
	return "identities"
}

type NationalIdentity struct {
	model.BaseModel
	IdentityID    string   `json:"identity_id" bson:"identity_id"`
	Identity      Identity `json:"-" bson:"-"`
	UserID        string   `json:"user_id" bson:"user_id"`
	User          User     `json:"-" gorm:"foreignkey:UserID"`
	Class     string    `json:"-" bson:"-"`
	Type 		  string   `json:"type" bson:"type"`
	NationID      string   `json:"nation_id" bson:"nation_id"`
	Fullname      string   `json:"fullname" bson:"fullname"`
	Address       string   `json:"address" bson:"address" gorm:"default:null"`
	City          string   `json:"city" bson:"city" gorm:"default:null"`
	State         string   `json:"state" bson:"state" gorm:"default:null"`
	Province      string   `json:"province" bson:"province" gorm:"default:null"`
	Country       string   `json:"country" bson:"country" gorm:"default:null"`
	PostalCode    string   `json:"postal_code" bson:"postal_code" gorm:"default:null"`
	BloodType     string   `json:"blood_type" bson:"blood_type" gorm:"default:null"`
	Religion      string   `json:"religion" bson:"religion" gorm:"default:null"`
	MarriedStatus string   `json:"married_status" bson:"married_status" gorm:"default:null"`
	IssueDate time.Time
	ExpiryDate time.Time
	RegistrationNo string
	IssuedOffice string
	Nikim string
}

func (NationalIdentity) TableName() string {
	return "national_identities"
}
