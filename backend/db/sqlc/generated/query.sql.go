// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
	"encoding/json"
)

const addBlock = `-- name: AddBlock :exec
INSERT INTO blocks (id, blockedById, blockedUserId)
VALUES (?, ?, ?)
`

type AddBlockParams struct {
	ID            string         `json:"id"`
	Blockedbyid   sql.NullString `json:"blockedbyid"`
	Blockeduserid sql.NullString `json:"blockeduserid"`
}

func (q *Queries) AddBlock(ctx context.Context, arg AddBlockParams) error {
	_, err := q.db.ExecContext(ctx, addBlock, arg.ID, arg.Blockedbyid, arg.Blockeduserid)
	return err
}

const addFollow = `-- name: AddFollow :exec
INSERT INTO follows (id, followerId, followingId)
VALUES (?, ?, ?)
`

type AddFollowParams struct {
	ID          string         `json:"id"`
	Followerid  sql.NullString `json:"followerid"`
	Followingid sql.NullString `json:"followingid"`
}

func (q *Queries) AddFollow(ctx context.Context, arg AddFollowParams) error {
	_, err := q.db.ExecContext(ctx, addFollow, arg.ID, arg.Followerid, arg.Followingid)
	return err
}

const addLike = `-- name: AddLike :exec
INSERT INTO likes (id, userId, postId)
VALUES (?, ?, ?)
`

type AddLikeParams struct {
	ID     string         `json:"id"`
	Userid sql.NullString `json:"userid"`
	Postid sql.NullString `json:"postid"`
}

func (q *Queries) AddLike(ctx context.Context, arg AddLikeParams) error {
	_, err := q.db.ExecContext(ctx, addLike, arg.ID, arg.Userid, arg.Postid)
	return err
}

const createNotification = `-- name: CreateNotification :exec
INSERT INTO notifications (id, userId, type, message)
VALUES (?, ?, ?, ?)
`

type CreateNotificationParams struct {
	ID      string         `json:"id"`
	Userid  sql.NullString `json:"userid"`
	Type    sql.NullString `json:"type"`
	Message sql.NullString `json:"message"`
}

func (q *Queries) CreateNotification(ctx context.Context, arg CreateNotificationParams) error {
	_, err := q.db.ExecContext(ctx, createNotification,
		arg.ID,
		arg.Userid,
		arg.Type,
		arg.Message,
	)
	return err
}

const createPost = `-- name: CreatePost :execresult
INSERT INTO posts (id, user_id, content)
VALUES (?, ?, ?)
`

type CreatePostParams struct {
	ID      string         `json:"id"`
	UserID  sql.NullString `json:"user_id"`
	Content sql.NullString `json:"content"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createPost, arg.ID, arg.UserID, arg.Content)
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
INSERT INTO users (id, email, password_hash, username, display_name)
VALUES (?, ?, ?, ?, ?)
`

type CreateUserParams struct {
	ID           string         `json:"id"`
	Email        string         `json:"email"`
	PasswordHash string         `json:"password_hash"`
	Username     string         `json:"username"`
	DisplayName  sql.NullString `json:"display_name"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createUser,
		arg.ID,
		arg.Email,
		arg.PasswordHash,
		arg.Username,
		arg.DisplayName,
	)
}

const getDMConversation = `-- name: GetDMConversation :many
SELECT id, senderid, receiverid, content, createdat
FROM dms
WHERE (senderId = ? AND receiverId = ?)
   OR (senderId = ? AND receiverId = ?)
ORDER BY createdAt ASC
`

type GetDMConversationParams struct {
	Senderid     sql.NullString `json:"senderid"`
	Receiverid   sql.NullString `json:"receiverid"`
	Senderid_2   sql.NullString `json:"senderid_2"`
	Receiverid_2 sql.NullString `json:"receiverid_2"`
}

func (q *Queries) GetDMConversation(ctx context.Context, arg GetDMConversationParams) ([]Dm, error) {
	rows, err := q.db.QueryContext(ctx, getDMConversation,
		arg.Senderid,
		arg.Receiverid,
		arg.Senderid_2,
		arg.Receiverid_2,
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
			&i.Senderid,
			&i.Receiverid,
			&i.Content,
			&i.Createdat,
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

const getRecentPosts = `-- name: GetRecentPosts :many
SELECT p.id, p.user_id, p.content, p.created_at, p.updated_at, p.is_repost, p.original_post_id, p.reply_to_id, p.root_post_id, p.is_reply, p.media_urls, p.likes_count, p.reposts_count, p.replies_count, p.views_count, p.visibility, p.language, p.location, p.device, p.is_pinned, p.is_deleted, u.username, u.display_name
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
	Language       sql.NullString  `json:"language"`
	Location       sql.NullString  `json:"location"`
	Device         sql.NullString  `json:"device"`
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
			&i.Language,
			&i.Location,
			&i.Device,
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
SELECT id, userid, type, message, createdat, isread
FROM notifications
WHERE userId = ? AND isRead = FALSE
ORDER BY createdAt DESC
`

func (q *Queries) GetUnreadNotifications(ctx context.Context, userid sql.NullString) ([]Notification, error) {
	rows, err := q.db.QueryContext(ctx, getUnreadNotifications, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Notification
	for rows.Next() {
		var i Notification
		if err := rows.Scan(
			&i.ID,
			&i.Userid,
			&i.Type,
			&i.Message,
			&i.Createdat,
			&i.Isread,
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

const getUserById = `-- name: GetUserById :one
SELECT id, email, username, display_name, bio, location, followers_count, following_count, posts_count
FROM users
WHERE id = ?
`

type GetUserByIdRow struct {
	ID             string         `json:"id"`
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
SELECT p.id, p.user_id, p.content, p.created_at, p.updated_at, p.is_repost, p.original_post_id, p.reply_to_id, p.root_post_id, p.is_reply, p.media_urls, p.likes_count, p.reposts_count, p.replies_count, p.views_count, p.visibility, p.language, p.location, p.device, p.is_pinned, p.is_deleted, u.username, u.display_name
FROM posts p
JOIN users u ON p.user_id = u.id
WHERE p.user_id IN (
    SELECT followingId
    FROM follows
    WHERE followerId = ?
) OR p.user_id = ?
ORDER BY p.created_at DESC
LIMIT ?
`

type GetUserTimelineParams struct {
	Followerid sql.NullString `json:"followerid"`
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
	Language       sql.NullString  `json:"language"`
	Location       sql.NullString  `json:"location"`
	Device         sql.NullString  `json:"device"`
	IsPinned       sql.NullBool    `json:"is_pinned"`
	IsDeleted      sql.NullBool    `json:"is_deleted"`
	Username       string          `json:"username"`
	DisplayName    sql.NullString  `json:"display_name"`
}

func (q *Queries) GetUserTimeline(ctx context.Context, arg GetUserTimelineParams) ([]GetUserTimelineRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserTimeline, arg.Followerid, arg.UserID, arg.Limit)
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
			&i.Language,
			&i.Location,
			&i.Device,
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

const searchPostsByHashtag = `-- name: SearchPostsByHashtag :many
SELECT p.id, p.user_id, p.content, p.created_at, p.updated_at, p.is_repost, p.original_post_id, p.reply_to_id, p.root_post_id, p.is_reply, p.media_urls, p.likes_count, p.reposts_count, p.replies_count, p.views_count, p.visibility, p.language, p.location, p.device, p.is_pinned, p.is_deleted, u.username, u.display_name
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
	Language       sql.NullString  `json:"language"`
	Location       sql.NullString  `json:"location"`
	Device         sql.NullString  `json:"device"`
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
			&i.Language,
			&i.Location,
			&i.Device,
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
INSERT INTO dms (id, senderId, receiverId, content)
VALUES (?, ?, ?, ?)
`

type SendDMParams struct {
	ID         string         `json:"id"`
	Senderid   sql.NullString `json:"senderid"`
	Receiverid sql.NullString `json:"receiverid"`
	Content    sql.NullString `json:"content"`
}

func (q *Queries) SendDM(ctx context.Context, arg SendDMParams) error {
	_, err := q.db.ExecContext(ctx, sendDM,
		arg.ID,
		arg.Senderid,
		arg.Receiverid,
		arg.Content,
	)
	return err
}

const updateFollowersCount = `-- name: UpdateFollowersCount :exec
UPDATE users
SET followers_count = (
    SELECT COUNT(*) FROM follows WHERE followingId = users.id
)
WHERE users.id = ?
`

func (q *Queries) UpdateFollowersCount(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, updateFollowersCount, id)
	return err
}

const updatePostLikesCount = `-- name: UpdatePostLikesCount :exec
UPDATE posts
SET likes_count = (
    SELECT COUNT(*) FROM likes WHERE postId = posts.id
)
WHERE posts.id = ?
`

func (q *Queries) UpdatePostLikesCount(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, updatePostLikesCount, id)
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
