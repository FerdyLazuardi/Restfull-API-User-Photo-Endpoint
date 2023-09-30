package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectoDB() {
	var err error
	dsn := "host=rosie.db.elephantsql.com user=weesadwg password=m2xYYpwHEzAk2M-ZNfTNIWuPMRWTNvRJ dbname=weesadwg port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database")
	}
}
