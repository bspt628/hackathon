-- name: AddBlock :exec
INSERT INTO blocks (id, blocked_by_id, blocked_user_id)
VALUES (?, ?, ?);

-- name: RemoveBlock :exec
DELETE FROM blocks
WHERE blocked_by_id = ? AND blocked_user_id = ?;

-- name: GetBlockStatus :one
SELECT EXISTS(
    SELECT 1
    FROM blocks
    WHERE blocked_by_id = ? AND blocked_user_id = ?
) AS blocked;

-- name: GetBlockedUsers :many
SELECT u.id, u.username, u.display_name
FROM blocks b
JOIN users u ON b.blocked_user_id = u.id
WHERE b.blocked_by_id = ?;