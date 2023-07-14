package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetConfig() *Config {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatalf("Error when load environment #{err.Error}")
	}

	return &Config{
		Server{
			Host: os.Getenv("HOST"),
			Port: os.Getenv("PORT"),
		},
		Redis{
			Password: os.Getenv("REDIS_PASSWORD"),
		},
		MySql{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		Midtrans{
			Key:          os.Getenv("MIDTRANS_KEY"),
			IsProduction: os.Getenv("ENVIRONMENT") == "production",
		},
		JWT{
			SecretKey: os.Getenv("JWT_SECRET"),
		},
	}
}
