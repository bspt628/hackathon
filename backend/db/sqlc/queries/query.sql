-- name: CreateUser :execresult
-- 実装済み
INSERT INTO users (id, firebase_uid, email, password_hash, username, display_name, created_at)
VALUES (?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP);

-- name: DeleteUser :execresult
-- 実装済み
DELETE FROM users WHERE id = ?;

-- name: GetEmailFromUsername :one
-- 実装済み
SELECT email
FROM users
WHERE username = ?;

-- name: GetUser :one
-- 実装済み
SELECT id, firebase_uid, email, username, display_name, bio, location, followers_count, following_count, posts_count
FROM users
WHERE id = ?;

-- name: SaveResetToken :exec
-- 実装済み
INSERT INTO password_reset_tokens (email, token, expiry)
VALUES (?, ?, ?);

-- name: ValidateResetToken :one
-- 実装済み
SELECT email FROM password_reset_tokens
WHERE token = ? AND expiry > NOW();

-- name: UpdatePassword :exec
-- 実装済み
UPDATE users
SET 
    password_hash = ?,
    last_password_change = NOW(),
    updated_at = CURRENT_TIMESTAMP
WHERE email = ?;

-- name: DeleteResetToken :exec
-- 実装済み
DELETE FROM password_reset_tokens
WHERE token = ?;

-- name: SignInCheck :one
-- 実装済み
SELECT id, password_hash FROM users WHERE username = ?;

-- name: UpdateUserProfile :exec
-- 実装済み
UPDATE users
SET 
    profile_image_url = COALESCE(?, profile_image_url),
    bio = COALESCE(?, bio),
    location = COALESCE(?, location),
    website = COALESCE(?, website),
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: UpdateUserSettings :exec
-- 実装済み
UPDATE users
SET 
    display_name = COALESCE(?, display_name),
    birth_date = COALESCE(?, birth_date),
    language = COALESCE(?, language),
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: UpdateUserNotifications :exec
-- 実装済み
UPDATE users
SET
    notification_settings = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: UpdateUserPrivacy :exec
-- 実装済み
UPDATE users
SET
    is_private = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: UpdateUserBanStatus :exec
-- 実装済み
UPDATE users
SET
    is_banned = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: UpdateUserUsername :exec
-- 実装済み 
UPDATE users
SET
    username = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: UpdateUserEmail :exec
-- 実装済み
UPDATE users
SET
    email = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: CreatePost :exec
-- 実装済み
INSERT INTO posts (
    id, user_id, content, media_urls, visibility, 
    original_post_id, reply_to_id, root_post_id, is_repost, is_reply, created_at
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP);


-- name: DeletePost :exec
-- 実装済み
UPDATE posts
SET 
    is_deleted = TRUE,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?;


-- name: AddFollow :exec
-- 実装済み
INSERT INTO follows (id, follower_id, following_id, created_at)
VALUES (?, ?, ?, CURRENT_TIMESTAMP);

-- name: RemoveFollow :execresult
-- 実装済み
DELETE FROM follows
WHERE follower_id = ? AND following_id = ?;

-- name: GetFollowStatus :one
-- 実装済み
SELECT EXISTS(
    SELECT 1
    FROM follows
    WHERE follower_id = ? AND following_id = ?
) AS following;


-- name: GetIDfromFirebaseUID :one
-- 実装済み
SELECT id FROM users WHERE firebase_uid = ?;

-- name: UpdateFollowersCount :execresult
-- 実装済み
UPDATE users
SET followers_count = (
    SELECT COUNT(*) FROM follows WHERE following_id = users.id
)
WHERE users.id = ?;

-- name: GetFollowersCount :one
-- 実装済み
SELECT followers_count FROM users WHERE id = ?;

-- name: UpdateFollowingsCount :execresult
-- 実装済み
UPDATE users
SET following_count = (
    SELECT COUNT(*) FROM follows WHERE following_id = users.id
)
WHERE users.id = ?;

-- name: GetFollowingsCount :one
-- 実装済み
SELECT following_count FROM users WHERE id = ?;

-- name: GetFollowers :many
-- 実装済み
SELECT u.id, u.username, u.display_name
FROM follows f
JOIN users u ON f.follower_id = u.id
WHERE f.following_id = ?;

-- name: GetFollowings :many
-- 実装済み
SELECT u.id, u.username, u.display_name
FROM follows f
JOIN users u ON f.following_id = u.id
WHERE f.follower_id = ?;

-- name: GetFollowersAndFollowings :many
-- 実装済み
SELECT u.id, u.username, u.display_name, f.follower_id, f.following_id
FROM follows f
JOIN users u ON f.follower_id = u.id
WHERE f.following_id = ? OR f.follower_id = ?;
