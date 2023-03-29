package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DBEnv 			string
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
		DBEnv:									os.Getenv("ENV"),
		DBHost:              		os.Getenv("DB_HOST"),
		DBPort:     						os.Getenv("DB_PORT"),
		DBName: 								os.Getenv("DB_NAME"),
		DBUser:     						os.Getenv("DB_USER"),
		DBPassword: 						os.Getenv("DB_PASSWORD"),
	}
}