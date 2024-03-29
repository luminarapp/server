package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v4"
	"github.com/luminarapp/server/auth"
	"github.com/luminarapp/server/models"
)

// GET /captures
func GetUserCaptures(c *gin.Context) {
	user, err := auth.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get captures
	captures, err := models.GetCapturesByUserId(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Paginate captures
	c.JSON(http.StatusOK, gin.H{"data": captures})
}

// GET /captures/:id
func GetCapture(c *gin.Context) {
	var capture models.Capture

	// Get capture with preloaded comments
	if err := models.DB.Preload("Comments").Where("id = ?", c.Param("id")).First(&capture).Error; err != nil {
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

	// Get user
	userId, err := auth.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if collection exists and belongs to user (if collection id is provided)
	if payload.CollectionID != "" {
		var collection models.Collection

		if err := models.DB.Where("id = ?", payload.CollectionID).First(&collection).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "collection id not found"})
			return
		}

		if collection.UserID != userId {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user not allowed to add capture to collection"})
			return
		}
	}

	// Create capture
	capture := models.Capture{
		ID: shortuuid.New(),
		UserID: userId,
		CollectionID: payload.CollectionID,
		Reference: payload.Reference,
	}
	
	if err := models.DB.Create(&capture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": capture})
}

// DELETE /captures/:id
func DeleteCapture(c *gin.Context) {
	var capture models.Capture

	if err := models.DB.Where("id = ?", c.Param("id")).First(&capture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "capture id not found"})
		return
	}

	// Authenticate user
	userId, err := auth.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if capture.UserID != userId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not allowed to delete capture"})
		return
	}

	// Delete capture
	if err := models.DB.Delete(&capture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": capture})
}

// PATCH /captures/:id
func UpdateCapture(c *gin.Context) {
    var payload models.UpdateCaptureRequest
    var capture models.Capture

    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Check if capture exists
    if err := models.DB.Where("id = ?", c.Param("id")).First(&capture).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "capture id not found"})
        return
    }

    // Authenticate user
    userId, err := auth.ExtractTokenID(c)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if capture.UserID != userId {
        c.JSON(http.StatusBadRequest, gin.H{"error": "user not allowed to update capture"})
        return
    }

	// Validate collection id
	if payload.CollectionID != "" {
		var collection models.Collection

		if err := models.DB.Where("id = ?", payload.CollectionID).First(&collection).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "collection id not found"})
			return
		}

		if collection.UserID != userId {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user not allowed to add capture to collection"})
			return
		}
	}

    // Update capture
    updates := map[string]interface{}{
        "collection_id": payload.CollectionID,
    }

    if err := models.DB.Model(&capture).Updates(updates).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": capture})
}
