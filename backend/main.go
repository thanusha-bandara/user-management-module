package main

import (
    "crm-backend/config"
    "crm-backend/controllers"
    "crm-backend/middleware"
    "github.com/gin-gonic/gin"
)

func main() {
    config.ConnectDB()
    r := gin.Default()

    auth := r.Group("/auth")
    {
        auth.POST("/register", controllers.Register)
        auth.POST("/login", controllers.Login)
    }

    protected := r.Group("/api")
    protected.Use(middleware.AuthRequired())
    {
        // protected routes
    }

    r.Run(":8080")
}
