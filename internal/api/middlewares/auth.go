package middlewares

import (
	"net/http"
	"sum/internal/api"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secretJWTKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractToken(c)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
		err := api.ValidToken(token, secretJWTKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func extractToken(c *gin.Context) string {
	const BearerSchema = "Bearer "
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}
	tokenString := authHeader[len(BearerSchema):]
	return tokenString
}
