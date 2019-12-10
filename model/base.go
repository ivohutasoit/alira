package model

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB

type BaseModel struct {
	ID        string     `gorm:"primary_key" json:"id" bson:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading .env file")
	}

	engine := os.Getenv("DATABASE.ENGINE")
	var conn *gorm.DB
	if engine == "sqlite3" {
		file := os.Getenv("DATABASE.FILE")
		conn, err = gorm.Open(engine, file)
	} else if engine == "postgres" {
		host := os.Getenv("DATABASE.HOST")
		port := os.Getenv("DATABASE.PORT")
		schema := os.Getenv("DATABASE.SCHEMA")
		username := os.Getenv("DATABASE.USERNAME")
		password := os.Getenv("DATABASE.PASSWORD")

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
