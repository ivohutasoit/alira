package domain

import (
	"time"

	alira "github.com/ivohutasoit/alira"
)

type Community struct {
	alira.Model
	Code            string
	Name            string
	Avatar          string
	Background      string
	Category        string
	Description     string
	LongDescription string
	BirthDate       time.Time `gorm:"default:null"`
	Interest        string    `gorm:"default:null"`
	Address         string
	City            string
	State           string
	Country         string
	PostalCode      string
	Telephone       string  `gorm:"default:null"`
	Mobile          string  `gorm:"default:null"`
	Website         string  `gorm:"default:null"`
	Geocode         string  `gorm:"default:null"`
	Longitude       float64 `gorm:"default:null"`
	Latitude        float64 `gorm:"default:null"`
	Finance         string  `gorm:"default:'NONE'"`
	Active          bool    `gorm:"default:true"`
}

func (Community) TableName() string {
	return "communities"
}

type CommunityMember struct {
	alira.Model
	CommunityID  string
	UserID       string
	Creator      bool
	Admin        bool
	JoinBy       string `gorm:"default:'INVITE'"`
	Approved     bool   `gorm:"default:false"`
	AdditionInfo string
}

func (CommunityMember) TableName() string {
	return "community_members"
}
