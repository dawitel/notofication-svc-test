package notificationservice

import (
	"encoding/json"
	"sync"

	"notification-service/internal/models"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type NotificationService struct {
	connections map[uuid.UUID]*websocket.Conn
	mu          sync.RWMutex
}

func NewNotificationService() *NotificationService {
	return &NotificationService{
		connections: make(map[uuid.UUID]*websocket.Conn),
	}
}

func (s *NotificationService) RegisterConnection(userID uuid.UUID, conn *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.connections[userID] = conn
}

func (s *NotificationService) RemoveConnection(userID uuid.UUID) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.connections, userID)
}

func (s *NotificationService) SendNotification(notification *models.Notification) error {
	s.mu.RLock()
	conn, exists := s.connections[notification.UserID]
	s.mu.RUnlock()

	if !exists {
		return nil // User not connected
	}

	message, err := json.Marshal(notification)
	if err != nil {
		return err
	}

	return conn.WriteMessage(websocket.TextMessage, message)
}
