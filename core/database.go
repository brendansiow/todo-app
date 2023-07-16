package core

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Migration bool

func connectToDatabase() {
	var err error
	var dsn string

	if Migration {
		dsn = os.Getenv("LOCALHOST_DB_URL")
	} else {
		dsn = os.Getenv("DB_URL")
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}
}
