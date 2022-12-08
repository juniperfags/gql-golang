package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	ServerPort     string
	ServerEndpoint string
}

func GetEnvironmentVariables() Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return Env{
		ServerPort:     os.Getenv("PORT"),
		ServerEndpoint: os.Getenv("ENDPOINT"),
	}

}
