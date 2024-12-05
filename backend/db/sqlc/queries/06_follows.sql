-- name: AddFollow :exec
INSERT INTO follows (id, follower_id, following_id)
VALUES (?, ?, ?);

-- name: RemoveFollow :execresult
DELETE FROM follows
WHERE follower_id = ? AND following_id = ?;

-- name: GetFollowStatus :one
SELECT EXISTS(
    SELECT 1
    FROM follows
    WHERE follower_id = ? AND following_id = ?
) AS following;

-- name: UpdateFollowersCount :execresult
-- いらないかも
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

-- name: IncrementFollowersCount :execresult
UPDATE users
SET followers_count = followers_count + 1
WHERE id = ?;

-- name: DecrementFollowersCount :execresult
UPDATE users
SET followers_count = followers_count - 1
WHERE id = ?;

-- name: IncrementFollowingsCount :execresult
UPDATE users
SET following_count = following_count + 1
WHERE id = ?;

-- name: DecrementFollowingsCount :execresult
UPDATE users
SET following_count = following_count - 1
WHERE id = ?;

