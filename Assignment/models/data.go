package models

import (
	"assignment/logger"
	"assignment/services"
	"context"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Record struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	City        string `json:"city"`
	County      string `json:"county"`
	Postal      string `json:"postal"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Web         string `json:"web"`
}

func (r *Record) Insert() error {
	query := `
        INSERT INTO Employee (first_name, last_name, company_name, address, city, county, postal, phone, email, web)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `
	_, err := services.DB.Exec(query,
		r.FirstName, r.LastName, r.CompanyName, r.Address, r.City, r.County,
		r.Postal, r.Phone, r.Email, r.Web,
	)
	if err != nil {
		logger.Error.Println(gin.H{"error": err})
		return fmt.Errorf("failed to insert record into database: %w", err)
	}


	return nil
}
func (r *Record) Cache() error {
	ctx := context.Background()
	data, err := json.Marshal(r)
	if err != nil {
		logger.Error.Println(gin.H{"error": err})
		return fmt.Errorf("failed to marshal record to JSON: %w", err)
	}
	key := fmt.Sprintf("record:%d", r.ID)
	err = services.RedisClient.Set(ctx, key, data, 0).Err()
	if err != nil {
		logger.Error.Println(gin.H{"error": err})
		return fmt.Errorf("failed to store record in cache: %w", err)
	}

	return nil
}
func (r *Record) Update() error {
	query := `
        UPDATE Employee
        SET first_name = ?, last_name = ?, company_name = ?, address = ?, city = ?, county = ?, postal = ?, phone = ?, email = ?, web = ?
        WHERE id = ?
    `
	_, err := services.DB.Exec(query,
		r.FirstName, r.LastName, r.CompanyName, r.Address, r.City,
		r.County, r.Postal, r.Phone, r.Email, r.Web, r.ID,
	)
	if err != nil {
		logger.Error.Println(gin.H{"error": err})
		return fmt.Errorf("failed to update record in database: %w", err)
	}

	return nil
}

func GetAllFromCache() ([]Record, error) {
	ctx := context.Background()
	cachedData, err := services.RedisClient.Get(ctx, "records").Result()
	if err != nil {
		logger.Error.Println(gin.H{"error": err})

		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("error fetching from cache: %v", err)
	}

	var records []Record
	if err := json.Unmarshal([]byte(cachedData), &records); err != nil {
		logger.Error.Println(gin.H{"error": err})
		return nil, fmt.Errorf("error unmarshalling cached data: %v", err)
	}

	return records, nil
}

func GetAllFromDB() ([]Record, error) {
	records := []Record{}
	query := "SELECT id,first_name, last_name, company_name, address, city, county, postal, phone, email, web FROM Employee"

	rows, err := services.DB.Query(query)
	if err != nil {
		logger.Error.Println(gin.H{"error": err})
		return nil, fmt.Errorf("error querying database: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var record Record

		if err := rows.Scan(&record.ID, &record.FirstName, &record.LastName, &record.CompanyName, &record.Address, &record.City, &record.County, &record.Postal, &record.Phone, &record.Email, &record.Web); err != nil {
			logger.Error.Println(gin.H{"error": err})
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		records = append(records, record)
	}

	return records, nil
}

func (r *Record) Delete() error {
	query := "DELETE FROM Employee WHERE id = ?"
	_, err := services.DB.Exec(query, r.ID)
	if err != nil {
		logger.Error.Println(gin.H{"error": err})
		return fmt.Errorf("failed to delete record from database: %w", err)
	}
	return nil
}

func (r *Record) RemoveFromCache() error {
	key := fmt.Sprintf("record:%d", r.ID)
	ctx := context.Background()
	_, err := services.RedisClient.Del(ctx, key).Result()
	if err != nil {
		logger.Error.Println(gin.H{"error": err})
		return fmt.Errorf("failed to remove record from cache: %w", err)
	}
	return nil
}
