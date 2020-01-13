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
	Code        string    `json:"code" bson:"code"`
	Name        string    `json:"name" bson:"name"`
	Category    string    `json:"category" bson:"category"`
	Description string    `json:"description" bson:"description"`
	BirthDate   time.Time `json:"birth_date" bson:"birth_date" gorm:"default:null"`
	Interest    string    `json:"interest" bson:"interest" gorm:"default:null"`
	Address     string    `json:"address" bson:"address"`
	City        string    `json:"city" bson:"city"`
	State       string    `json:"state" bson:"state"`
	Country     string    `json:"country" bson:"country"`
	PostalCode  string    `json:"postal_code" bson:"postal_code"`
	Telephone   string    `json:"telephone" bson:"telephone" gorm:"default:null"`
	Mobile      string    `json:"mobile" bson:"mobile" gorm:"default:null"`
	Website     string    `json:"website" bson:"website" gorm:"default:null"`
	Geocode     string    `json:"geocode" bson:"geocode" gorm:"default:null"`
	Longitude   float64   `json:"longitude" bson:"longitude" gorm:"default:null"`
	Latitude    float64   `json:"latitude" bson:"latitude" gorm:"default:null"`
	Finance     bool      `json:"finance" bson:"finance" gorm:"default:false"`
	Active      bool      `json:"active" bson:"active" gorm:"default:true"`
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
