package middleware

import (
	"crm-backend/controllers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthRequired: ඕනෑම ආරක්ෂිත API එකක් රන් වීමට කලින් JWT එක චෙක් කරන වැට
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "පද්ධතියට ඇතුළු වීමට අවසර නැත (Token Missing)"})
			c.Abort()
			return
		}

		// Bearer <token> කෑල්ල වෙන් කර ගැනීම
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "වැරදි Token Format එකක්"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return controllers.JWTKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "ඔබගේ Token එක කල් ඉකුත් වී ඇත හෝ අවලංගුයි"})
			c.Abort()
			return
		}

		// ඉදිරි කෝඩ් වලට පාවිච්චි කරන්න Claims ටික Context එකට දැමීම
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
