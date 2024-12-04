-- name: GetPostReposts :many
SELECT u.id, u.username, u.display_name, r.is_quote_repost, r.additional_comment
FROM reposts r
JOIN users u ON r.user_id = u.id
WHERE r.original_post_id = ?;

-- name: CreateRepost :exec
INSERT INTO reposts (id, user_id, original_post_id, is_quote_repost, additional_comment, reposted_at)
VALUES (?, ?, ?, ?, ?, CURRENT_TIMESTAMP);