-- name: CreateUser :execresult
INSERT INTO users (id, email, password_hash, username, display_name)
VALUES (?, ?, ?, ?, ?);

-- name: UpdateUserInfo :exec
UPDATE users
SET bio = ?, location = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;

-- name: GetEmailFromUsername :one
SELECT email
FROM users
WHERE username = ?;

-- name: CreatePost :execresult
INSERT INTO posts (id, user_id, content)
VALUES (?, ?, ?);

-- name: GetRecentPosts :many
SELECT p.*, u.username, u.display_name
FROM posts p
JOIN users u ON p.user_id = u.id
WHERE p.is_deleted = FALSE
ORDER BY p.created_at DESC
LIMIT ?;

-- name: GetUserById :one
SELECT id, email, username, display_name, bio, location, followers_count, following_count, posts_count
FROM users
WHERE id = ?;

-- name: AddLike :exec
INSERT INTO likes (id, userId, postId)
VALUES (?, ?, ?);

-- name: CreateRepost :exec
INSERT INTO reposts (id, user_id, original_post_id, is_quote_repost, additional_comment)
VALUES (?, ?, ?, ?, ?);

-- name: AddFollow :exec
INSERT INTO follows (id, followerId, followingId)
VALUES (?, ?, ?);

-- name: AddBlock :exec
INSERT INTO blocks (id, blockedById, blockedUserId)
VALUES (?, ?, ?);

-- name: CreateNotification :exec
INSERT INTO notifications (id, userId, type, message)
VALUES (?, ?, ?, ?);

-- name: SendDM :exec
INSERT INTO dms (id, senderId, receiverId, content)
VALUES (?, ?, ?, ?);

-- name: UpdateFollowersCount :exec
UPDATE users
SET followers_count = (
    SELECT COUNT(*) FROM follows WHERE followingId = users.id
)
WHERE users.id = ?;

-- name: UpdatePostLikesCount :exec
UPDATE posts
SET likes_count = (
    SELECT COUNT(*) FROM likes WHERE postId = posts.id
)
WHERE posts.id = ?;

-- name: GetUserTimeline :many
SELECT p.*, u.username, u.display_name
FROM posts p
JOIN users u ON p.user_id = u.id
WHERE p.user_id IN (
    SELECT followingId
    FROM follows
    WHERE followerId = ?
) OR p.user_id = ?
ORDER BY p.created_at DESC
LIMIT ?;

-- name: SearchPostsByHashtag :many
SELECT p.*, u.username, u.display_name
FROM posts p
JOIN users u ON p.user_id = u.id
WHERE p.content LIKE ?
ORDER BY p.created_at DESC
LIMIT ?;

-- name: GetUserStats :one
SELECT
    u.id,
    u.username,
    u.followers_count,
    u.following_count,
    u.posts_count,
    COUNT(DISTINCT l.id) AS total_likes_received
FROM users u
LEFT JOIN posts p ON u.id = p.user_id
LEFT JOIN likes l ON p.id = l.postId
WHERE u.id = ?
GROUP BY u.id;

-- name: GetUnreadNotifications :many
SELECT *
FROM notifications
WHERE userId = ? AND isRead = FALSE
ORDER BY createdAt DESC;

-- name: GetDMConversation :many
SELECT *
FROM dms
WHERE (senderId = ? AND receiverId = ?)
   OR (senderId = ? AND receiverId = ?)
ORDER BY createdAt ASC;

-- パスワードリセット用のトークンを保存するクエリ
-- name: SaveResetToken :exec
INSERT INTO password_reset_tokens (email, token, expiry)
VALUES ($1, $2, $3);

-- トークンを検証して対応するメールを取得するクエリ
-- name: ValidateResetToken :one
SELECT email FROM password_reset_tokens
WHERE token = $1 AND expiry > NOW();

-- パスワードを更新するクエリ
-- name: UpdatePasswordByEmail :exec
UPDATE users
SET password_hash = $2, last_password_change = NOW()
WHERE email = $1;

-- 使用済みのリセットトークンを削除するクエリ
-- name: DeleteResetToken :exec
DELETE FROM password_reset_tokens
WHERE token = $1;