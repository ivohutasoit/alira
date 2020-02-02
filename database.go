package alira

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var connection *gorm.DB

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

	connection = conn
}

func GetConnection() *gorm.DB {
	return connection
}
