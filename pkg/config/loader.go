package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  .env file not found, using environment variables")
	}

	dbPort, err := strconv.Atoi(getEnv("PORT", "5432"))
	if err != nil {
		log.Fatal("invalid db port")
	}

	appPort, err := strconv.Atoi(getEnv("PORT", "8080"))
	if err != nil {
		log.Fatal("invalid app port")
	}

	return &Config{
		AppConfig: AppConfig{
			AppEnv: getEnv("APP_ENV", "local"),
			Port:   appPort,
		},
		DBConfig: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     dbPort,
			Username: getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			Database: getEnv("DB_NAME", "postgres"),
		},
		JwtConfig: JwtConfig{
			Secret: getEnv("JWT_SECRET", "secret"),
			Issuer: getEnv("JWT_ISSUER", "issuer"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
