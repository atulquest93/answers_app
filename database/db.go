package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connection *sql.DB

func Open() *gorm.DB {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db_username := os.Getenv("db_username")
	db_password := os.Getenv("db_password")
	db_host := os.Getenv("db_host")
	db_port := os.Getenv("db_port")
	db_database := os.Getenv("db_database")

	result := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_username, db_password, db_host, db_port, db_database)

	dsn := result
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error in connecting", err)
	}
	return connection
}

func Close() error {
	return connection.Close()
}
