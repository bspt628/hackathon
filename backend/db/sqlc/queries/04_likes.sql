-- name: GetPostLikes :many
SELECT u.id, u.username, u.display_name
FROM likes l
JOIN users u ON l.user_id = u.id
WHERE l.post_id = ?;

-- name: AddLike :exec
INSERT INTO likes (id, user_id, post_id)
VALUES (?, ?, ?);

-- name: RemoveLike :exec
DELETE FROM likes
WHERE user_id = ? AND post_id = ?;

-- name: GetLikeStatus :one
SELECT EXISTS(
    SELECT 1
    FROM likes
    WHERE user_id = ? AND post_id = ?
) AS liked;

-- name: UpdateLikesCount :execresult
UPDATE posts
SET likes_count = (
    SELECT COUNT(*) FROM likes WHERE post_id = posts.id
)
WHERE posts.id = ?;

-- name: GetLikesCount :one
SELECT likes_count FROM posts WHERE id = ?;

-- name: GetLikes :many
SELECT u.id, u.username, u.display_name
FROM likes l
JOIN users u ON l.user_id = u.id
WHERE l.post_id = ?;

-- name: IncrementLikesCount :exec
UPDATE posts
SET likes_count = likes_count + 1
WHERE id = ?;

-- name: DecrementLikesCount :exec
UPDATE posts
SET likes_count = likes_count - 1
WHERE (id = ? AND likes_count > 0);

-- name: GetLikeID :one
SELECT id
FROM likes
WHERE user_id = ? AND post_id = ?;