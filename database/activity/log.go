package actiity

import (
	"time"

	"github.com/google/uuid"
	alira "github.com/ivohutasoit/alira"
	"github.com/jinzhu/gorm"
)

type Log struct {
	alira.Model
	Class      string `gorm:"default:FINANCIAL"` // FINANCIAL, TRANSACTION
	ActionDate time.Time
	ActionBy   string `sql:"index"`
	Status     string
	Error      string
}

func (model *Log) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	scope.SetColumn("ActionDate", time.Now())
	return nil
}

func (Log) TableName() string {
	return "logs"
}

type FinancialLog struct {
	ID             string `gorm:"primary_key"`
	Log            Log    `gorm:"foreignkey:ID"`
	Type           string `gorm:"default:DEBIT"` // DEBIT: Money Out from wallet; CREDIT: Money In to wallet
	ServiceID      string // TOPUP, PURCHASE, PAYMENT
	Amount         float64
	Fee            float64
	ReferenceNo    string
	GLNo           string
	Channel        string
	BankCode       string
	AdditionalInfo string
}

func (model *FinancialLog) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	return nil
}

func (FinancialLog) TableName() string {
	return "financial_logs"
}
