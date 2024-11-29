// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package sqlc

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

const addBlock = `-- name: AddBlock :exec
INSERT INTO blocks (id, blocked_by_id, blocked_user_id)
VALUES (?, ?, ?)
`

type AddBlockParams struct {
	ID            string         `json:"id"`
	BlockedByID   sql.NullString `json:"blocked_by_id"`
	BlockedUserID sql.NullString `json:"blocked_user_id"`
}

func (q *Queries) AddBlock(ctx context.Context, arg AddBlockParams) error {
	_, err := q.db.ExecContext(ctx, addBlock, arg.ID, arg.BlockedByID, arg.BlockedUserID)
	return err
}

const addFollow = `-- name: AddFollow :exec
INSERT INTO follows (id, follower_id, following_id, created_at)
VALUES (?, ?, ?, CURRENT_TIMESTAMP)
`

type AddFollowParams struct {
	ID          string         `json:"id"`
	FollowerID  sql.NullString `json:"follower_id"`
	FollowingID sql.NullString `json:"following_id"`
}

func (q *Queries) AddFollow(ctx context.Context, arg AddFollowParams) error {
	_, err := q.db.ExecContext(ctx, addFollow, arg.ID, arg.FollowerID, arg.FollowingID)
	return err
}

const addLike = `-- name: AddLike :exec
INSERT INTO likes (id, user_id, post_id)
VALUES (?, ?, ?)
`

type AddLikeParams struct {
	ID     string         `json:"id"`
	UserID sql.NullString `json:"user_id"`
	PostID sql.NullString `json:"post_id"`
}

func (q *Queries) AddLike(ctx context.Context, arg AddLikeParams) error {
	_, err := q.db.ExecContext(ctx, addLike, arg.ID, arg.UserID, arg.PostID)
	return err
}

const createNotification = `-- name: CreateNotification :exec
INSERT INTO notifications (id, user_id, type, message)
VALUES (?, ?, ?, ?)
`

type CreateNotificationParams struct {
	ID      string         `json:"id"`
	UserID  sql.NullString `json:"user_id"`
	Type    sql.NullString `json:"type"`
	Message sql.NullString `json:"message"`
}

func (q *Queries) CreateNotification(ctx context.Context, arg CreateNotificationParams) error {
	_, err := q.db.ExecContext(ctx, createNotification,
		arg.ID,
		arg.UserID,
		arg.Type,
		arg.Message,
	)
	return err
}

const createPost = `-- name: CreatePost :exec
INSERT INTO posts (
    id, user_id, content, media_urls, visibility, 
    original_post_id, reply_to_id, root_post_id, is_repost, is_reply, created_at
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type CreatePostParams struct {
	ID             string          `json:"id"`
	UserID         sql.NullString  `json:"user_id"`
	Content        sql.NullString  `json:"content"`
	MediaUrls      json.RawMessage `json:"media_urls"`
	Visibility     sql.NullString  `json:"visibility"`
	OriginalPostID sql.NullString  `json:"original_post_id"`
	ReplyToID      sql.NullString  `json:"reply_to_id"`
	RootPostID     sql.NullString  `json:"root_post_id"`
	IsRepost       sql.NullBool    `json:"is_repost"`
	IsReply        sql.NullBool    `json:"is_reply"`
	CreatedAt      sql.NullTime    `json:"created_at"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) error {
	_, err := q.db.ExecContext(ctx, createPost,
		arg.ID,
		arg.UserID,
		arg.Content,
		arg.MediaUrls,
		arg.Visibility,
		arg.OriginalPostID,
		arg.ReplyToID,
		arg.RootPostID,
		arg.IsRepost,
		arg.IsReply,
		arg.CreatedAt,
	)
	return err
}

const createRepost = `-- name: CreateRepost :exec
INSERT INTO reposts (id, user_id, original_post_id, is_quote_repost, additional_comment)
VALUES (?, ?, ?, ?, ?)
`

type CreateRepostParams struct {
	ID                string         `json:"id"`
	UserID            sql.NullString `json:"user_id"`
	OriginalPostID    sql.NullString `json:"original_post_id"`
	IsQuoteRepost     sql.NullBool   `json:"is_quote_repost"`
	AdditionalComment sql.NullString `json:"additional_comment"`
}

func (q *Queries) CreateRepost(ctx context.Context, arg CreateRepostParams) error {
	_, err := q.db.ExecContext(ctx, createRepost,
		arg.ID,
		arg.UserID,
		arg.OriginalPostID,
		arg.IsQuoteRepost,
		arg.AdditionalComment,
	)
	return err
}

const createUser = `-- name: CreateUser :execresult
INSERT INTO users (id, firebase_uid, email, password_hash, username, display_name)
VALUES (?, ?, ?, ?, ?, ?)
`

type CreateUserParams struct {
	ID           string         `json:"id"`
	FirebaseUid  string         `json:"firebase_uid"`
	Email        string         `json:"email"`
	PasswordHash string         `json:"password_hash"`
	Username     string         `json:"username"`
	DisplayName  sql.NullString `json:"display_name"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser,
		arg.ID,
		arg.FirebaseUid,
		arg.Email,
		arg.PasswordHash,
		arg.Username,
		arg.DisplayName,
	)
}

const deleteResetToken = `-- name: DeleteResetToken :exec
DELETE FROM password_reset_tokens
WHERE token = ?
`

// 使用済みのリセットトークンを削除するクエリ
func (q *Queries) DeleteResetToken(ctx context.Context, token string) error {
	_, err := q.db.ExecContext(ctx, deleteResetToken, token)
	return err
}

const deleteUser = `-- name: DeleteUser :execresult
DELETE FROM users WHERE id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id string) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteUser, id)
}

const getDMConversation = `-- name: GetDMConversation :many
SELECT id, sender_id, receiver_id, content, created_at
FROM dms
WHERE (sender_id = ? AND receiver_id = ?)
   OR (sender_id = ? AND receiver_id = ?)
ORDER BY createdAt ASC
`

type GetDMConversationParams struct {
	SenderID     sql.NullString `json:"sender_id"`
	ReceiverID   sql.NullString `json:"receiver_id"`
	SenderID_2   sql.NullString `json:"sender_id_2"`
	ReceiverID_2 sql.NullString `json:"receiver_id_2"`
}

func (q *Queries) GetDMConversation(ctx context.Context, arg GetDMConversationParams) ([]Dm, error) {
	rows, err := q.db.QueryContext(ctx, getDMConversation,
		arg.SenderID,
		arg.ReceiverID,
		arg.SenderID_2,
		arg.ReceiverID_2,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Dm
	for rows.Next() {
		var i Dm
		if err := rows.Scan(
			&i.ID,
			&i.SenderID,
			&i.ReceiverID,
			&i.Content,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEmailFromUsername = `-- name: GetEmailFromUsername :one
SELECT email
FROM users
WHERE username = ?
`

func (q *Queries) GetEmailFromUsername(ctx context.Context, username string) (string, error) {
	row := q.db.QueryRowContext(ctx, getEmailFromUsername, username)
	var email string
	err := row.Scan(&email)
	return email, err
}

const getFollowStatus = `-- name: GetFollowStatus :one
SELECT EXISTS(
    SELECT 1
    FROM follows
    WHERE follower_id = ? AND following_id = ?
) AS following
`

type GetFollowStatusParams struct {
	FollowerID  sql.NullString `json:"follower_id"`
	FollowingID sql.NullString `json:"following_id"`
}

func (q *Queries) GetFollowStatus(ctx context.Context, arg GetFollowStatusParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, getFollowStatus, arg.FollowerID, arg.FollowingID)
	var following bool
	err := row.Scan(&following)
	return following, err
}

const getIdfromFirebaseUID = `-- name: GetIdfromFirebaseUID :one
SELECT id FROM users WHERE firebase_uid = ?
`

func (q *Queries) GetIdfromFirebaseUID(ctx context.Context, firebaseUid string) (string, error) {
	row := q.db.QueryRowContext(ctx, getIdfromFirebaseUID, firebaseUid)
	var id string
	err := row.Scan(&id)
	return id, err
}

const getRecentPosts = `-- name: GetRecentPosts :many
SELECT p.id, p.user_id, p.content, p.created_at, p.updated_at, p.is_repost, p.original_post_id, p.reply_to_id, p.root_post_id, p.is_reply, p.media_urls, p.likes_count, p.reposts_count, p.replies_count, p.views_count, p.visibility, p.is_pinned, p.is_deleted, u.username, u.display_name
FROM posts p
JOIN users u ON p.user_id = u.id
WHERE p.is_deleted = FALSE
ORDER BY p.created_at DESC
LIMIT ?
`

type GetRecentPostsRow struct {
	ID             string          `json:"id"`
	UserID         sql.NullString  `json:"user_id"`
	Content        sql.NullString  `json:"content"`
	CreatedAt      sql.NullTime    `json:"created_at"`
	UpdatedAt      sql.NullTime    `json:"updated_at"`
	IsRepost       sql.NullBool    `json:"is_repost"`
	OriginalPostID sql.NullString  `json:"original_post_id"`
	ReplyToID      sql.NullString  `json:"reply_to_id"`
	RootPostID     sql.NullString  `json:"root_post_id"`
	IsReply        sql.NullBool    `json:"is_reply"`
	MediaUrls      json.RawMessage `json:"media_urls"`
	LikesCount     sql.NullInt32   `json:"likes_count"`
	RepostsCount   sql.NullInt32   `json:"reposts_count"`
	RepliesCount   sql.NullInt32   `json:"replies_count"`
	ViewsCount     sql.NullInt32   `json:"views_count"`
	Visibility     sql.NullString  `json:"visibility"`
	IsPinned       sql.NullBool    `json:"is_pinned"`
	IsDeleted      sql.NullBool    `json:"is_deleted"`
	Username       string          `json:"username"`
	DisplayName    sql.NullString  `json:"display_name"`
}

func (q *Queries) GetRecentPosts(ctx context.Context, limit int32) ([]GetRecentPostsRow, error) {
	rows, err := q.db.QueryContext(ctx, getRecentPosts, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRecentPostsRow
	for rows.Next() {
		var i GetRecentPostsRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.IsRepost,
			&i.OriginalPostID,
			&i.ReplyToID,
			&i.RootPostID,
			&i.IsReply,
			&i.MediaUrls,
			&i.LikesCount,
			&i.RepostsCount,
			&i.RepliesCount,
			&i.ViewsCount,
			&i.Visibility,
			&i.IsPinned,
			&i.IsDeleted,
			&i.Username,
			&i.DisplayName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUnreadNotifications = `-- name: GetUnreadNotifications :many
SELECT id, user_id, type, message, created_at, is_read
FROM notifications
WHERE user_id = ? AND is_read = FALSE
ORDER BY created_at DESC
`

func (q *Queries) GetUnreadNotifications(ctx context.Context, userID sql.NullString) ([]Notification, error) {
	rows, err := q.db.QueryContext(ctx, getUnreadNotifications, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Notification
	for rows.Next() {
		var i Notification
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Type,
			&i.Message,
			&i.CreatedAt,
			&i.IsRead,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, password_hash FROM users WHERE username = ?
`

type GetUserByEmailRow struct {
	ID           string `json:"id"`
	PasswordHash string `json:"password_hash"`
}

func (q *Queries) GetUserByEmail(ctx context.Context, username string) (GetUserByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, username)
	var i GetUserByEmailRow
	err := row.Scan(&i.ID, &i.PasswordHash)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, firebase_uid, email, username, display_name, bio, location, followers_count, following_count, posts_count
FROM users
WHERE id = ?
`

type GetUserByIdRow struct {
	ID             string         `json:"id"`
	FirebaseUid    string         `json:"firebase_uid"`
	Email          string         `json:"email"`
	Username       string         `json:"username"`
	DisplayName    sql.NullString `json:"display_name"`
	Bio            sql.NullString `json:"bio"`
	Location       sql.NullString `json:"location"`
	FollowersCount sql.NullInt32  `json:"followers_count"`
	FollowingCount sql.NullInt32  `json:"following_count"`
	PostsCount     sql.NullInt32  `json:"posts_count"`
}

func (q *Queries) GetUserById(ctx context.Context, id string) (GetUserByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i GetUserByIdRow
	err := row.Scan(
		&i.ID,
		&i.FirebaseUid,
		&i.Email,
		&i.Username,
		&i.DisplayName,
		&i.Bio,
		&i.Location,
		&i.FollowersCount,
		&i.FollowingCount,
		&i.PostsCount,
	)
	return i, err
}

const getUserStats = `-- name: GetUserStats :one
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
GROUP BY u.id
`

type GetUserStatsRow struct {
	ID                 string        `json:"id"`
	Username           string        `json:"username"`
	FollowersCount     sql.NullInt32 `json:"followers_count"`
	FollowingCount     sql.NullInt32 `json:"following_count"`
	PostsCount         sql.NullInt32 `json:"posts_count"`
	TotalLikesReceived int64         `json:"total_likes_received"`
}

func (q *Queries) GetUserStats(ctx context.Context, id string) (GetUserStatsRow, error) {
	row := q.db.QueryRowContext(ctx, getUserStats, id)
	var i GetUserStatsRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FollowersCount,
		&i.FollowingCount,
		&i.PostsCount,
		&i.TotalLikesReceived,
	)
	return i, err
}

const getUserTimeline = `-- name: GetUserTimeline :many
SELECT p.id, p.user_id, p.content, p.created_at, p.updated_at, p.is_repost, p.original_post_id, p.reply_to_id, p.root_post_id, p.is_reply, p.media_urls, p.likes_count, p.reposts_count, p.replies_count, p.views_count, p.visibility, p.is_pinned, p.is_deleted, u.username, u.display_name
FROM posts p
JOIN users u ON p.user_id = u.id
WHERE p.user_id IN (
    SELECT following_id
    FROM follows
    WHERE follower_id = ?
) OR p.user_id = ?
ORDER BY p.created_at DESC
LIMIT ?
`

type GetUserTimelineParams struct {
	FollowerID sql.NullString `json:"follower_id"`
	UserID     sql.NullString `json:"user_id"`
	Limit      int32          `json:"limit"`
}

type GetUserTimelineRow struct {
	ID             string          `json:"id"`
	UserID         sql.NullString  `json:"user_id"`
	Content        sql.NullString  `json:"content"`
	CreatedAt      sql.NullTime    `json:"created_at"`
	UpdatedAt      sql.NullTime    `json:"updated_at"`
	IsRepost       sql.NullBool    `json:"is_repost"`
	OriginalPostID sql.NullString  `json:"original_post_id"`
	ReplyToID      sql.NullString  `json:"reply_to_id"`
	RootPostID     sql.NullString  `json:"root_post_id"`
	IsReply        sql.NullBool    `json:"is_reply"`
	MediaUrls      json.RawMessage `json:"media_urls"`
	LikesCount     sql.NullInt32   `json:"likes_count"`
	RepostsCount   sql.NullInt32   `json:"reposts_count"`
	RepliesCount   sql.NullInt32   `json:"replies_count"`
	ViewsCount     sql.NullInt32   `json:"views_count"`
	Visibility     sql.NullString  `json:"visibility"`
	IsPinned       sql.NullBool    `json:"is_pinned"`
	IsDeleted      sql.NullBool    `json:"is_deleted"`
	Username       string          `json:"username"`
	DisplayName    sql.NullString  `json:"display_name"`
}

func (q *Queries) GetUserTimeline(ctx context.Context, arg GetUserTimelineParams) ([]GetUserTimelineRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserTimeline, arg.FollowerID, arg.UserID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserTimelineRow
	for rows.Next() {
		var i GetUserTimelineRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.IsRepost,
			&i.OriginalPostID,
			&i.ReplyToID,
			&i.RootPostID,
			&i.IsReply,
			&i.MediaUrls,
			&i.LikesCount,
			&i.RepostsCount,
			&i.RepliesCount,
			&i.ViewsCount,
			&i.Visibility,
			&i.IsPinned,
			&i.IsDeleted,
			&i.Username,
			&i.DisplayName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeFollow = `-- name: RemoveFollow :exec
DELETE FROM follows
WHERE follower_id = ? AND following_id = ?
`

type RemoveFollowParams struct {
	FollowerID  sql.NullString `json:"follower_id"`
	FollowingID sql.NullString `json:"following_id"`
}

func (q *Queries) RemoveFollow(ctx context.Context, arg RemoveFollowParams) error {
	_, err := q.db.ExecContext(ctx, removeFollow, arg.FollowerID, arg.FollowingID)
	return err
}

const saveResetToken = `-- name: SaveResetToken :exec

INSERT INTO password_reset_tokens (email, token, expiry)
VALUES (?, ?, ?)
`

type SaveResetTokenParams struct {
	Email  string    `json:"email"`
	Token  string    `json:"token"`
	Expiry time.Time `json:"expiry"`
}

// ここから自作
// パスワードリセット用のトークンを保存するクエリ
// params: email, token, expiry
func (q *Queries) SaveResetToken(ctx context.Context, arg SaveResetTokenParams) error {
	_, err := q.db.ExecContext(ctx, saveResetToken, arg.Email, arg.Token, arg.Expiry)
	return err
}

const searchPostsByHashtag = `-- name: SearchPostsByHashtag :many
SELECT p.id, p.user_id, p.content, p.created_at, p.updated_at, p.is_repost, p.original_post_id, p.reply_to_id, p.root_post_id, p.is_reply, p.media_urls, p.likes_count, p.reposts_count, p.replies_count, p.views_count, p.visibility, p.is_pinned, p.is_deleted, u.username, u.display_name
FROM posts p
JOIN users u ON p.user_id = u.id
WHERE p.content LIKE ?
ORDER BY p.created_at DESC
LIMIT ?
`

type SearchPostsByHashtagParams struct {
	Content sql.NullString `json:"content"`
	Limit   int32          `json:"limit"`
}

type SearchPostsByHashtagRow struct {
	ID             string          `json:"id"`
	UserID         sql.NullString  `json:"user_id"`
	Content        sql.NullString  `json:"content"`
	CreatedAt      sql.NullTime    `json:"created_at"`
	UpdatedAt      sql.NullTime    `json:"updated_at"`
	IsRepost       sql.NullBool    `json:"is_repost"`
	OriginalPostID sql.NullString  `json:"original_post_id"`
	ReplyToID      sql.NullString  `json:"reply_to_id"`
	RootPostID     sql.NullString  `json:"root_post_id"`
	IsReply        sql.NullBool    `json:"is_reply"`
	MediaUrls      json.RawMessage `json:"media_urls"`
	LikesCount     sql.NullInt32   `json:"likes_count"`
	RepostsCount   sql.NullInt32   `json:"reposts_count"`
	RepliesCount   sql.NullInt32   `json:"replies_count"`
	ViewsCount     sql.NullInt32   `json:"views_count"`
	Visibility     sql.NullString  `json:"visibility"`
	IsPinned       sql.NullBool    `json:"is_pinned"`
	IsDeleted      sql.NullBool    `json:"is_deleted"`
	Username       string          `json:"username"`
	DisplayName    sql.NullString  `json:"display_name"`
}

func (q *Queries) SearchPostsByHashtag(ctx context.Context, arg SearchPostsByHashtagParams) ([]SearchPostsByHashtagRow, error) {
	rows, err := q.db.QueryContext(ctx, searchPostsByHashtag, arg.Content, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SearchPostsByHashtagRow
	for rows.Next() {
		var i SearchPostsByHashtagRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Content,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.IsRepost,
			&i.OriginalPostID,
			&i.ReplyToID,
			&i.RootPostID,
			&i.IsReply,
			&i.MediaUrls,
			&i.LikesCount,
			&i.RepostsCount,
			&i.RepliesCount,
			&i.ViewsCount,
			&i.Visibility,
			&i.IsPinned,
			&i.IsDeleted,
			&i.Username,
			&i.DisplayName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const sendDM = `-- name: SendDM :exec
INSERT INTO dms (id, sender_id, receiver_id, content)
VALUES (?, ?, ?, ?)
`

type SendDMParams struct {
	ID         string         `json:"id"`
	SenderID   sql.NullString `json:"sender_id"`
	ReceiverID sql.NullString `json:"receiver_id"`
	Content    sql.NullString `json:"content"`
}

func (q *Queries) SendDM(ctx context.Context, arg SendDMParams) error {
	_, err := q.db.ExecContext(ctx, sendDM,
		arg.ID,
		arg.SenderID,
		arg.ReceiverID,
		arg.Content,
	)
	return err
}

const updateFollowersCount = `-- name: UpdateFollowersCount :exec
UPDATE users
SET followers_count = (
    SELECT COUNT(*) FROM follows WHERE following_id = users.id
)
WHERE users.id = ?
`

func (q *Queries) UpdateFollowersCount(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, updateFollowersCount, id)
	return err
}

const updatePasswordByEmail = `-- name: UpdatePasswordByEmail :exec
UPDATE users
SET password_hash = ?, last_password_change = NOW()
WHERE email = ?
`

type UpdatePasswordByEmailParams struct {
	PasswordHash string `json:"password_hash"`
	Email        string `json:"email"`
}

// パスワードを更新するクエリ
func (q *Queries) UpdatePasswordByEmail(ctx context.Context, arg UpdatePasswordByEmailParams) error {
	_, err := q.db.ExecContext(ctx, updatePasswordByEmail, arg.PasswordHash, arg.Email)
	return err
}

const updatePostLikesCount = `-- name: UpdatePostLikesCount :exec
UPDATE posts
SET likes_count = (
    SELECT COUNT(*) FROM likes WHERE post_id = posts.id
)
WHERE posts.id = ?
`

func (q *Queries) UpdatePostLikesCount(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, updatePostLikesCount, id)
	return err
}

const updateUserBanStatus = `-- name: UpdateUserBanStatus :exec
UPDATE users
SET
    is_banned = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
`

type UpdateUserBanStatusParams struct {
	IsBanned sql.NullBool `json:"is_banned"`
	ID       string       `json:"id"`
}

func (q *Queries) UpdateUserBanStatus(ctx context.Context, arg UpdateUserBanStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateUserBanStatus, arg.IsBanned, arg.ID)
	return err
}

const updateUserEmail = `-- name: UpdateUserEmail :exec
UPDATE users
SET
    email = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
`

type UpdateUserEmailParams struct {
	Email string `json:"email"`
	ID    string `json:"id"`
}

func (q *Queries) UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) error {
	_, err := q.db.ExecContext(ctx, updateUserEmail, arg.Email, arg.ID)
	return err
}

const updateUserInfo = `-- name: UpdateUserInfo :exec
UPDATE users
SET bio = ?, location = ?
WHERE id = ?
`

type UpdateUserInfoParams struct {
	Bio      sql.NullString `json:"bio"`
	Location sql.NullString `json:"location"`
	ID       string         `json:"id"`
}

func (q *Queries) UpdateUserInfo(ctx context.Context, arg UpdateUserInfoParams) error {
	_, err := q.db.ExecContext(ctx, updateUserInfo, arg.Bio, arg.Location, arg.ID)
	return err
}

const updateUserName = `-- name: UpdateUserName :exec
UPDATE users
SET
    username = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
`

type UpdateUserNameParams struct {
	Username string `json:"username"`
	ID       string `json:"id"`
}

func (q *Queries) UpdateUserName(ctx context.Context, arg UpdateUserNameParams) error {
	_, err := q.db.ExecContext(ctx, updateUserName, arg.Username, arg.ID)
	return err
}

const updateUserNotifications = `-- name: UpdateUserNotifications :exec
UPDATE users
SET
    notification_settings = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
`

type UpdateUserNotificationsParams struct {
	NotificationSettings json.RawMessage `json:"notification_settings"`
	ID                   string          `json:"id"`
}

func (q *Queries) UpdateUserNotifications(ctx context.Context, arg UpdateUserNotificationsParams) error {
	_, err := q.db.ExecContext(ctx, updateUserNotifications, arg.NotificationSettings, arg.ID)
	return err
}

const updateUserPrivacy = `-- name: UpdateUserPrivacy :exec
UPDATE users
SET
    is_private = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
`

type UpdateUserPrivacyParams struct {
	IsPrivate sql.NullBool `json:"is_private"`
	ID        string       `json:"id"`
}

func (q *Queries) UpdateUserPrivacy(ctx context.Context, arg UpdateUserPrivacyParams) error {
	_, err := q.db.ExecContext(ctx, updateUserPrivacy, arg.IsPrivate, arg.ID)
	return err
}

const updateUserProfile = `-- name: UpdateUserProfile :exec
UPDATE users
SET 
    profile_image_url = COALESCE(?, profile_image_url),
    bio = COALESCE(?, bio),
    location = COALESCE(?, location),
    website = COALESCE(?, website),
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
`

type UpdateUserProfileParams struct {
	ProfileImageUrl sql.NullString `json:"profile_image_url"`
	Bio             sql.NullString `json:"bio"`
	Location        sql.NullString `json:"location"`
	Website         sql.NullString `json:"website"`
	ID              string         `json:"id"`
}

func (q *Queries) UpdateUserProfile(ctx context.Context, arg UpdateUserProfileParams) error {
	_, err := q.db.ExecContext(ctx, updateUserProfile,
		arg.ProfileImageUrl,
		arg.Bio,
		arg.Location,
		arg.Website,
		arg.ID,
	)
	return err
}

const updateUserSettings = `-- name: UpdateUserSettings :exec
UPDATE users
SET 
    display_name = COALESCE(?, display_name),
    birth_date = COALESCE(?, birth_date),
    language = COALESCE(?, language),
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
`

type UpdateUserSettingsParams struct {
	DisplayName sql.NullString `json:"display_name"`
	BirthDate   sql.NullTime   `json:"birth_date"`
	Language    sql.NullString `json:"language"`
	ID          string         `json:"id"`
}

func (q *Queries) UpdateUserSettings(ctx context.Context, arg UpdateUserSettingsParams) error {
	_, err := q.db.ExecContext(ctx, updateUserSettings,
		arg.DisplayName,
		arg.BirthDate,
		arg.Language,
		arg.ID,
	)
	return err
}

const validateResetToken = `-- name: ValidateResetToken :one
SELECT email FROM password_reset_tokens
WHERE token = ? AND expiry > NOW()
`

// トークンを検証して対応するメールを取得するクエリ
func (q *Queries) ValidateResetToken(ctx context.Context, token string) (string, error) {
	row := q.db.QueryRowContext(ctx, validateResetToken, token)
	var email string
	err := row.Scan(&email)
	return email, err
}
