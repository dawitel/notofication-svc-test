// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: 20250207233845_long_union.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createAPIKey = `-- name: CreateAPIKey :one
INSERT INTO api_keys (user_id, key, name)
VALUES ($1, $2, $3)
RETURNING id, user_id, key, name, created_at, last_used_at, is_active
`

type CreateAPIKeyParams struct {
	UserID uuid.UUID `json:"user_id"`
	Key    string    `json:"key"`
	Name   string    `json:"name"`
}

func (q *Queries) CreateAPIKey(ctx context.Context, arg CreateAPIKeyParams) (ApiKey, error) {
	row := q.queryRow(ctx, q.createAPIKeyStmt, createAPIKey, arg.UserID, arg.Key, arg.Name)
	var i ApiKey
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Key,
		&i.Name,
		&i.CreatedAt,
		&i.LastUsedAt,
		&i.IsActive,
	)
	return i, err
}

const deactivateAPIKey = `-- name: DeactivateAPIKey :exec
UPDATE api_keys
SET is_active = false
WHERE id = $1 AND user_id = $2
`

type DeactivateAPIKeyParams struct {
	ID     uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
}

func (q *Queries) DeactivateAPIKey(ctx context.Context, arg DeactivateAPIKeyParams) error {
	_, err := q.exec(ctx, q.deactivateAPIKeyStmt, deactivateAPIKey, arg.ID, arg.UserID)
	return err
}

const getAPIKeyByKey = `-- name: GetAPIKeyByKey :one
SELECT id, user_id, key, name, created_at, last_used_at, is_active FROM api_keys
WHERE key = $1 AND is_active = true
`

func (q *Queries) GetAPIKeyByKey(ctx context.Context, key string) (ApiKey, error) {
	row := q.queryRow(ctx, q.getAPIKeyByKeyStmt, getAPIKeyByKey, key)
	var i ApiKey
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Key,
		&i.Name,
		&i.CreatedAt,
		&i.LastUsedAt,
		&i.IsActive,
	)
	return i, err
}

const updateAPIKeyLastUsed = `-- name: UpdateAPIKeyLastUsed :exec
UPDATE api_keys
SET last_used_at = NOW()
WHERE id = $1
`

func (q *Queries) UpdateAPIKeyLastUsed(ctx context.Context, id uuid.UUID) error {
	_, err := q.exec(ctx, q.updateAPIKeyLastUsedStmt, updateAPIKeyLastUsed, id)
	return err
}
