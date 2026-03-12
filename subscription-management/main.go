package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"subscription-management/controller"
	"subscription-management/initializers"
	"subscription-management/middleware"

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

	// Subscription routes
	protected.GET("/subscriptions", controller.GetSubscriptions)
	protected.GET("/subscriptions/:id", controller.GetSubscription)
	protected.GET("/subscriptions/total", controller.GetSubscriptionTotal)
	protected.POST("/subscriptions", controller.CreateSubscription)
	protected.PUT("/subscriptions/:id", controller.UpdateSubscription)
	protected.DELETE("/subscriptions/:id", controller.DeleteSubscription)

	// Installment Payment routes
	protected.GET("/installment-payments", controller.GetInstallmentPayments)
	protected.GET("/installment-payments/:id", controller.GetInstallmentPayment)
	protected.POST("/installment-payments", controller.CreateInstallmentPayment)
	protected.PUT("/installment-payments/:id", controller.UpdateInstallmentPayment)
	protected.DELETE("/installment-payments/:id", controller.DeleteInstallmentPayment)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port) // explicitly bind to all interfaces
}
