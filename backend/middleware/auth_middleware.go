package middleware

import (
    "net/http"
    "strings"
    "crm-backend/controllers"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
)

func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not allowed to access the system. (Token Missing)"})
            c.Abort()
            return
        }
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong Token Format."})
            c.Abort()
            return
        }
        tokenString := parts[1]
        claims := jwt.MapClaims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token)(interface{}, error) {
            return controllers.JWTKey, nil
        })
        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Your token is invalid or has expired."})
            c.Abort()
            return
        }
        c.Set("user_id", claims["user_id"])
        c.Set("role_id", claims["role_id"])
        c.Next()
    }
}