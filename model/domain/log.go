package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/ivohutasoit/alira/model"
	"github.com/jinzhu/gorm"
)

type Log struct {
	model.BaseModel
	Class      string    `json:"-" bson:"-" gorm:"default:FINANCIAL"` // FINANCIAL, TRANSACTION
	ActionDate time.Time `json:"action_date" bson:"action_date"`
	ActionBy   string    `json:"action_by" bson:"action_by" sql:"index"`
	Status     string    `json:"status" bson:"status"`
	Error      string    `json:"error" bson:"error"`
}

func (log *Log) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	scope.SetColumn("ActionDate", time.Now())
	return nil
}

func (Log) TableName() string {
	return "logs"
}

type FinancialLog struct {
	ID             string  `json:"id" bson:"id" gorm:"primary_key"`
	Log            Log     `json:"-" gorm:"foreignkey:ID"`
	Type           string  `json:"type" bson:"type" gorm:"default:DEBIT"` // DEBIT: Money Out from wallet; CREDIT: Money In to wallet
	ServiceID      string  `json:"service_id" bson:"service_id"`          // TOPUP, PURCHASE, PAYMENT
	Amount         float64 `json:"amount" bson:"amount"`
	Fee            float64 `json:"fee" bson:"fee"`
	ReferenceNo    string  `json:"reference_no" bson:"reference_no"`
	GLNo           string  `json:"gl_no" bson:"gl_no"`
	Channel        string  `json:"channel" bson:"channel"`
	BankCode       string  `json:"bank_code" bson:"bank_code"`
	AdditionalInfo string  `json:"additional_info" bson:"additional_info"`
}

func (log *FinancialLog) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	return nil
}

func (FinancialLog) TableName() string {
	return "financial_logs"
}
