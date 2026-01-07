package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv string

	DBDriver string
	DBHost   string
	DBPort   string
	DBName   string
	DBUser   string
	DBPass   string
	DBSSL    string
}

func Load() *Config {
	// Load .env hanya untuk local
	if os.Getenv("APP_ENV") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Println("⚠️ .env file not loaded (probably production)")
		}
	}

	return &Config{
		AppEnv: os.Getenv("APP_ENV"),

		DBDriver: os.Getenv("DB_DRIVER"),
		DBHost:   os.Getenv("DB_HOST"),
		DBPort:   os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
		DBUser:   os.Getenv("DB_USER"),
		DBPass:   os.Getenv("DB_PASS"),
		DBSSL:    os.Getenv("DB_SSL"),
	}
}
