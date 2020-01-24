package account

import alira "github.com/ivohutasoit/alira/model"

type Role struct {
	alira.BaseModel
	Code string `form:"code" json:"code" bson:"code" xml:"code"`
	Name string `form:"name" json:"name" bson:"name" xml:"name"`
}

func (Role) TableName() string {
	return "roles"
}

type UserRole struct {
	alira.BaseModel
	UserID string `form:"user_id" json:"user_id" bson:"user_id" xml:"user_id"`
	RoleID string `form:"role_id" json:"role_id" bson:"role_id" xml:"role_id"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
