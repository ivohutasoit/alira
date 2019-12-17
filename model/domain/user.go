package domain

import (
	"encoding/gob"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/ivohutasoit/alira/model"
	"github.com/jinzhu/gorm"
)

func init() {
	gob.Register(&User{})
	gob.Register(&Profile{})
	gob.Register(&Token{})
}

type AccessTokenClaims struct {
	jwt.StandardClaims
	UserID string
	Admin  bool
}

type RefreshTokenClaims struct {
	jwt.StandardClaims
	UserID string
	Sub    int
}

// User hold information about an user
type User struct {
	model.BaseModel
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Mobile   string `json:"mobile" bson:"mobile"`
	Avatar   string `json:"avatar" bson:"avatar"`
	Active   bool   `json:"active" bson:"active" gorm:"default:false"`
}

func (User) TableName() string {
	return "users"
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	return nil
}

type Profile struct {
	model.BaseModel
	User        User   `json:"-" gorm:"foreignkey:ID"`
	Name        string `json:"name" bson:"name"`
	FirstName   string `json:"first_name" bson:"first_name"`
	LastName    string `json:"last_name" bson:"last_name"`
	NickName    string `json:"nick_name" bson:"nick_name"`
	Gender      string `json:"gender" bson:"gender"`
	Description string `json:"description" bson:"description"`
}

func (Profile) TableName() string {
	return "profiles"
}

type Token struct {
	model.BaseModel
	Class        string    `json:"-" bson:"-"`
	Referer      string    `json:"referer" bson:"referer"`
	UserID       string    `json:"user_id" bson:"user_id"`
	User         User      `json:"-" gorm:"foreignkey:UserID"`
	Agent        string    `json:"agent" bson:"agent"`
	AccessToken  string    `json:"access_token bson:"access_token"`
	RefreshToken string    `json:"refresh_token" bson:"refresh_token"`
	IPAddress    string    `json:"ip_address" bson:"ip_address"`
	NotBefore    time.Time `json:"not_before" bson:"not_before"`
	NotAfter     time.Time `json:"not_after" bson:"not_after"`
	Valid        bool      `json:"valid" bson:"valid" gorm:"default:true"`
}

func (Token) TableName() string {
	return "tokens"
}

func (token *Token) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	return nil
}
