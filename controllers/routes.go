package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/luminarapp/server/config"
	"github.com/luminarapp/server/middlewares"
)

func SetupRouter() {
	router := gin.Default()

	protected := router.Group(config.Config().ApiRootPath)
	protected.Use(middlewares.JwtAuthMiddleware())

	// Capture routes
	protected.GET("/captures", GetUserCaptures)
	protected.POST("/captures", CreateCapture)
	protected.GET("/captures/:id", GetCapture)
	protected.DELETE("/captures/:id", DeleteCapture)

	// Comment routes
	protected.POST("/captures/:id/comments", CreateComment)
	protected.DELETE("/captures/:id/comments/:commentId", DeleteComment)

	// Collection routes
	protected.GET("/collections", GetUserCollections)
	protected.GET("/collections/:id", GetCollection)
	protected.POST("/collections", CreateCollection)
	protected.DELETE("/collections/:id", DeleteCollection)

	// User routes
	protected.GET("/users/me", CurrentUser)
	router.POST("/users/auth", UserAuthChallenge)

	router.Run(fmt.Sprintf(":%s", config.Config().ServerPort))
}