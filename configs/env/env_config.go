package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DBHost     	string 
	DBPort     	string
	DBName 			string
	DBUser     	string
	DBPassword 	string
}

func LoadConfig() *DBConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Connection can't be established.\nErr: %s", err)
	}
	return &DBConfig{
		DBHost:              		os.Getenv("DB_HOST"),
		DBPort:     						os.Getenv("DB_PORT"),
		DBName: 								os.Getenv("DB_NAME"),
		DBUser:     						os.Getenv("DB_USER"),
		DBPassword: 						os.Getenv("DB_PASSWORD"),
	}
}