-- name: GetPostReplies :many
SELECT p.*, u.username, u.display_name
FROM posts p
JOIN users u ON p.user_id = u.id
WHERE p.root_post_id = ?
ORDER BY p.created_at ASC;

-- name: IncrementReplyCount :exec
UPDATE posts
SET replies_count = replies_count + 1
WHERE id = ?;

-- name: DecrementReplyCount :exec
UPDATE posts
SET replies_count = replies_count - 1
WHERE (id = ? AND replies_count > 0);