package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"mototrack/controller"
	"mototrack/initializers"
	"mototrack/middleware"

	"github.com/gin-contrib/cors"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()

	router.Use(gin.Logger())
	// router.Use(cors.Default())
	allowedOrigins := []string{"http://localhost:3000"}
	if prodOrigin := os.Getenv("ALLOWED_ORIGIN"); prodOrigin != "" {
		allowedOrigins = append(allowedOrigins, prodOrigin)
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Define a simple GET endpoint
	router.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Auth routes
	router.POST("/auth/login", controller.Login)
	router.POST("/auth/register", controller.CreateUser)

	protected := router.Group("/api")
	protected.Use(middleware.AuthenticateMiddleware)

	protected.GET("/users", controller.GetUsers)
	protected.GET("/garage", controller.GetGarage)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port) // explicitly bind to all interfaces
}
