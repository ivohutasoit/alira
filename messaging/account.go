package messaging

// UserProfile godoc
type UserProfile struct {
	ID              string `form:"id" json:"id" bson:"id" xml:"id"`
	Username        string `form:"username" json:"username" bson:"username" xml:"username"`
	Email           string `form:"email" json:"email" bson:"email" xml:"email"`
	PrimaryMobile   string `form:"primary_mobile" json:"primary_mobile" bson:"primary_mobile" xml:"primary_mobile"`
	SecondaryMobile string `form:"secondary_mobile" json:"secondary_mobile" bson:"secondary_mobile" xml:"secondary_mobile"`
	FirstName       string `form:"first_name" json:"first_name" bson:"first_name" xml:"first_name"`
	MiddleName      string `form:"middle_name" json:"middle_name" bson:"middle_name" xml:"middle_name"`
	LastName        string `form:"last_name" json:"last_name" bson:"last_name" xml:"last_name"`
	Avatar          string `form:"avatar" json:"avatar" bson:"avatar" xml:"avatar"`
	Active          bool   `form:"active" json:"active" bson:"active" xml:"active"`
	UsePin          bool   `form:"use_pin" json:"use_pin" bson:"use_pin" xml:"use_pin"`
}
