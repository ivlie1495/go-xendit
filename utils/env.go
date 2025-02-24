package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() map[string]string {
	env, err := godotenv.Read()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return env
}

var Env = LoadEnv()
