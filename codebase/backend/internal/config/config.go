package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	DBDriver     string
	DatabasePath string
	DBHost       string
	DBUser       string
	DBPassword   string
	DBName       string
	JWTSecret    string
	FrontendURL  string
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		Port:         getEnv("PORT", "8080"),
		DBDriver:     getEnv("DB_DRIVER", "sqlite"),
		DatabasePath: getEnv("DATABASE_PATH", "./data/gestao.db"),
		DBHost:       getEnv("DB_HOST", ""),
		DBUser:       getEnv("DB_USER", ""),
		DBPassword:   getEnv("DB_PASSWORD", ""),
		DBName:       getEnv("DB_NAME", "gestao_psicologos"),
		JWTSecret:    getEnv("JWT_SECRET", ""),
		FrontendURL:  getEnv("FRONTEND_URL", "http://localhost:3000"),
	}
}

func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}
