-- name: CreateUser :execresult
INSERT INTO users (id, firebase_uid, email, password_hash, username, display_name, created_at)
VALUES (?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP);

-- name: DeleteUser :execresult
DELETE FROM users WHERE id = ?;

-- name: GetEmailFromUsername :one
SELECT email
FROM users
WHERE username = ?;

-- name: GetRecentPosts :many
SELECT p.*, u.username, u.display_name
FROM posts p
JOIN users u ON p.user_id = u.id
WHERE p.is_deleted = FALSE
ORDER BY p.created_at DESC
LIMIT ?;

-- name: GetUserById :one
SELECT id, firebase_uid, email, username, display_name, bio, location, followers_count, following_count, posts_count
FROM users
WHERE id = ?;

-- name: AddLike :exec
INSERT INTO likes (id, user_id, post_id)
VALUES (?, ?, ?);

-- name: CreateRepost :exec
INSERT INTO reposts (id, user_id, original_post_id, is_quote_repost, additional_comment)
VALUES (?, ?, ?, ?, ?);



-- name: AddBlock :exec
INSERT INTO blocks (id, blocked_by_id, blocked_user_id)
VALUES (?, ?, ?);

-- name: CreateNotification :exec
INSERT INTO notifications (id, user_id, type, message)
VALUES (?, ?, ?, ?);

-- name: SendDM :exec
INSERT INTO dms (id, sender_id, receiver_id, content)
VALUES (?, ?, ?, ?);



-- name: UpdatePostLikesCount :exec
UPDATE posts
SET likes_count = (
    SELECT COUNT(*) FROM likes WHERE post_id = posts.id
)
WHERE posts.id = ?;

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
WHERE user_id = ? AND is_read = FALSE
ORDER BY created_at DESC;

-- name: GetDMConversation :many
SELECT *
FROM dms
WHERE (sender_id = ? AND receiver_id = ?)
   OR (sender_id = ? AND receiver_id = ?)
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
-- name: UpdatePassword :exec
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

-- name: UpdateUserName :exec
UPDATE users
SET
    username = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: UpdateUserEmail :exec
UPDATE users
SET
    email = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: CreatePost :exec
INSERT INTO posts (
    id, user_id, content, media_urls, visibility, 
    original_post_id, reply_to_id, root_post_id, is_repost, is_reply, created_at
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP);


-- name: DeletePost :exec
UPDATE posts
SET 
    is_deleted = TRUE,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;


-- name: AddFollow :exec
INSERT INTO follows (id, follower_id, following_id, created_at)
VALUES (?, ?, ?, CURRENT_TIMESTAMP);

-- name: RemoveFollow :execresult
DELETE FROM follows
WHERE follower_id = ? AND following_id = ?;

-- name: GetFollowStatus :one
SELECT EXISTS(
    SELECT 1
    FROM follows
    WHERE follower_id = ? AND following_id = ?
) AS following;


-- name: GetIDfromFirebaseUID :one
SELECT id FROM users WHERE firebase_uid = ?;

-- name: GetFirebaseUIDfromUserID :one
SELECT firebase_uid FROM users WHERE id = ?;


-- name: UpdateFollowersCount :execresult
UPDATE users
SET followers_count = (
    SELECT COUNT(*) FROM follows WHERE following_id = users.id
)
WHERE users.id = ?;

-- name: GetFollowersCount :one
SELECT followers_count FROM users WHERE id = ?;

-- name: UpdateFollowingsCount :execresult
UPDATE users
SET following_count = (
    SELECT COUNT(*) FROM follows WHERE following_id = users.id
)
WHERE users.id = ?;

-- name: GetFollowingsCount :one
SELECT following_count FROM users WHERE id = ?;

-- name: GetFollowers :many
SELECT u.id, u.username, u.display_name
FROM follows f
JOIN users u ON f.follower_id = u.id
WHERE f.following_id = ?;

-- name: GetFollowings :many
SELECT u.id, u.username, u.display_name
FROM follows f
JOIN users u ON f.following_id = u.id
WHERE f.follower_id = ?;

-- name: GetFollowersAndFollowings :many
SELECT u.id, u.username, u.display_name, f.follower_id, f.following_id
FROM follows f
JOIN users u ON f.follower_id = u.id
WHERE f.following_id = ? OR f.follower_id = ?;

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

-- name: GetPostLikes :many
SELECT u.id, u.username, u.display_name
FROM likes l
JOIN users u ON l.user_id = u.id
WHERE l.post_id = ?;

-- name: GetPostReposts :many
SELECT u.id, u.username, u.display_name, r.is_quote_repost, r.additional_comment
FROM reposts r
JOIN users u ON r.user_id = u.id
WHERE r.original_post_id = ?;

-- name: GetPostReplies :many
SELECT p.*, u.username, u.display_name
FROM posts p
JOIN users u ON p.user_id = u.id
WHERE p.root_post_id = ?
ORDER BY p.created_at ASC;

-- name: GetPostRepliesCount :one
SELECT COUNT(*) AS reply_count
FROM posts
WHERE root_post_id = ?;

-- name: IncrementReplyCount :exec
UPDATE posts
SET replies_count = replies_count + 1
WHERE id = ?;

-- name: DecrementReplyCount :exec
UPDATE posts
SET replies_count = replies_count - 1
WHERE (id = ? AND replies_count > 0);

-- name: GetReplyToID :one
SELECT reply_to_id
FROM posts
WHERE id = ?;

-- name: RestorePost :execresult
UPDATE posts
SET is_deleted = false
WHERE id = ? AND is_deleted = true AND TIMESTAMPDIFF(MINUTE, updated_at, NOW()) <= 20;
