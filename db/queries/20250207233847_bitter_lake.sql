-- name: CreateNotification :one
INSERT INTO notifications (user_id, channel, message)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetNotificationsByUserID :many
SELECT * FROM notifications
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: UpdateNotificationStatus :exec
UPDATE notifications
SET status = $2, delivered_at = CASE WHEN $2 = 'delivered' THEN NOW() ELSE NULL END
WHERE id = $1;