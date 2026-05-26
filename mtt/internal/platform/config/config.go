package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT   string
	DB_URL string
}

func read(key, def string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	if def == "" {
		log.Fatalf("Error: %s key not found", key)
	}
	return def
}

func LoadConfig() *Config {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error: failed to load .env")
		}
	} else {
		log.Println("No .env file found. Defaulting on system variable")
	}

	return &Config{
		PORT:   read("PORT", "8080"),
		DB_URL: read("DATABASE_URL", ""),
	}
}
