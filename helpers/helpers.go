package helpers

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func ReadEnv(key string) string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}

func AppPort() string {
	return ":" + ReadEnv("APP_PORT")
}
