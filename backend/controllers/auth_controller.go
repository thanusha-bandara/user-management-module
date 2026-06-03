package controllers

import (
	"context"
	"crm-backend/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var JWTKey = []byte("CRM_SUPER_SECRET_KEY_2026")

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func LoginUser(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "enter valid email and password"})
		return
	}

	var userID int
	var username string
	var passwordHash string
	var roleID int
	var status string

	// search user
	query := `SELECT user_id, username, password_hash, role_id, status FROM users WHERE email = $1`
	err := config.DB.QueryRow(context.Background(), query, input.Email).Scan(&userID, &username, &passwordHash, &roleID, &status)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no user found with the provided email"})
		return
	}

	if status != "active" {
		c.JSON(http.StatusForbidden, gin.H{"error": "your account is temporarily disabled"})
		return
	}

	// compare the provided password
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong password"})
		return
	}

	// JWT 
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"role_id":  roleID,
		"exp":      expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create JWT token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "login successful! ",
		"token":    tokenString,
		"user_id":  userID,
		"username": username,
		"role_id":  roleID,
	})
}

// Logout
func LogoutUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "logout successful!"})
}
