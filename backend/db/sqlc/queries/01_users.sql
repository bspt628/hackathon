-- name: CreateUser :exec
INSERT INTO users (id, firebase_uid, email, password_hash, username, display_name)
VALUES (?, ?, ?, ?, ?, ?);

-- name: DeleteUser :execresult
DELETE FROM users WHERE id = ?;

-- name: GetEmailFromUsername :one
SELECT email
FROM users
WHERE username = ?;

-- name: GetUser :one
SELECT id, firebase_uid, email, username, display_name, bio, location, followers_count, following_count, posts_count
FROM users
WHERE id = ?;

-- name: GetUserPasswordFromUsername :one
SELECT id, password_hash 
FROM users
WHERE username = ?;

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

-- name: UpdateUserProfile :exec
UPDATE users
SET 
    profile_image_url = COALESCE(?, profile_image_url),
    bio = COALESCE(?, bio),
    location = COALESCE(?, location),
    website = COALESCE(?, website)
WHERE id = ?;

-- name: UpdateUserSettings :exec
UPDATE users
SET 
    display_name = COALESCE(?, display_name),
    birth_date = COALESCE(?, birth_date),
    language = COALESCE(?, language)
WHERE id = ?;

-- name: UpdateUserNotifications :exec
UPDATE users
SET
    notification_settings = ?
WHERE id = ?;

-- name: UpdateUserPrivacy :exec
UPDATE users
SET
    is_private = ?
WHERE id = ?;

-- name: UpdateUserBanStatus :exec
UPDATE users
SET
    is_banned = ?
WHERE id = ?;

-- name: UpdateUserUsername :exec
UPDATE users
SET
    username = ?
WHERE id = ?;

-- name: UpdateUserEmail :exec
UPDATE users
SET
    email = ?
WHERE id = ?;

-- name: GetIDfromFirebaseUID :one
SELECT id FROM users WHERE firebase_uid = ?;