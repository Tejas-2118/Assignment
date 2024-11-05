package services

import (
	"context"
	"database/sql"

	"assignment/config"
	"assignment/logger"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var RedisClient *redis.Client

func InitDatabase(cfg *config.Config) error {
	var err error
	DB, err = sql.Open("mysql", cfg.MySQLDSN)
	if err != nil {
		return err
	}
	return DB.Ping()
}


func InitCache(cfg *config.Config) error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       0,
	})

	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		logger.Error.Println(gin.H{"error": err})
		return err
	}
	logger.Info.Println("Connected to Redis")
	return nil
}
