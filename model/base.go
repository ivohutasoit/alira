package model

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
)

var db *gorm.DB

type BaseModel struct {
	ID        string     `gorm:"primary_key" json:"id" bson:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

func (base *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.New().String())
	return nil
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading .env file")
	}

	engine := os.Getenv("DATABASE_ENGINE")
	var conn *gorm.DB
	if engine == "sqlite3" {
		file := os.Getenv("DATABASE_FILE")
		conn, err = gorm.Open(engine, file)
	} else if engine == "postgres" {
		host := os.Getenv("DATABASE_HOST")
		port := os.Getenv("DATABASE_PORT")
		schema := os.Getenv("DATABASE_SCHEMA")
		username := os.Getenv("DATABASE_USERNAME")
		password := os.Getenv("DATABASE_PASSWORD")

		uri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			host, port, username, schema, password)

		fmt.Println(uri)
		conn, err = gorm.Open(engine, uri)
	}

	if err != nil {
		fmt.Print(err)
	}

	db = conn
}

func GetDatabase() *gorm.DB {
	return db
}
