package handlers

import (
	"database/sql"
	"notification-service/internal/repositories/apikey_repository"
	"notification-service/internal/repositories/notification_repository"
	"notification-service/internal/repositories/user_repository"
	"notification-service/internal/services/notificationservice"
	"notification-service/server/handlers/apikey"
	"notification-service/server/handlers/auth"
	"notification-service/server/handlers/notification"
	"notification-service/server/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB, notificationSvc *notificationservice.NotificationService) {
	userRepo := user_repository.NewUserRepository(db)
	apikeyRepo := apikey_repository.NewAPIKeyRepository(db)
	notificationRepo := notification_repository.NewNotificationRepository(db)

	authHandler := auth.NewHandler(userRepo)
	apikeyHandler := apikey.NewHandler(apikeyRepo)
	notificationHandler := notification.NewHandler(notificationRepo, notificationSvc)

	api := r.Group("/api")
	{
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)

		authorized := api.Group("/", middlewares.AuthMiddleware())
		{
			authorized.POST("/api-keys", apikeyHandler.Create)
			authorized.DELETE("/api-keys/:id", apikeyHandler.Delete)
		}

		notifications := api.Group("/notifications", middlewares.APIKeyMiddleware())
		{
			notifications.POST("/send", notificationHandler.SendNotification)
			notifications.GET("/ws", notificationHandler.WebSocket)
		}
	}
}
