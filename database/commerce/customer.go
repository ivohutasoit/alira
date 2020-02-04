package commerce

import (
	"time"

	"github.com/google/uuid"
	alira "github.com/ivohutasoit/alira"
	"github.com/jinzhu/gorm"
)

type Customer struct {
	alira.Model
	Class   string `gorm:"default:'SHOPOWNER'"` // DISTRIBUTOR
	Code    string
	Status  string `grom:"default:'INACTIVE'"`
	Payment bool   `grom:"default:false"`
}

func (Customer) TableName() string {
	return "customers"
}

type CustomerUser struct {
	alira.Model
	CustomerID    string
	UserID        string
	Username      string
	Email         string
	PrimaryMobile string
	Role          string `gorm:"default:'OWNER'"`
	Active        bool   `gorm:"default:false"`
	Delete        bool   `gorm:"default:false"`
}

func (CustomerUser) TableName() string {
	return "customer_users"
}

type Store struct {
	alira.Model
	CustomerID string
	Class      string `gorm:"default:'GENERAL'"`
	Code       string
	Name       string
	Address    string
	City       string
	State      string
	Country    string
	PostalCode string
	Geocode    string  `gorm:"default:null"`
	Longitude  float64 `gorm:"default:null"`
	Latitude   float64 `gorm:"default:null"`
	Status     string  `gorm:"default:'OPEN'"`
	Default    bool    `gorm:"default:false"`
}

func (Store) TableName() string {
	return "stores"
}

type StoreUser struct {
	alira.Model
	CustomerUserID string
	StoreID        string
	Role           string `gorm:"default:'TELLER'"`
	NotBefore      time.Time
	NotAfter       time.Time `gorm:"default:null"`
	Status         string
	Active         bool `gorm:"default:false"`
}

func (model *StoreUser) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	scope.SetColumn("NotBefore", time.Now())
	return nil
}

func (StoreUser) TableName() string {
	return "store_users"
}

type StoreProductCategory struct {
	alira.Model
	StoreID     string
	Code        string
	Name        string
	Description string
}

func (StoreProductCategory) TableName() string {
	return "storeproduct_categories"
}

type StoreProduct struct {
	alira.Model
	StoreID    string
	ProductID  string
	CategoryID string
	Name       string
	Image      string
	RackNo     string
	Available  int64 `gorm:"default:0"`
	Opname     int64 `gorm:"default:0"`
	Return     int64 `gorm:"default:0"`
	Sold       int64 `gorm:"default:0"`
}

func (StoreProduct) TableName() string {
	return "store_products"
}

type StoreProductPrice struct {
	alira.Model
	StoreProductID string
	Quantity       int64 `gorm:"default:0"`
	Unit           string
	Currency       string
	BuyPrice       float64 `gorm:"default:0"`
	SellPrice      float64 `gorm:"default:0"`
	NotBefore      time.Time
	NotAfter       time.Time
}

func (model *StoreProductPrice) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	scope.SetColumn("NotBefore", time.Now())
	scope.SetColumn("AgreedAt", time.Now())
	return nil
}

func (StoreProductPrice) TableName() string {
	return "storeproduct_prices"
}
