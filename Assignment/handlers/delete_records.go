package handlers

import (
	"assignment/logger"
	"assignment/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteRecord(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error.Println(gin.H{"error": err})
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid record ID"})
		return
	}

	record := models.Record{ID: id}
	if err := record.Delete(); err != nil {
		logger.Error.Println(gin.H{"error": err})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting data"})
		return
	}

	if err := record.RemoveFromCache(); err != nil {
		logger.Error.Println(gin.H{"error": err})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error removing data from cache"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Record deleted successfully"})
}
