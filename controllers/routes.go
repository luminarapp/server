package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/luminarapp/server/config"
)

func SetupRouter() {
	router := gin.Default()
	
	// TODO: Seperate private and public routes with middleware authentication
	// router.GET("/books", FindBooks)

	router.Run(fmt.Sprintf(":%s", config.Config().ServerPort))
}