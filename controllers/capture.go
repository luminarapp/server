package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v4"
	"github.com/luminarapp/server/models"
)

// GET /captures/:id
func GetCapture(c *gin.Context) {
	var capture models.Capture

	if err := models.DB.Where("id = ?", c.Param("id")).First(&capture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "capture id not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": capture})
}

// POST /captures
func CreateCapture(c *gin.Context) {
	var payload models.CreateCaptureRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	capture := models.Capture{
		ID: shortuuid.New(),
		Author: payload.Author,
		Source: payload.Source,
	}
	
	if err := models.DB.Create(&capture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": capture})
}
