package domain

import "github.com/ivohutasoit/alira/model"

type Wallet struct {
	model.BaseModel
	Class       string  `json:"-" bson:"-" gorm:"default:USERWALLET"`
	OwnerID     string  `json:"owner_id" bson:"owner_id" sql:"index"`
	Type        string  `json:"type" bson:"type" gorm:"default:CONVENTIONAL"`
	Balance     float64 `json:"balance" bson:"balance" gorm:"default:0"`
	UsageLimit  float64 `json:"usage_limit" bson:"usage_limit" gorm:"default:0"`
	Limit       float64 `json:"limit" bson:"limit" gorm:"default:0"`
	Acummulated bool    `json:"accumulated" bson:"accumulated" gorm:"default:false"`
}

func (Wallet) TableName() string {
	return "wallets"
}
