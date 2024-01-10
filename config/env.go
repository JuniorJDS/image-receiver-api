package config

import (
	"log"

	"github.com/joho/godotenv"
)

func GetSettings() map[string]string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err.Error())
	}

	settings := make(map[string]string)

	settings["API_V1"] = "/api/v1"
	settings["PORT"] = "5000"

	return settings
}