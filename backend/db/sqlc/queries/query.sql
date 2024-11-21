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

-- ここから自作
-- パスワードリセット用のトークンを保存するクエリ

-- name: SaveResetToken :exec
-- params: email, token, expiry
INSERT INTO password_reset_tokens (email, token, expiry)
VALUES (?, ?, ?);


-- トークンを検証して対応するメールを取得するクエリ
-- name: ValidateResetToken :one
SELECT email FROM password_reset_tokens
WHERE token = ? AND expiry > NOW();

-- パスワードを更新するクエリ
-- name: UpdatePasswordByEmail :exec
UPDATE users
SET password_hash = ?, last_password_change = NOW()
WHERE email = ?;

-- 使用済みのリセットトークンを削除するクエリ
-- name: DeleteResetToken :exec
DELETE FROM password_reset_tokens
WHERE token = ?;

-- name: GetUserByEmail :one
SELECT id, password_hash FROM users WHERE username = ?;

-- name: UpdateUserProfile :exec
UPDATE users
SET 
    profile_image_url = COALESCE(?, profile_image_url),
    bio = COALESCE(?, bio),
    location = COALESCE(?, location),
    website = COALESCE(?, website),
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: UpdateUserSettings :exec
UPDATE users
SET 
    display_name = COALESCE(?, display_name),
    birth_date = COALESCE(?, birth_date),
    language = COALESCE(?, language),
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: UpdateUserNotifications :exec
UPDATE users
SET
    notification_settings = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: UpdateUserPrivacy :exec
UPDATE users
SET
    is_private = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: UpdateUserBanStatus :exec
UPDATE users
SET
    is_banned = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;