// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: 04_likes.sql

package sqlc

import (
	"context"
	"database/sql"
)

const addLike = `-- name: AddLike :execresult
INSERT INTO likes (id, user_id, post_id)
VALUES (?, ?, ?)
`

type AddLikeParams struct {
	ID     string         `json:"id"`
	UserID sql.NullString `json:"user_id"`
	PostID sql.NullString `json:"post_id"`
}

func (q *Queries) AddLike(ctx context.Context, arg AddLikeParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addLike, arg.ID, arg.UserID, arg.PostID)
}

const decrementLikesCount = `-- name: DecrementLikesCount :exec
UPDATE posts
SET likes_count = likes_count - 1
WHERE (id = ? AND likes_count > 0)
`

func (q *Queries) DecrementLikesCount(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, decrementLikesCount, id)
	return err
}

const getLikeID = `-- name: GetLikeID :one
SELECT id
FROM likes
WHERE user_id = ? AND post_id = ?
`

type GetLikeIDParams struct {
	UserID sql.NullString `json:"user_id"`
	PostID sql.NullString `json:"post_id"`
}

func (q *Queries) GetLikeID(ctx context.Context, arg GetLikeIDParams) (string, error) {
	row := q.db.QueryRowContext(ctx, getLikeID, arg.UserID, arg.PostID)
	var id string
	err := row.Scan(&id)
	return id, err
}

const getLikeStatus = `-- name: GetLikeStatus :one
SELECT EXISTS(
    SELECT 1
    FROM likes
    WHERE user_id = ? AND post_id = ?
) AS liked
`

type GetLikeStatusParams struct {
	UserID sql.NullString `json:"user_id"`
	PostID sql.NullString `json:"post_id"`
}

func (q *Queries) GetLikeStatus(ctx context.Context, arg GetLikeStatusParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, getLikeStatus, arg.UserID, arg.PostID)
	var liked bool
	err := row.Scan(&liked)
	return liked, err
}

const getLikes = `-- name: GetLikes :many
SELECT u.id, u.username, u.display_name
FROM likes l
JOIN users u ON l.user_id = u.id
WHERE l.post_id = ?
`

type GetLikesRow struct {
	ID          string         `json:"id"`
	Username    string         `json:"username"`
	DisplayName sql.NullString `json:"display_name"`
}

func (q *Queries) GetLikes(ctx context.Context, postID sql.NullString) ([]GetLikesRow, error) {
	rows, err := q.db.QueryContext(ctx, getLikes, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetLikesRow
	for rows.Next() {
		var i GetLikesRow
		if err := rows.Scan(&i.ID, &i.Username, &i.DisplayName); err != nil {
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

const getPostLikes = `-- name: GetPostLikes :many
SELECT u.id, u.username, u.display_name
FROM likes l
JOIN users u ON l.user_id = u.id
WHERE l.post_id = ?
`

type GetPostLikesRow struct {
	ID          string         `json:"id"`
	Username    string         `json:"username"`
	DisplayName sql.NullString `json:"display_name"`
}

func (q *Queries) GetPostLikes(ctx context.Context, postID sql.NullString) ([]GetPostLikesRow, error) {
	rows, err := q.db.QueryContext(ctx, getPostLikes, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPostLikesRow
	for rows.Next() {
		var i GetPostLikesRow
		if err := rows.Scan(&i.ID, &i.Username, &i.DisplayName); err != nil {
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

const getPostLikesCount = `-- name: GetPostLikesCount :one
SELECT likes_count FROM posts WHERE id = ?
`

func (q *Queries) GetPostLikesCount(ctx context.Context, id string) (sql.NullInt32, error) {
	row := q.db.QueryRowContext(ctx, getPostLikesCount, id)
	var likes_count sql.NullInt32
	err := row.Scan(&likes_count)
	return likes_count, err
}

const incrementLikesCount = `-- name: IncrementLikesCount :exec
UPDATE posts
SET likes_count = likes_count + 1
WHERE id = ?
`

func (q *Queries) IncrementLikesCount(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, incrementLikesCount, id)
	return err
}

const removeLike = `-- name: RemoveLike :execresult
DELETE FROM likes
WHERE user_id = ? AND post_id = ?
`

type RemoveLikeParams struct {
	UserID sql.NullString `json:"user_id"`
	PostID sql.NullString `json:"post_id"`
}

func (q *Queries) RemoveLike(ctx context.Context, arg RemoveLikeParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, removeLike, arg.UserID, arg.PostID)
}

const updateLikesCount = `-- name: UpdateLikesCount :execresult
UPDATE posts
SET likes_count = (
    SELECT COUNT(*) FROM likes WHERE post_id = posts.id
)
WHERE posts.id = ?
`

func (q *Queries) UpdateLikesCount(ctx context.Context, id string) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateLikesCount, id)
}
