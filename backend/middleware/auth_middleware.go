package middleware

import (
	"crm-backend/controllers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthRequired
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Token missing"})
			c.Abort()
			return
		}

		// Bearer
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token Format"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return controllers.JWTKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token"})
			c.Abort()
			return
		}

		// Set user_id and role_id in context for further use in handlers
		c.Set("user_id", claims["user_id"])
		c.Set("role_id", claims["role_id"])
		c.Next()
	}
}

// AdminRequired: Check if the user is an Admin (role_id = 1)
func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleID, exists := c.Get("role_id")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Role ID not found in token"})
			c.Abort()
			return
		}

		// claims["role_id"] might be float64 after JSON parsing in JWT
		var roleInt int
		switch v := roleID.(type) {
		case float64:
			roleInt = int(v)
		case int:
			roleInt = v
		}

		if roleInt != 1 { // 1 is Admin
			c.JSON(http.StatusForbidden, gin.H{"error": "Access Denied: Admins only"})
			c.Abort()
			return
		}
		c.Next()
	}
}
