package main

import (
	"assignment/config"
	"assignment/handlers"
	"assignment/logger"
	"assignment/services"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func createTable(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS Employee (
	id INT AUTO_INCREMENT PRIMARY KEY,
	first_name VARCHAR(100),
	last_name VARCHAR(100),
	company_name VARCHAR(100),
	address VARCHAR(225),
	city VARCHAR(100),
	county VARCHAR(100),
	postal VARCHAR(20),
	phone VARCHAR(20),
	email VARCHAR(100),
	web VARCHAR(100)
    );`

	_, err := db.Exec(query)
	return err
}

func main() {
	cfg := config.LoadConfig()
	if err := services.InitDatabase(cfg); err != nil {
		logger.Error.Println(gin.H{"error": err})
		panic("Failed to connect to database")
	}
	defer services.DB.Close()

	if err := createTable(services.DB); err != nil {
		logger.Error.Println(gin.H{"error creating table": err})
	}
	logger.Info.Println("Table created successfully!")

	err := services.InitCache(cfg)
	if err != nil {
		logger.Error.Println(gin.H{"Error initializing Redis": err})
	}

	r := gin.Default()
	r.POST("/records", handlers.ImportData)
	r.GET("/records", handlers.GetAllRecords)
	r.PUT("/records", handlers.UpdateRecord)
	r.DELETE("/records/:id", handlers.DeleteRecord)

	err = r.Run(":8080")
	if err != nil {
		logger.Error.Println(gin.H{"error": err})
		panic(err)
	}
}
