package config

import (
	"assignment/logger"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MySQLDSN      string
	RedisAddr     string
	RedisPassword string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		logger.Error.Printf("Error loading .env file: %v", err)
	}
	logger.Info.Println("Configuration Loaded successfully")
	return &Config{
		MySQLDSN: fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME")),
		RedisAddr:     os.Getenv("REDIS_ADDR"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
	}

}
