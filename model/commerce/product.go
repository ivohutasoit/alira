package commerce

import alira "github.com/ivohutasoit/alira/model"

type Product struct {
	alira.BaseModel
	Barcode     string `form:"barcode" json:"barcode" bson:"barcode" xml:"barcode"`
	Name        string `form:"name" json:"name" bson:"name" xml:"name"`
	LongName    string `form:"long_name" json:"long_name" bson:"long_name" xml:"long_name"`
	Manufacture string `form:"manufacture" json:"manufacture" bson:"manufacture" xml:"manufacture"`
}

func (Product) TableName() string {
	return "products"
}
