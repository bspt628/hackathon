-- name: CreatePost :exec
INSERT INTO posts (
    id, user_id, content, media_urls, visibility, 
    original_post_id, reply_to_id, root_post_id, is_repost, is_reply
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);


-- name: DeletePost :exec
UPDATE posts
SET 
    is_deleted = TRUE
WHERE id = ?;

-- name: CheckPostExists :one
SELECT EXISTS (
    SELECT 1 
    FROM posts 
    WHERE id = ?
);


-- name: CheckRootPostValidity :one
SELECT root_post_id IS NULL AS is_valid
FROM posts
WHERE id = ?;

-- name: CountReplyPosts :one
SELECT COUNT(*) AS reply_count
FROM posts
WHERE root_post_id = ?;

-- name: GetPost :one
SELECT p.*, u.username, u.display_name
FROM posts p
JOIN users u ON p.user_id = u.id
WHERE p.id = ?;

-- name: GetReplyToID :one
SELECT reply_to_id
FROM posts
WHERE id = ?;

-- name: RestorePost :execresult
UPDATE posts
SET is_deleted = false
WHERE id = ? AND is_deleted = true AND TIMESTAMPDIFF(MINUTE, updated_at, CURRENT_TIMESTAMP) <= 20;

-- name: GetUserTimeline :many
SELECT p.*, u.username, u.display_name
FROM posts p
JOIN users u ON p.user_id = u.id
WHERE p.user_id IN (
    SELECT following_id
    FROM follows
    WHERE follower_id = ?
) OR p.user_id = ?
ORDER BY p.created_at DESC
LIMIT ?;

-- name: GetAllPosts :many
SELECT p.*, u.username, u.display_name
FROM posts p
JOIN users u ON p.user_id = u.id
ORDER BY p.created_at DESC
LIMIT ?;

-- name: GetFollowedPosts :many
SELECT p.*, u.username, u.display_name
FROM posts p
JOIN users u ON p.user_id = u.id
WHERE p.user_id IN (
    SELECT following_id
    FROM follows
    WHERE follower_id = ?
)
ORDER BY p.created_at DESC
LIMIT ?;