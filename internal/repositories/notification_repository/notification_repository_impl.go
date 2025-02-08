package notification_repository

import (
	"context"
	"database/sql"
	"notification-service/internal/db"
	"notification-service/internal/models"

	"github.com/google/uuid"
)

type notificationRepositoryImpl struct {
	queries *db.Queries
}

func NewNotificationRepository(conn *sql.DB) NotificationRepository {
	return &notificationRepositoryImpl{
		queries: db.New(conn),
	}
}

func (r *notificationRepositoryImpl) CreateNotification(ctx context.Context, userID string, channel string, message string) (*models.Notification, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	notification, err := r.queries.CreateNotification(ctx, db.CreateNotificationParams{
		UserID:  uid,
		Channel: channel,
		Message: message,
	})
	if err != nil {
		return nil, err
	}

	return &models.Notification{
		ID:          notification.ID,
		UserID:      notification.UserID,
		Channel:     notification.Channel,
		Message:     notification.Message,
		Status:      notification.Status,
		CreatedAt:   notification.CreatedAt,
		DeliveredAt: &notification.DeliveredAt.Time,
	}, nil
}

func (r *notificationRepositoryImpl) GetNotificationsByUserID(ctx context.Context, userID string, limit int32, offset int32) ([]*models.Notification, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	notifications, err := r.queries.GetNotificationsByUserID(ctx, db.GetNotificationsByUserIDParams{
		UserID: uid,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	result := make([]*models.Notification, len(notifications))
	for i, n := range notifications {
		result[i] = &models.Notification{
			ID:          n.ID,
			UserID:      n.UserID,
			Channel:     n.Channel,
			Message:     n.Message,
			Status:      n.Status,
			CreatedAt:   n.CreatedAt,
			DeliveredAt: &n.DeliveredAt.Time,
		}
	}

	return result, nil
}

func (r *notificationRepositoryImpl) UpdateNotificationStatus(ctx context.Context, id string, status string) error {
	nid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return r.queries.UpdateNotificationStatus(ctx, db.UpdateNotificationStatusParams{
		ID:     nid,
		Status: status,
	})
}
