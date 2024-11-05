-- name: GetUser :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username;

-- name: CreateUser :execresult
INSERT INTO users (id, username, email, password, status)
VALUES (?, ?, ?, ?, ?);

-- name: UpdateUser :exec
UPDATE users
SET username = ?, email = ?, password = ?, status = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;

-- name: GetTweet :one
SELECT * FROM tweets
WHERE id = ? LIMIT 1;

-- name: ListTweets :many
SELECT * FROM tweets
WHERE user_id = ?
ORDER BY created_at DESC;

-- name: CreateTweet :execresult
INSERT INTO tweets (id, user_id, content)
VALUES (?, ?, ?);

-- name: DeleteTweet :exec
UPDATE tweets
SET is_deleted = TRUE
WHERE id = ?;

-- name: CreateRetweet :execresult
INSERT INTO retweets (id, user_id, tweet_id)
VALUES (?, ?, ?);

-- name: DeleteRetweet :exec
DELETE FROM retweets
WHERE id = ?;

-- name: CreateLike :execresult
INSERT INTO likes (id, user_id, tweet_id)
VALUES (?, ?, ?);

-- name: DeleteLike :exec
DELETE FROM likes
WHERE id = ?;

-- name: CreateReply :execresult
INSERT INTO replies (id, user_id, tweet_id, content)
VALUES (?, ?, ?, ?);

-- name: ListReplies :many
SELECT * FROM replies
WHERE tweet_id = ?
ORDER BY created_at;

-- name: CreateFollow :execresult
INSERT INTO follows (id, follower_id, following_id)
VALUES (?, ?, ?);

-- name: DeleteFollow :exec
DELETE FROM follows
WHERE follower_id = ? AND following_id = ?;

-- name: CreateBlock :execresult
INSERT INTO blocks (id, blocked_by_id, blocked_user_id)
VALUES (?, ?, ?);

-- name: DeleteBlock :exec
DELETE FROM blocks
WHERE blocked_by_id = ? AND blocked_user_id = ?;

-- name: GetProfile :one
SELECT * FROM profiles
WHERE user_id = ? LIMIT 1;

-- name: UpdateProfile :exec
UPDATE profiles
SET bio = ?, location = ?, website = ?, is_private = ?
WHERE user_id = ?;

-- name: ListNotifications :many
SELECT * FROM notifications
WHERE user_id = ?
ORDER BY created_at DESC
LIMIT ?;

-- name: MarkNotificationAsRead :exec
UPDATE notifications
SET is_read = TRUE
WHERE id = ?;

-- name: CreateDM :execresult
INSERT INTO dms (id, sender_id, receiver_id, content)
VALUES (?, ?, ?, ?);

-- name: ListDMs :many
SELECT * FROM dms
WHERE (sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)
ORDER BY created_at DESC
LIMIT ?;