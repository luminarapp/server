package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v4"
	"github.com/luminarapp/server/models"
)

// GET /captures/:id/comments
func GetCaptureComments(c *gin.Context) {
	var capture models.Capture

	if err := models.DB.Where("id = ?", c.Param("id")).First(&capture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "capture id not found"})
		return
	}

	comments, err := models.GetCaptureComments(capture.Comments)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comments})
}

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

	// Create comment
	comment := models.Comment{
		ID: shortuuid.New(),
		Author: payload.Author,
		Body: payload.Body,
	}

	if err := models.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add comment to capture
	capture.Comments = append(capture.Comments, comment.ID)

	if err := models.DB.Save(&capture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// DELETE /captures/:id/comments/:commentId
func DeleteComment(c *gin.Context) {
	var capture models.Capture

	if err := models.DB.Where("id = ?", c.Param("id")).First(&capture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "capture id not found"})
		return
	}

	var comment models.Comment

	if err := models.DB.Where("id = ?", c.Param("commentId")).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "comment id not found"})
		return
	}

	if err := models.DB.Delete(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Remove comment from capture
	for i, commentId := range capture.Comments {
		if commentId == comment.ID {
			capture.Comments = append(capture.Comments[:i], capture.Comments[i+1:]...)
			break
		}
	}

	if err := models.DB.Save(&capture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comment})
}