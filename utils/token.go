package utils

import (
	"strings"

	"github.com/google/uuid"
)

// Create new magic token
func GenerateMagicToken() string {
	// Magic token is in format XXXX-XXXX, must contain min. 1 number each side of dash
	
	// Generate random 8 character string
	uuid := uuid.New().String()

	// Split uuid at dashes
	uuid = strings.Split(uuid, "-")[0] 

	// Split uuid into two strings
	magicToken := uuid[:4] + "-" + uuid[4:8]

	return magicToken
}