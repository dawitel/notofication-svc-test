package main

import (
	"log"
	"os"

	adapterpostgres "notification-service/internal/infrastructure/adapter-postgres"
	"notification-service/internal/services/notificationservice"
	"notification-service/server/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Warning: .env file not found")
	}

	r := gin.Default()

	d, err := adapterpostgres.New()
	if err != nil {
		log.Fatal("failed to connect to the postgrs dataase")
	}

	notificationSvc := notificationservice.NewNotificationService()

	handlers.RegisterRoutes(r, d.Conn, notificationSvc)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
