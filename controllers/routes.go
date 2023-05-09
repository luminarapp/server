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
	protected.PATCH("/captures/:id", UpdateCapture)
	protected.DELETE("/captures/:id", DeleteCapture)

	// Comment routes
	protected.POST("/captures/:id/comments", CreateComment)
	protected.PATCH("/captures/:id/comments/:commentId", UpdateComment)
	protected.DELETE("/captures/:id/comments/:commentId", DeleteComment)

	// Collection routes
	protected.GET("/collections", GetUserCollections)
	protected.POST("/collections", CreateCollection)
	protected.GET("/collections/:id", GetCollection)
	protected.PATCH("/collections/:id", UpdateCollection)
	protected.DELETE("/collections/:id", DeleteCollection)

	// User routes
	protected.GET("/users/me", CurrentUser)
	protected.PATCH("/users/me", UpdateCurrentUser)
	
	// Public Auth routes
	router.POST("/users/auth", UserAuthChallenge)

	router.Run(fmt.Sprintf(":%s", config.Config().ServerPort))
}