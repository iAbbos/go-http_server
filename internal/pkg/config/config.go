package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	APP         string
	Environment string
	Server      struct {
		Host string
		Port string
	}
	FilesDir string
}

func NewConfig() (*Config, error) {
	newGetEnv(".env")
	var config Config

	config.APP = getEnv("APP", "app")
	config.Environment = getEnv("ENVIRONMENT", "develop")

	config.Server.Host = getEnv("SERVER_HOST", "localhost")
	config.Server.Port = getEnv("SERVER_PORT", ":4221")

	fileDir := getEnv("FILE_DIR", "")

	config.FilesDir = fileDir

	return &config, nil
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}

func newGetEnv(envFile string) {
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
