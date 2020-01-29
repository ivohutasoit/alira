package account

import (
	"time"

	"github.com/google/uuid"
	alira "github.com/ivohutasoit/alira"
	"github.com/jinzhu/gorm"
)

type Identity struct {
	alira.Model
	Class     string
	UserID    string
	Code      string
	Email     string `gorm:"default:null"`
	NotBefore time.Time
	NotAfter  time.Time `gorm:"default:null"`
	Valid     bool      `gorm:"default:true"`
}

func (model *Identity) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	scope.SetColumn("NotBefore", time.Now())
	return nil
}

func (Identity) TableName() string {
	return "identities"
}

type NationalIdentity struct {
	alira.Model
	IdentityID     string
	UserID         string
	Document       string
	NationID       string
	Fullname       string
	BirthPlace     string
	BirthDate      time.Time `gorm:"default:null"`
	Address        string    `gorm:"default:null"`
	City           string    `gorm:"default:null"`
	State          string    `gorm:"default:null"`
	Province       string    `gorm:"default:null"`
	Country        string    `gorm:"default:null"`
	PostalCode     string    `gorm:"default:null"`
	BloodType      string    `gorm:"default:null"`
	Religion       string    `gorm:"default:null"`
	MarriedStatus  string    `gorm:"default:null"`
	Type           string    `gorm:"default:null"`
	Nationality    string    `gorm:"default:null"`
	IssueDate      time.Time `gorm:"default:null"`
	ExpiryDate     time.Time `gorm:"default:null"`
	RegistrationNo string    `gorm:"default:null"`
	IssuedOffice   string    `gorm:"default:null"`
	Nikim          string    `gorm:"default:null"`
}

func (NationalIdentity) TableName() string {
	return "national_identities"
}
