package domain

import (
	"encoding/gob"
	"time"

	"github.com/ivohutasoit/alira/model"
)

func init() {
	gob.Register(&Community{})
	gob.Register(&CommunityMember{})
}

type Community struct {
	model.BaseModel
	Name        string `json:"name" bson:"name"`
	Category    string `json:"category" bson:"category"`
	Description string `json:"description" bson:"description"`
	BirthDate   time.Time `json:"birth_date" bson:"birth_date"`
	Interest    string `json:"interest" bson:"interest"`
	Street      string `json:"street" bson:"street"` 
	City        string `json:"city" bson:"city"`
	Country     string `json:"country" bson:"country"`
	PostalCode  string `json:"postal_code" bson:"postal_code"`
	Telephone   string `json:"telephone" bson:"telephone"`
	Mobile      string `json:"mobile" bson:"mobile"`
	Website     string `json:"website" bson:"website"`
	Geocode     string `json:"geocode" bson:"geocode"`
	Longitude   float64 `json:"longitude" bson:"longitude"`
	Latitude    float64 `json:"latitude" bson:"latitude"`
	Finance     bool `json:"finance" bson:"finance"`
	Active      bool `json:"active" bson:"active"`
}

func (Community) TableName() string {
	return "communities"
}

type CommunityMember struct {
	model.BaseModel
	CommunityID  string    `json:"community_id" bson:"community_id"`
	Community    Community `json:"-" gorm:"foreignkey:CommunityID"`
	UserID       string    `json:"user_id" bson:"user_id"`
	User         User      `json:"-" gorm:"foreignkey:UserID"`
	Creator      bool      `json:"creator" bson:"creator"`
	Admin        bool      `json:"admin" bson:"admin"`
	AdditionInfo string    `json:"additional_info" bson:"additional_info"`
}

func (CommunityMember) TableName() string {
	return "community_members"
}
