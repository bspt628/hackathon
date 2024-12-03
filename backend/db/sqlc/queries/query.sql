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

-- name: AddLike :execresult
INSERT INTO likes (id, user_id, post_id, created_at)
VALUES (?, ?, ?, CURRENT_TIMESTAMP);

-- name: IncrementLikesCount :exec
UPDATE posts
SET likes_count = likes_count + 1
WHERE id = ?;

-- name: RemoveLike :exec
DELETE FROM likes
WHERE user_id = ? AND post_id = ?;

-- name: CheckLikeExists :one
SELECT EXISTS (
    SELECT 1
    FROM likes
    WHERE user_id = ? AND post_id = ?
) AS liked;

-- name: DecrementLikesCount :exec
UPDATE posts
SET likes_count = likes_count - 1
WHERE id = ?;

-- name: GetPostIDFromLike :one
SELECT post_id
FROM likes
WHERE id = ?;

-- name: GetLikeID :one
SELECT id
FROM likes
WHERE user_id = ? AND post_id = ?;

-- name: GetUserLikes :many
SELECT post_id
FROM likes
WHERE user_id = ?
ORDER BY created_at DESC
LIMIT ?;

-- name: GetLikeStatus :one
SELECT EXISTS(
    SELECT 1
    FROM likes
    WHERE user_id = ? AND post_id = ?
) AS liked;

-- name: GetPostLikesCount :one
SELECT likes_count
FROM posts
WHERE id = ?;

-- name: GetAllPosts :many
SELECT p.*, u.username, u.display_name
FROM posts p
JOIN users u ON p.user_id = u.id
ORDER BY p.created_at DESC
LIMIT ?;


-- name: GetFollowingUsersPosts :many
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

-- name: SearchPosts :many
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