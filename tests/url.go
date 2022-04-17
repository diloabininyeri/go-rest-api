package tests

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func GetUrl() string {

	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
	port := fmt.Sprintf("http://localhost:%s", os.Getenv("APP_PORT"))
	return port
}
