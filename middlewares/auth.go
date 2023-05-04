package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luminarapp/server/auth"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Verify token
		_, err := auth.VerifyToken(c)

		// Return error if token is invalid
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		// Continue if token is valid
		c.Next()
	}
}