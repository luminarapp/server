package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v4"
	"github.com/luminarapp/server/auth"
	"github.com/luminarapp/server/models"
)

// GET /collections
func GetUserCollections(c *gin.Context) {
	user, err := auth.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get collections
	collections, err := models.GetCollectionsByUserId(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return collections
	c.JSON(http.StatusOK, gin.H{"data": collections})
}

// GET /collections/:id
func GetCollection(c *gin.Context) {
	var collection models.Collection

	// Get collection with preloaded captures
	if err := models.DB.Preload("Captures").Where("id = ?", c.Param("id")).First(&collection).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "collection id not found"})
		return
	}

	// Check if user is authorized to view collection
	// TODO: Add collection visibility / sharing settings
	userId, err := auth.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if collection.UserID != userId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "collection is private"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": collection})
}

// POST /collections
func CreateCollection(c *gin.Context) {
	var payload models.CreateCollectionRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user
	userId, err := auth.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create collection
	collection := models.Collection{
		ID: shortuuid.New(),
		UserID: userId,
		Name: payload.Name,
		Description: payload.Description,
	}

	if err := models.DB.Create(&collection).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": collection})
}

// DELETE /collections/:id
func DeleteCollection(c *gin.Context) {
	var collection models.Collection

	if err := models.DB.Where("id = ?", c.Param("id")).First(&collection).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "collection id not found"})
		return
	}

	// Authenticate user
	userId, err := auth.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if collection.UserID != userId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user is missing permissions to delete collection"})
		return
	}

	// Delete collection
	if err := models.DB.Delete(&collection).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": collection})
}