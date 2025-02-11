// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateAPIKey(ctx context.Context, arg CreateAPIKeyParams) (ApiKey, error)
	CreateNotification(ctx context.Context, arg CreateNotificationParams) (Notification, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeactivateAPIKey(ctx context.Context, arg DeactivateAPIKeyParams) error
	GetAPIKeyByKey(ctx context.Context, key string) (ApiKey, error)
	GetNotificationsByUserID(ctx context.Context, arg GetNotificationsByUserIDParams) ([]Notification, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (User, error)
	UpdateAPIKeyLastUsed(ctx context.Context, id uuid.UUID) error
	UpdateNotificationStatus(ctx context.Context, arg UpdateNotificationStatusParams) error
}

var _ Querier = (*Queries)(nil)
