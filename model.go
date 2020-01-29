package alira

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

type Model struct {
	ID        string `gorm:"primary_key"`
	CreatedBy string `gorm:"default:null"`
	CreatedAt time.Time
	UpdatedBy string `gorm:"default:null"`
	UpdatedAt time.Time
	DeletedBy string     `gorm:"default:null"`
	DeletedAt *time.Time `sql:"index"`
}

func (model *Model) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	return nil
}
