package handlers

import (
	"assignment/logger"
	"assignment/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func ImportData(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		logger.Error.Println(gin.H{"error": "file is required"})
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	f, err := file.Open()
	if err != nil {
		logger.Error.Println(gin.H{"error": "could not open file"})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not open file"})
		return
	}
	defer f.Close()

	excelFile, err := excelize.OpenReader(f)
	if err != nil {
		logger.Error.Println(gin.H{"error": "Could not parse Excel file"})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not parse Excel file"})
		return
	}

	rows, err := excelFile.GetRows("uk-500")
	if err != nil {
		logger.Error.Println(gin.H{"error": "Could not read rows"})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read rows"})
		return
	}

	records := make([]models.Record, 0)
	for i, row := range rows {
		if i == 0 {
			logger.Info.Println("Skipping Header")
			continue
		}
		if len(row) < 10 {
			logger.Error.Println(gin.H{"error": "Invalid file format"})
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file format"})
			return
		}
		record := models.Record{
			FirstName:   row[0],
			LastName:    row[1],
			CompanyName: row[2],
			Address:     row[3],
			City:        row[4],
			County:      row[5],
			Postal:      row[6],
			Phone:       row[7],
			Email:       row[8],
			Web:         row[9],
		}
		records = append(records, record)
	}

	for _, record := range records {
		if err := record.Insert(); err != nil {
			logger.Error.Println(gin.H{"error": err})
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving data to database"})
			return
		}
		if err := record.Cache(); err != nil {
			logger.Error.Println(gin.H{"error": err})
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error caching data"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data imported successfully"})
}
