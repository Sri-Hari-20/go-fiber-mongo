package config

import (
	"github.com/joho/godotenv"

	"log"
	"os"
)

func Config(key string) string {
    err := godotenv.Load("creds.env")

    if err != nil {
        log.Fatal("Cannot load given key: " + key)
    }

    return os.Getenv(key)
}