package handlers

import (
	"assignment/logger"
	"assignment/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateRecord(c *gin.Context) {
	var record models.Record
	if err := c.ShouldBindJSON(&record); err != nil {
		logger.Error.Println(gin.H{"error": err})
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := record.Update(); err != nil {
		logger.Error.Println(gin.H{"error": err})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating data"})
		return
	}
	if err := record.Cache(); err != nil {
		logger.Error.Println(gin.H{"error": err})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating cache"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Record updated successfully"})
}
