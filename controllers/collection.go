package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v4"
	"github.com/luminarapp/server/models"
)

// GET /collection/:id
func GetCollection(c *gin.Context) {
	var collection models.Collection

	if err := models.DB.Where("id = ?", c.Param("id")).First(&collection).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "collection id not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": collection})
}

// POST /collection
func CreateCollection(c *gin.Context) {
	var payload models.CreateCollectionRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := models.Collection{
		ID: shortuuid.New(),
		Name: payload.Name,
		Description: payload.Description,
	}

	if err := models.DB.Create(&collection).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": collection})
}

// GET /collection/:id/captures
func GetCollectionCaptures(c *gin.Context) {
	var collection models.Collection

	if err := models.DB.Where("id = ?", c.Param("id")).First(&collection).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "collection id not found"})
		return
	}

	captures, err := models.GetCollectionCaptures(collection.Captures)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": captures})
}

// POST /collection/:id/captures
func AddCaptureToCollection(c *gin.Context) {
	var payload models.AddCaptureToCollectionRequest
	var collection models.Collection
	var capture models.Capture

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Where("id = ?", c.Param("id")).First(&collection).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "collection id not found"})
		return
	}

	if err := models.DB.Where("id = ?", payload.CaptureID).First(&capture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "capture id not found"})
		return
	}

	// Check if capture is already in collection
	for _, id := range collection.Captures {
		if id == payload.CaptureID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "capture already in collection"})
			return
		}
	}

	collection.Captures = append(collection.Captures, capture.ID)

	if err := models.DB.Save(&collection).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": collection})
}