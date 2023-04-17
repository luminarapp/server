package models

import (
	"github.com/luminarapp/server/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	databasePath := config.Config().DatabasePath + "/database.db"
	database, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	
	if err != nil {
		panic("Failed to connect to database!")
	}
	
	err = database.AutoMigrate(&Book{})
	if err != nil {
		panic("Failed to migrate database!")
	}

	DB = database
}