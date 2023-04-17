package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

// Config is the struct that holds all the configuration variables
type configBuilder struct {
	DatabasePath string `env:"DB_PATH" envDefault:"./data"`
	ServerPort   string `env:"SERVER_PORT" envDefault:"8080"`
}

// BuildConfig builds the configuration variables
func Build() configBuilder {
	// Load the configuration variables from the .env file and the environment variables
	dotenvFile := ".env"
	if err := godotenv.Load(dotenvFile); err != nil {
		log.Fatalf("Error loading %s file", dotenvFile)
	}

	// Create a new configuration variable
	var config configBuilder

	// Parse the configuration variables
	if err := env.Parse(&config); err != nil {
		log.Fatalf("%+v\n", err)
	}

	// Return the configuration variables
	return config
}

// Config returns the configuration variables
func Config() configBuilder {
	return Build()
}