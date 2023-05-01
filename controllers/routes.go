package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/luminarapp/server/config"
)

func SetupRouter() {
	router := gin.Default()
	
	// TODO: Seperate private and public routes with middleware authentication
	// Capture routes
	router.GET("/capture/:id", GetCapture)
	router.POST("/capture", CreateCapture)
	router.GET("/capture/:id/comments", GetCaptureComments)

	// Collection routes
	router.GET("/collection/:id", GetCollection)
	router.POST("/collection", CreateCollection)
	router.GET("/collection/:id/captures", GetCollectionCaptures)
	router.POST("/collection/:id/captures", AddCaptureToCollection)

	// Comment routes
	router.GET("/comment/:id", GetComment)
	router.POST("/comment", CreateComment)

	router.Run(fmt.Sprintf(":%s", config.Config().ServerPort))
}