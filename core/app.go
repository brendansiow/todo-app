package core

import (
	"log"

	"github.com/joho/godotenv"
)

func Initialize() {
	loadEnvFile()
	connectToDatabase()
}

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading dot env file")
	}
}
