package configuration

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	DBHost     string
	DBPassword string
	DBUser     string
	DBName     string
	DBPort     string
}

func InitializeConfiguration() *Configuration {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Fatal("Could not load environment variables")

	}

	return &Configuration{
		DBHost:     os.Getenv("DB_HOST"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBUser:     os.Getenv("DB_USER"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
	}
}
