package apikey_repository

import (
	"context"
	"database/sql"
	"notification-service/internal/db"
	"notification-service/internal/models"

	"github.com/google/uuid"
)

type apiKeyRepositoryImpl struct {
	queries *db.Queries
}

func NewAPIKeyRepository(conn *sql.DB) APIKeyRepository {
	return &apiKeyRepositoryImpl{
		queries: db.New(conn),
	}
}

func (r *apiKeyRepositoryImpl) CreateAPIKey(ctx context.Context, userID string, key string, name string) (*models.APIKey, error) {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	apiKey, err := r.queries.CreateAPIKey(ctx, db.CreateAPIKeyParams{
		UserID: uid,
		Key:    key,
		Name:   name,
	})
	if err != nil {
		return nil, err
	}

	return &models.APIKey{
		ID:         apiKey.ID,
		UserID:     apiKey.UserID,
		Key:        apiKey.Key,
		Name:       apiKey.Name,
		CreatedAt:  apiKey.CreatedAt,
		LastUsedAt: &apiKey.LastUsedAt.Time,
		IsActive:   apiKey.IsActive,
	}, nil
}

func (r *apiKeyRepositoryImpl) GetAPIKeyByKey(ctx context.Context, key string) (*models.APIKey, error) {
	apiKey, err := r.queries.GetAPIKeyByKey(ctx, key)
	if err != nil {
		return nil, err
	}

	return &models.APIKey{
		ID:         apiKey.ID,
		UserID:     apiKey.UserID,
		Key:        apiKey.Key,
		Name:       apiKey.Name,
		CreatedAt:  apiKey.CreatedAt,
		LastUsedAt: &apiKey.LastUsedAt.Time,
		IsActive:   apiKey.IsActive,
	}, nil
}

func (r *apiKeyRepositoryImpl) UpdateAPIKeyLastUsed(ctx context.Context, id string) error {
	keyID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return r.queries.UpdateAPIKeyLastUsed(ctx, keyID)
}

func (r *apiKeyRepositoryImpl) DeactivateAPIKey(ctx context.Context, id string, userID string) error {
	keyID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	uid, err := uuid.Parse(userID)
	if err != nil {
		return err
	}

	return r.queries.DeactivateAPIKey(ctx, db.DeactivateAPIKeyParams{
		ID:     keyID,
		UserID: uid,
	})
}
