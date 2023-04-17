package utils

import (
	"errors"
	"log"
	"os"

	"github.com/luminarapp/server/config"
)

// CreateDirIfNotExists creates a directory if it doesn't exist
func createDirIfNotExists(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

// Run necessary setup functions
func SetupDirectories() {
	// Make sure database path exists
	createDirIfNotExists(config.Config().DatabasePath)
}