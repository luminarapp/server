package main

import (
	"github.com/luminarapp/server/controllers"
	"github.com/luminarapp/server/models"
	"github.com/luminarapp/server/utils"
)

func main() {
	// r := gin.Default()

	utils.SetupDirectories()
	models.ConnectDatabase()
	controllers.SetupRouter()

	// r.Run()
}
