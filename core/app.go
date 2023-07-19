package core

import (
	"log"
	"time"

	"github.com/joho/godotenv"
)

func Initialize() {
	loadEnvFile()
	connectToDatabase()
	setTimeZone()
}

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading dot env file")
	}
}

func setTimeZone() {
	time.LoadLocation("Asia/Malaysia")
}
