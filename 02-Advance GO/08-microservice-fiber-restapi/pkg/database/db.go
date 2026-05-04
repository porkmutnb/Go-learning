package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=chermew password=P@ssw0rd dbname=postgres port=5432 sslmode=disable search_path=develop"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	fmt.Println("Database connected successfully! 🎉")
	DB = db
}
