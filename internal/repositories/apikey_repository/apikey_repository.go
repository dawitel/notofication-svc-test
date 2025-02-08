package apikey_repository

import (
	"context"
	"notification-service/internal/models"
)

type APIKeyRepository interface {
	CreateAPIKey(ctx context.Context, userID string, key string, name string) (*models.APIKey, error)
	GetAPIKeyByKey(ctx context.Context, key string) (*models.APIKey, error)
	UpdateAPIKeyLastUsed(ctx context.Context, id string) error
	DeactivateAPIKey(ctx context.Context, id string, userID string) error
}