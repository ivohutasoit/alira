package commerce

import (
	"time"

	alira "github.com/ivohutasoit/alira"
)

type Order struct {
	alira.Model
	StoreID         string
	Code            string
	Currency        string
	Subtotal        float64
	ServiceTax      float64
	Total           float64
	Rounding        float64
	Status          string
	PaidDate        time.Time
	PaymentMode     string
	Channel         string
	Bank            string
	ReferenceNo     string
	LocalCurrency   string
	BuyRate         float64
	SellRate        float64
	EquivalentTotal float64
}

func (Order) TableName() string {
	return "orders"
}

type OrderDetail struct {
	alira.Model
	OrderID   string
	ProductID string
	Unit      string
	Currency  string
	Price     string
	Quantity  string
	Total     string
}

func (OrderDetail) TableName() string {
	return "order_details"
}
