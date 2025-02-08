-- name: CreateAPIKey :one
INSERT INTO api_keys (user_id, key, name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAPIKeyByKey :one
SELECT * FROM api_keys
WHERE key = $1 AND is_active = true;

-- name: UpdateAPIKeyLastUsed :exec
UPDATE api_keys
SET last_used_at = NOW()
WHERE id = $1;

-- name: DeactivateAPIKey :exec
UPDATE api_keys
SET is_active = false
WHERE id = $1 AND user_id = $2;