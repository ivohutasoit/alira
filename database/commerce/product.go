package commerce

import alira "github.com/ivohutasoit/alira"

type Product struct {
	alira.Model
	Barcode     string
	Name        string
	LongName    string
	Manufacture string
}

func (Product) TableName() string {
	return "products"
}
