package notification

import (
	"net/http"
	"notification-service/internal/repositories/notification_repository"
	"notification-service/internal/services/notificationservice"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Handler struct {
	notificationRepo notification_repository.NotificationRepository
	notificationSvc  *notificationservice.NotificationService
}

func NewHandler(notificationRepo notification_repository.NotificationRepository, notificationSvc *notificationservice.NotificationService) *Handler {
	return &Handler{
		notificationRepo: notificationRepo,
		notificationSvc:  notificationSvc,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) SendNotification(c *gin.Context) {
	var input struct {
		Channel string `json:"channel" binding:"required"`
		Message string `json:"message" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	notification, err := h.notificationRepo.CreateNotification(
		c.Request.Context(),
		userID,
		input.Channel,
		input.Message,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create notification"})
		return
	}

	if err := h.notificationSvc.SendNotification(notification); err != nil {
		c.JSON(http.StatusCreated, gin.H{
			"notification": notification,
			"warning":      "Notification stored but real-time delivery failed",
		})
		return
	}

	c.JSON(http.StatusCreated, notification)
}

func (h *Handler) WebSocket(c *gin.Context) {
	userID, err := uuid.Parse(c.GetString("user_id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade connection"})
		return
	}

	h.notificationSvc.RegisterConnection(userID, conn)
	defer h.notificationSvc.RemoveConnection(userID)

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
}
