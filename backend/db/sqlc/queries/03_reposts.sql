-- name: CreateRepost :exec
INSERT INTO reposts (id, user_id, original_post_id, is_quote_repost, additional_comment)
VALUES (?, ?, ?, ?, ?);

-- name: DeleteRepost :exec
DELETE FROM reposts
WHERE user_id = ? AND original_post_id = ?;

-- name: IncrementRepostsCount :execresult
UPDATE posts
SET reposts_count = reposts_count + 1
WHERE id = ?;

-- name: DecrementRepostsCount :execresult
UPDATE posts
SET reposts_count = reposts_count - 1
WHERE id = ?;

-- name: GetPostReposts :many
SELECT u.id, u.username, u.display_name, r.is_quote_repost, r.additional_comment
FROM reposts r
JOIN users u ON r.user_id = u.id
WHERE r.original_post_id = ?;

-- name: GetRepostStatus :one
SELECT EXISTS(
    SELECT 1
    FROM reposts
    WHERE user_id = ? AND original_post_id = ?
) AS reposting;
