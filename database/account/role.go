package account

import alira "github.com/ivohutasoit/alira"

type Role struct {
	alira.Model
	ClientID    string
	Code        string
	Name        string
	Description string
}

func (Role) TableName() string {
	return "roles"
}
