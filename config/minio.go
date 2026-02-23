package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigMinio struct {
	Endpoint   string
	AccessKey  string
	SecretKey  string
	BucketName string
	Secure     bool
}

func LoadMinio() *ConfigMinio {
	// Load .env hanya untuk local
	if os.Getenv("APP_ENV") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Println("⚠️ .env file not loaded (probably production)")
		}
	}

	return &ConfigMinio{
		Endpoint:   os.Getenv("MINIO_ENDPOINT"),
		AccessKey:  os.Getenv("MINIO_ACCESS_KEY"),
		SecretKey:  os.Getenv("MINIO_SECRET_KEY"),
		BucketName: os.Getenv("MINIO_BUCKET_NAME"),
		Secure:     os.Getenv("MINIO_SECURE") == "true", // default true, set false jika tidak menggunakan HTTPS
	}
}
