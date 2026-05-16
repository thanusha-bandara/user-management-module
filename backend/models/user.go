package models

import "time"

// User struct represents the users table in the database
type User struct {
	UserID       int       `json:"user_id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	RoleID       int       `json:"role_id"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}

// UserRegisterDTO: This struct is used for binding the JSON input from the frontend during user registration
type UserRegisterDTO struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	RoleID   int    `json:"role_id" binding:"required"`
}
