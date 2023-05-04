package auth

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/luminarapp/server/config"
)

// Generate new JWT token
func GenerateToken(id string) string {
	claims := jwt.MapClaims{
		"id": id,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add((time.Hour * 24) * time.Duration(config.Config().SessionTokenLifetime)).Unix(),
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(config.Config().SessionSecret))

	return tokenString
}

// Extract JWT token from Authorization header or query string
func ExtractToken(c *gin.Context) string {
	// Check if token provided in query string
	token := c.Query("token")
	if token != "" {
		return token
	}

	// Check if token provided in Authorization header (Bearer)
	bearerToken := c.Request.Header.Get("Authorization")

	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	// Return empty string if no token provided
	return ""
}

// Extract User ID from JWT token
func ExtractTokenID(c *gin.Context) (string, error) {
	token, err := VerifyToken(c)

	// Return error if token is invalid
	if err != nil {
		return "", err
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		id := claims["id"].(string)
		return id, nil
	}

	// Return empty string if no token provided
	return "", nil	
}

// Verify JWT token
func VerifyToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := ExtractToken(c)

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		// Return secret key
		return []byte(config.Config().SessionSecret), nil
	})

	// Return error if token is invalid
	if err != nil {
		return nil, err
	}

	// Return token
	return token, nil
}
