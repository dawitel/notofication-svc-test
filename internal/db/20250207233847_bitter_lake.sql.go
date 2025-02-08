// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: 20250207233847_bitter_lake.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createNotification = `-- name: CreateNotification :one
INSERT INTO notifications (user_id, channel, message)
VALUES ($1, $2, $3)
RETURNING id, user_id, channel, message, status, created_at, delivered_at
`

type CreateNotificationParams struct {
	UserID  uuid.UUID `json:"user_id"`
	Channel string    `json:"channel"`
	Message string    `json:"message"`
}

func (q *Queries) CreateNotification(ctx context.Context, arg CreateNotificationParams) (Notification, error) {
	row := q.queryRow(ctx, q.createNotificationStmt, createNotification, arg.UserID, arg.Channel, arg.Message)
	var i Notification
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Channel,
		&i.Message,
		&i.Status,
		&i.CreatedAt,
		&i.DeliveredAt,
	)
	return i, err
}

const getNotificationsByUserID = `-- name: GetNotificationsByUserID :many
SELECT id, user_id, channel, message, status, created_at, delivered_at FROM notifications
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetNotificationsByUserIDParams struct {
	UserID uuid.UUID `json:"user_id"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

func (q *Queries) GetNotificationsByUserID(ctx context.Context, arg GetNotificationsByUserIDParams) ([]Notification, error) {
	rows, err := q.query(ctx, q.getNotificationsByUserIDStmt, getNotificationsByUserID, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Notification
	for rows.Next() {
		var i Notification
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Channel,
			&i.Message,
			&i.Status,
			&i.CreatedAt,
			&i.DeliveredAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateNotificationStatus = `-- name: UpdateNotificationStatus :exec
UPDATE notifications
SET status = $2, delivered_at = CASE WHEN $2 = 'delivered' THEN NOW() ELSE NULL END
WHERE id = $1
`

type UpdateNotificationStatusParams struct {
	ID     uuid.UUID `json:"id"`
	Status string    `json:"status"`
}

func (q *Queries) UpdateNotificationStatus(ctx context.Context, arg UpdateNotificationStatusParams) error {
	_, err := q.exec(ctx, q.updateNotificationStatusStmt, updateNotificationStatus, arg.ID, arg.Status)
	return err
}
