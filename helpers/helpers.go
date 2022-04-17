package helpers

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func ReadEnv(key string) string {

	//err := godotenv.Load("../.env")
	cwd, _ := os.Getwd()
	envPath := fmt.Sprintf("%s/.env", cwd)
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}

func AppPort() string {
	return ":" + ReadEnv("APP_PORT")
}
