package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v4"
	"github.com/luminarapp/server/models"
)

// POST /comment
func CreateComment(c *gin.Context) {
	var payload models.CreateCommentRequest

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
	var capture models.Capture

	if err := models.DB.Where("id = ?", payload.CaptureID).First(&capture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "capture id not found"})
		return
	}

	capture.Comments = append(capture.Comments, comment.ID)

	if err := models.DB.Save(&capture).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// GET /comment/:id
func GetComment(c *gin.Context) {
	var comment models.Comment

	if err := models.DB.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "comment id not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comment})
}