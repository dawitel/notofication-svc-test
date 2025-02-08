package user_repository

import (
	"context"
	"database/sql"
	"notification-service/internal/db"
	"notification-service/internal/models"

	"github.com/google/uuid"
)

type userRepositoryImpl struct {
	queries *db.Queries
}

func NewUserRepository(conn *sql.DB) UserRepository {
	return &userRepositoryImpl{
		queries: db.New(conn),
	}
}

func (r *userRepositoryImpl) CreateUser(ctx context.Context, email string, passwordHash string) (*models.User, error) {
	user, err := r.queries.CreateUser(ctx, db.CreateUserParams{
		Email:        email,
		PasswordHash: passwordHash,
	})
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
	}, nil
}

func (r *userRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
	}, nil
}

func (r *userRepositoryImpl) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	user, err := r.queries.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
	}, nil
}
