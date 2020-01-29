package account

import (
	"time"

	alira "github.com/ivohutasoit/alira"
)

// User hold information about an user
type User struct {
	alira.Model
	Username       string
	Email          string
	Mobile         string
	Avatar         string
	Active         bool
	FirstTimeLogin bool
	UsePin         bool `grom:"default:false"`
	Pin            string
	Profile        Profile `gorm:"foreignkey:ID"`
}

func (User) TableName() string {
	return "users"
}

type Profile struct {
	ID              string `gorm:"primary_key"`
	CreatedBy       string `gorm:"default:null"`
	CreatedAt       time.Time
	UpdatedBy       string `gorm:"default:null"`
	UpdatedAt       time.Time
	DeletedBy       string     `gorm:"default:null"`
	DeletedAt       *time.Time `sql:"index"`
	FirstName       string
	MiddleName      string
	LastName        string
	NickName        string
	Gender          string
	SecondaryMobile string
}

func (Profile) TableName() string {
	return "profiles"
}
