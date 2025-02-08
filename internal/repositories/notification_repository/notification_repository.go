package notification_repository

import (
	"context"
	"notification-service/internal/models"
)

type NotificationRepository interface {
	CreateNotification(ctx context.Context, userID string, channel string, message string) (*models.Notification, error)
	GetNotificationsByUserID(ctx context.Context, userID string, limit int32, offset int32) ([]*models.Notification, error)
	UpdateNotificationStatus(ctx context.Context, id string, status string) error
}