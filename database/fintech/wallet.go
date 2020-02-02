package domain

import alira "github.com/ivohutasoit/alira"

type Wallet struct {
	alira.Model
	Class       string  `gorm:"default:'USERWALLET'"`
	OwnerID     string  `sql:"index"`
	Type        string  `gorm:"default:'CONVENTIONAL'"`
	Balance     float64 `gorm:"default:0"`
	UsageLimit  float64 `gorm:"default:0"`
	Limit       float64 `gorm:"default:0"`
	Acummulated bool    `gorm:"default:false"`
}

func (Wallet) TableName() string {
	return "wallets"
}
