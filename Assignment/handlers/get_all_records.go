package handlers

import (
	"assignment/logger"
	models "assignment/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllRecords(c *gin.Context) {
	records, err := models.GetAllFromCache()
	if err != nil || len(records) == 0 {
		logger.Info.Println(gin.H{"Length of records": len(records)})
		logger.Error.Println(gin.H{"error": err})
		records, err = models.GetAllFromDB()
		if err != nil {
			logger.Error.Println(gin.H{"error": err})
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
			return
		}
		for _, record := range records {
			err = record.Cache()
			if err != nil {
				logger.Error.Println(gin.H{"error": err})
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching data"})
			}
		}
	}
	c.JSON(http.StatusOK, records)
}
