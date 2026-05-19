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

// JWT සඳහා රහස්‍ය කී එක (Production වලදී මෙවැනි ඒවා .env එකට දමයි)
var JWTKey = []byte("CRM_SUPER_SECRET_KEY_2026")

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginUser: පරිශීලකයා පද්ධතියට ඇතුළත් කරගෙන JWT Token එකක් ලබා දීම
func LoginUser(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "කරුණාකර නිවැරදි Email සහ Password ඇතුළත් කරන්න"})
		return
	}

	var userID int
	var username string
	var passwordHash string
	var roleID int
	var status string

	// Database එකෙන් පරිශීලකයා සෙවීම
	query := `SELECT user_id, username, password_hash, role_id, status FROM users WHERE email = $1`
	err := config.DB.QueryRow(context.Background(), query, input.Email).Scan(&userID, &username, &passwordHash, &roleID, &status)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "ඇතුළත් කළ Email ලිපිනය පද්ධතියේ නොමැත"})
		return
	}

	if status != "active" {
		c.JSON(http.StatusForbidden, gin.H{"error": "ඔබගේ ගිණුම තාවකාලිකව අක්‍රීය කර ඇත"})
		return
	}

	// Password එක සසඳා බැලීම
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "ඇතුළත් කළ මුරපදය (Password) වැරදියි"})
		return
	}

	// JWT Claims සැකසීම (Valid for 24 Hours)
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "JWT Token එක සෑදීමට අපොහොසත් වුණා"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "ලොගින් වීම සාර්ථකයි! 🎉",
		"token":    tokenString,
		"user_id":  userID,
		"username": username,
		"role_id":  roleID,
	})
}

// LogoutUser: පරිශීලකයා පද්ධතියෙන් ඉවත් වීම (Frontend එකෙන් Token එක අයින් කිරීමට අමතරව ආරක්ෂාවට)
func LogoutUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "සාර්ථකව පද්ධතියෙන් ඉවත් වුණා! 👋"})
}
