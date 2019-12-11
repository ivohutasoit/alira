package domain

import (
	"encoding/gob"

	"github.com/dgrijalva/jwt-go"
	"github.com/ivohutasoit/alira/model"
)

func init() {
	gob.Register(&User{})
	gob.Register(&UserProfile{})
}

type AccessToken struct {
	jwt.StandardClaims
	Admin bool
}

type RefreshToken struct {
	jwt.StandardClaims
	Sub int
}

// User hold information about an user
type User struct {
	model.BaseModel
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Mobile   string `json:"mobile" bson:"mobile"`
	Avatar   string `json:"avatar" bson:"avatar"`
	Active   bool   `json:"active" bson:"active"`
}

func (User) TableName() string {
	return "users"
}

type UserProfile struct {
	model.BaseModel
	Name        string `json:"name" bson:"name"`
	FirstName   string `json:"first_name" bson:"first_name"`
	LastName    string `json:"last_name" bson:"last_name"`
	NickName    string `json:"nick_name" bson:"nick_name"`
	Description string `json:"description" bson:"description"`
}

func (UserProfile) TableName() string {
	return "profiles"
}
