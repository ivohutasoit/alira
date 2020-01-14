package domain

import (
	"encoding/gob"
	"time"

	"github.com/google/uuid"
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
	scope.SetColumn("ID", uuid.New().String())
	scope.SetColumn("NotBefore", time.Now())
	return nil
}

func (Identity) TableName() string {
	return "identities"
}

type NationalIdentity struct {
	model.BaseModel
	IdentityID     string    `json:"identity_id" bson:"identity_id"`
	Identity       Identity  `json:"-" bson:"-"`
	UserID         string    `json:"user_id" bson:"user_id"`
	User           User      `json:"-" gorm:"foreignkey:UserID"`
	Document       string    `json:"document" bson:"document"`
	NationID       string    `json:"nation_id" bson:"nation_id"`
	Fullname       string    `json:"fullname" bson:"fullname"`
	BirthPlace     string    `json:"birth_place" bson:"birth_place"`
	BirthDate      time.Time `json:"birth_date" bson:"birth_date" gorm:"default:null"`
	Address        string    `json:"address" bson:"address" gorm:"default:null"`
	City           string    `json:"city" bson:"city" gorm:"default:null"`
	State          string    `json:"state" bson:"state" gorm:"default:null"`
	Province       string    `json:"province" bson:"province" gorm:"default:null"`
	Country        string    `json:"country" bson:"country" gorm:"default:null"`
	PostalCode     string    `json:"postal_code" bson:"postal_code" gorm:"default:null"`
	BloodType      string    `json:"blood_type" bson:"blood_type" gorm:"default:null"`
	Religion       string    `json:"religion" bson:"religion" gorm:"default:null"`
	MarriedStatus  string    `json:"married_status" bson:"married_status" gorm:"default:null"`
	Type           string    `json:"type" bson:"type"`
	Nationality    string    `json:"nationality" bson:"nationality" gorm:"default:null"`
	IssueDate      time.Time `json:"issued_date" bson:"issued_date" gorm:"default:null"`
	ExpiryDate     time.Time `json:"expiry_date" bson:"expiry_date" gorm:"default:null"`
	RegistrationNo string    `json:"reg_no" bson:"reg_no" gorm:"default:null"`
	IssuedOffice   string    `json:"issued_office" bson:"issued_office" gorm:"default:null"`
	Nikim          string    `json:"nikim" bson:"nikim" gorm:"default:null"`
}

func (NationalIdentity) TableName() string {
	return "national_identities"
}
