package user_repository

import (
	"context"
	"notification-service/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, email string, passwordHash string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
}