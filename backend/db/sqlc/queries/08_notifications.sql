

-- name: GetUnreadNotifications :many
SELECT *
FROM notifications
WHERE user_id = ? AND is_read = FALSE
ORDER BY created_at DESC;

-- name: CreateNotification :exec
INSERT INTO notifications (id, user_id, type, message, is_read)
VALUES (?, ?, ?, ?, FALSE);

-- name: GetNotificarionsByUserID :many
SELECT id, user_id, type, message, created_at, is_read
FROM notifications
WHERE user_id = ?
ORDER BY created_at DESC;

-- name: MarkNotificationsAsRead :execresult
UPDATE notifications
SET is_read = TRUE
WHERE id = ?;

-- name: DeleteNotification :execresult
DELETE FROM notifications
WHERE id = ?;

-- name: CountUnreadNotifications :one
SELECT COUNT(*)
FROM notifications
WHERE user_id = ? AND is_read = FALSE;

-- name: CountAllNotifications :one
SELECT COUNT(*)
FROM notifications
WHERE user_id = ?;

-- name: GetNotificationByID :one
SELECT *
FROM notifications
WHERE id = ?;

-- name: GetNotificationByUserIDAndID :one
SELECT *
FROM notifications
WHERE user_id = ? AND id = ?;

-- name: GetNotificationByUserIDAndType :one
SELECT *
FROM notifications
WHERE user_id = ? AND type = ?;
