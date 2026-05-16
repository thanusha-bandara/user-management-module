package main

import (
	"crm-backend/config"
	"crm-backend/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1.connect to database
	config.ConnectDB()

	// 2. Gin Router
	r := gin.Default()

	// 3. API Endpoints
	r.POST("/api/v1/auth/register", controllers.RegisterUser)

	r.Run(":8080")
}
