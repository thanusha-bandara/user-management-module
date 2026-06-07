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

	// Force public registration to Customer role (RoleID = 4)
	input.RoleID = 4

	// 1. Password Encrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot hash password"})
		return
	}

	// 2.input data Database
	query := `INSERT INTO users (full_name, username, email, password_hash, role_id) VALUES ($1, $2, $3, $4, $5) RETURNING user_id`

	var newUserID int
	err = config.DB.QueryRow(context.Background(), query, input.FullName, input.Username, input.Email, string(hashedPassword), input.RoleID).Scan(&newUserID)

	if err != nil {

		println("Database Error:", err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to insert user into database",
			"details": err.Error(),
		})
		return
	}

	// 3. Success Response
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully! ",
		"user_id": newUserID,
	})
}

// GetAllUsers retrieves all users from the database
func GetAllUsers(c *gin.Context) {
	rows, err := config.DB.Query(context.Background(), "SELECT user_id, full_name, username, email, role_id, status, created_at FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.UserID, &u.FullName, &u.Username, &u.Email, &u.RoleID, &u.Status, &u.CreatedAt); err != nil {
			continue
		}
		users = append(users, u)
	}

	c.JSON(http.StatusOK, users)
}

// GetUserByID retrieves a single user by ID
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var u models.User
	err := config.DB.QueryRow(context.Background(), "SELECT user_id, full_name, username, email, role_id, status, created_at FROM users WHERE user_id = $1", id).
		Scan(&u.UserID, &u.FullName, &u.Username, &u.Email, &u.RoleID, &u.Status, &u.CreatedAt)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, u)
}

// UpdateUser updates user details (excluding password)
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var input models.UserUpdateDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.Exec(context.Background(),
		"UPDATE users SET full_name=$1, username=$2, email=$3, role_id=$4, status=$5 WHERE user_id=$6",
		input.FullName, input.Username, input.Email, input.RoleID, input.Status, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser removes a user from the database
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	_, err := config.DB.Exec(context.Background(), "DELETE FROM users WHERE user_id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// UpdateOwnProfile allows a user to update their own username and email
func UpdateOwnProfile(c *gin.Context) {
	// Retrieve user_id from token context
	tokenUserID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input models.UserOwnUpdateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.Exec(context.Background(),
		"UPDATE users SET full_name=$1, username=$2, email=$3 WHERE user_id=$4",
		input.FullName, input.Username, input.Email, tokenUserID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

// UpdatePassword allows a user to change their password
func UpdatePassword(c *gin.Context) {
	tokenUserID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input models.UserPasswordUpdateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch current password hash
	var currentHash string
	err := config.DB.QueryRow(context.Background(), "SELECT password_hash FROM users WHERE user_id = $1", tokenUserID).Scan(&currentHash)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Compare old password
	if err := bcrypt.CompareHashAndPassword([]byte(currentHash), []byte(input.OldPassword)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect old password"})
		return
	}

	// Hash new password
	newHash, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash new password"})
		return
	}

	// Update password in DB
	_, err = config.DB.Exec(context.Background(), "UPDATE users SET password_hash=$1 WHERE user_id=$2", string(newHash), tokenUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}
