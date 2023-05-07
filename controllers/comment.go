package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v4"
	"github.com/luminarapp/server/auth"
	"github.com/luminarapp/server/models"
)

// POST /captures/:id/comments
func CreateComment(c *gin.Context) {
	var payload models.CreateCommentRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if capture exists
	var capture models.Capture

	if err := models.DB.Where("id = ?", c.Param("id")).First(&capture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "capture id not found"})
		return
	}

	// Get user
	userId, err := auth.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create comment
	comment := models.Comment{
		ID: shortuuid.New(),
		UserID: userId,
		CaptureID: capture.ID,
		Body: payload.Body,
	}

	if err := models.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return comment
	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// DELETE /captures/:id/comments/:commentId
func DeleteComment(c *gin.Context) {
	// Get capture
	var capture models.Capture

	if err := models.DB.Where("id = ?", c.Param("id")).First(&capture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "capture id not found"})
		return
	}

	// Get comment
	var comment models.Comment

	if err := models.DB.Where("id = ?", c.Param("commentId")).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "comment id not found"})
		return
	}

	// Authenticate user
	userId, err := auth.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if comment.UserID != userId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user is not authorized to delete this comment"})
		return
	}

	// Delete comment
	if err := models.DB.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Save(&capture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comment})
}