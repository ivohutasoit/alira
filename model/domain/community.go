package domain

import (
	"github.com/ivohutasoit/alira/model"
	"time"
)

type Community struct {
	model.BaseModel
	Name        string
	Category    string
	Description string
	BirthDate   time.Time
	Interest    string
	Street      string
	City        string
	Country     string
	PostalCode  string
	Telephone   string
	Mobile      string
	Website     string
	Geocode     string
	Longitude   float64
	Latitude    float64
	Finance     bool
	Active      bool
}

func (Community) TableName() string {
	return "communities"
}

type CommunityUser struct {
	model.BaseModel
	CommunityID  string
	Community    Community
	UserID       string
	User         User
	AdditionInfo string
}

func (CommunityUser) TableName() string {
	return "community_users"
}
