package main

import (
	"crm-backend/config"
	"crm-backend/controllers"
	"crm-backend/middleware"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware: Frontend (Vue.js)
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	// 1. connect to database
	config.ConnectDB()

	r := gin.Default()
	r.Use(CORSMiddleware())

	// 2. Public Endpoints - Registration, Login, Logout
	authRoutes := r.Group("/api/v1/auth")
	{
		authRoutes.POST("/register", controllers.RegisterUser)
		authRoutes.POST("/login", controllers.LoginUser)
		authRoutes.POST("/logout", controllers.LogoutUser)
	}

	// 3. Protected Endpoints - Require JWT Token
	protectedRoutes := r.Group("/api/v1")
	protectedRoutes.Use(middleware.AuthRequired())
	{
		protectedRoutes.GET("/users/:id", controllers.GetUserByID)
		protectedRoutes.PUT("/users/me/profile", controllers.UpdateOwnProfile)
		protectedRoutes.PUT("/users/me/password", controllers.UpdatePassword)

		// Admin Only APIs
		adminRoutes := protectedRoutes.Group("")
		adminRoutes.Use(middleware.AdminRequired())
		{
			adminRoutes.GET("/users", controllers.GetAllUsers)
			adminRoutes.PUT("/users/:id", controllers.UpdateUser)
			adminRoutes.DELETE("/users/:id", controllers.DeleteUser)
		}
	}

	
	r.Run(":8080")
}
