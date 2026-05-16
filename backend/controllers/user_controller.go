package controllers

import (
	"context"
	"crm-backend/config"
	"crm-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	var input models.UserRegisterDTO

	// Frontend JSON to input struct binding
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 1. Password Encrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password hash කිරීමට අපොහොසත් වුණා"})
		return
	}

	// 2.input data Database
	query := `INSERT INTO users (username, email, password_hash, role_id) VALUES ($1, $2, $3, $4) RETURNING user_id`

	var newUserID int
	err = config.DB.QueryRow(context.Background(), query, input.Username, input.Email, string(hashedPassword), input.RoleID).Scan(&newUserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database එකට ඇතුළත් කරන්න බැරි වුණා (Email එක දැනටමත් ඇති)"})
		return
	}

	// 3. Success Response
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully! ",
		"user_id": newUserID,
	})
}
