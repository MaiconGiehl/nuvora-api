package env

import (
	"os"

	"github.com/go-chi/jwtauth"
	"github.com/joho/godotenv"
	"github.com/maicongiehl/nuvora-api/internal/core/application/shared/logger"
)

type DBConfig struct {
	DBHost        string
	DBPort        string
	DBName        string
	DBUser        string
	DBPassword    string
	WebServerPort string
	JWTSecret     string
	JWTExpiresIn  string
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig(path string, logger logger.Logger) *DBConfig {
	err := godotenv.Load(path)
	if err != nil {
		logger.Fatalf("EnvConfig.LoadConfig: Unable to load env, %s", err.Error())
	}

	cfg := &DBConfig{
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		DBName:        os.Getenv("DB_NAME"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		WebServerPort: os.Getenv("WEB_SERVER_PORT"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
		JWTExpiresIn:  os.Getenv("JWT_EXPIRESIN"),
	}

	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg
}
