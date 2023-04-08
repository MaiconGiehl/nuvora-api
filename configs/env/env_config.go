package env

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
)

type DBConfig struct {
	DBHost     	string 
	DBPort     	string
	DBName 			string
	DBUser     	string
	DBPassword 	string
}

func LoadConfig(logger logger.Logger) *DBConfig {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Fatalf("EnvConfig.LoadConfig: Unable to load env, %s", err.Error())
	}

	return &DBConfig{
		DBHost:              		os.Getenv("DB_HOST"),
		DBPort:     						os.Getenv("DB_PORT"),
		DBName: 								os.Getenv("DB_NAME"),
		DBUser:     						os.Getenv("DB_USER"),
		DBPassword: 						os.Getenv("DB_PASSWORD"),
	}
}