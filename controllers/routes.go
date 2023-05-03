package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/luminarapp/server/config"
)

func SetupRouter() {
	router := gin.Default()
	
	// TODO: Seperate private and public routes with middleware authentication
	// TODO: User routes (login, register, me, etc.)

	// Capture routes
	router.GET("/captures/:id", GetCapture)
	router.POST("/captures", CreateCapture)
	router.DELETE("/captures/:id", DeleteCapture)

	// Comment routes
	router.GET("/captures/:id/comments", GetCaptureComments)
	router.POST("/captures/:id/comments", CreateComment)
	router.DELETE("/captures/:id/comments/:commentId", DeleteComment)

	// Collection routes
	router.GET("/collections/:id", GetCollection)
	router.POST("/collections", CreateCollection)
	router.DELETE("/collections/:id", DeleteCollection)
	router.GET("/collections/:id/captures", GetCollectionCaptures)
	router.POST("/collections/:id/captures", AddCaptureToCollection)

	router.Run(fmt.Sprintf(":%s", config.Config().ServerPort))
}