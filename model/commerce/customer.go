package commerce

import (
	"time"

	alira "github.com/ivohutasoit/alira/model"
)

type Customer struct {
	alira.BaseModel
	Class     string `form:"class" json:"class" bson:"class" xml:"class" gorm:"default:SHOPOWNER"` // DISTRIBUTOR
	CreatedBy string `form:"-" json:"-" bson:"-" xml:"-"`
	UpdatedBy string `form:"-" json:"-" bson:"-" xml:"-"`
	Code      string `form:"code" json:"code" bson:"code" xml:"code"`
	Status    string `form:"status" json:"status" bson:"status" xml:"status" grom:"default:INACTIVE"`
	Payment   bool   `form:"payment" json:"payment" bson:"payment" xml:"payment" grom:"default:false"`
}

func (Customer) TableName() string {
	return "customers"
}

type CustomerUser struct {
	alira.BaseModel
	CustomerID string `form:"customer_id" json:"customer_id" bson:"customer_id" xml:"customer_id"`
	UserID     string `form:"user_id" json:"user_id" bson:"user_id" xml:"user_id"`
	Role       string `form:"role" json:"role" bson:"role" xml:"role" gorm:"default:OWNER"`
}

type Store struct {
	alira.BaseModel
	CustomerID string  `form:"customer_id" json:"customer_id" bson:"customer_id" xml:"customer_id"`
	Class      string  `form:"class" json:"class" bson:"class" xml:"class" gorm:"default:GENERAL"`
	Name       string  `form:"name" json:"name" bson:"name" xml:"name"`
	Address    string  `form:"address" json:"address" bson:"address" xml:"address"`
	City       string  `form:"city" json:"city" bson:"city" xml:"city"`
	State      string  `form:"state" json:"state" bson:"state" xml:"state"`
	Country    string  `form:"country" json:"country" bson:"country" xml:"country"`
	PostalCode string  `form:"postal_code" json:"postal_code" bson:"postal_code" xml:"postal_code"`
	Geocode    string  `form:"geocode" json:"geocode" bson:"geocode" xml:"geocode" gorm:"default:null"`
	Longitude  float64 `form:"longitude" json:"longitude" bson:"longitude" xml:"longitude" gorm:"default:null"`
	Latitude   float64 `form:"latitude" json:"latitude" bson:"latitude" xml:"latitude" gorm:"default:null"`
}

func (Store) TableName() string {
	return "stores"
}

type StoreProduct struct {
	alira.BaseModel
	StoreID   string `form:"store_id" json:"store_id" bson:"store_id" xml:"store_id"`
	ProductID string `form:"product_id" json:"product_id" bson:"product_id" xml:"product_id"`
	Name      string `form:"name" json:"name" bson:"name" xml:"name"`
	Image     string `form:"image" json:"image" bson:"image" xml:"image"`
	RackNo    string `form:"rack_no" json:"rack_no" bson:"rack_no" xml:"rack_no"`
	Available int64  `form:"available" json:"available" bson:"available" xml:"available" gorm:"default:0"`
	Opname    int64  `form:"opname" json:"opname" bson:"opname" xml:"opname" gorm:"default:0"`
	Return    int64  `form:"return" json:"return" bson:"return" xml:"return" gorm:"default:0"`
	Sold      int64  `form:"sold" json:"sold" bson:"sold" xml:"sold" gorm:"default:0"`
}

func (StoreProduct) TableName() string {
	return "store_products"
}

type StoreProductPrice struct {
	alira.BaseModel
	StoreProductID string    `form:"storeproduct_id" json:"storeproduct_id" bson:"storeproduct_id" xml:"storeproduct_id"`
	Quantity       int64     `form:"qty" json:"qty" bson:"qty" xml:"qty" gorm:"default:0"`
	Unit           string    `form:"unit" json:"unit" bson:"unit" xml:"unit"`
	BuyPrice       float64   `form:"buy_price" json:"buy_price" bson:"buy_price" xml:"buy_price" gorm:"default:0"`
	SellPrice      float64   `form:"sell_price" json:"sell_price" bson:"sell_price" xml:"sell_price" gorm:"default:0"`
	NotBefore      time.Time `form:"not_before" json:"not_before" bson:"not_before" xml:"not_before"`
	NotAfter       time.Time `form:"not_after" json:"not_after" bson:"not_after" xml:"not_after"`
}

func (StoreProductPrice) TableName() string {
	return "storeproduct_prices"
}
