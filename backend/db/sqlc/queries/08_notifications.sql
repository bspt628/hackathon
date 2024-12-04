

-- name: GetUnreadNotifications :many
SELECT *
FROM notifications
WHERE user_id = ? AND is_read = FALSE
ORDER BY created_at DESC;

-- name: CreateNotification :exec
INSERT INTO notifications (id, user_id, type, message)
VALUES (?, ?, ?, ?);