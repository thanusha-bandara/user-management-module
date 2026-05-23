package models

import "time"

// User struct represents the users table in the database
type User struct {
	UserID       int       `json:"user_id"`
	FullName     string    `json:"full_name"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	RoleID       int       `json:"role_id"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}

// UserRegisterDTO: This struct is used for binding the JSON input from the frontend during user registration
type UserRegisterDTO struct {
	FullName string `json:"full_name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	RoleID   int    `json:"role_id"`
}

// UserUpdateDTO for updating user
type UserUpdateDTO struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleID   int    `json:"role_id"`
	Status   string `json:"status"`
}

// UserOwnUpdateDTO for users updating their own profile
type UserOwnUpdateDTO struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// UserPasswordUpdateDTO for changing password
type UserPasswordUpdateDTO struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}
