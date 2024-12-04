// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: 03_reposts.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createRepost = `-- name: CreateRepost :exec
INSERT INTO reposts (id, user_id, original_post_id, is_quote_repost, additional_comment, reposted_at)
VALUES (?, ?, ?, ?, ?, CURRENT_TIMESTAMP)
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

const getPostReposts = `-- name: GetPostReposts :many
SELECT u.id, u.username, u.display_name, r.is_quote_repost, r.additional_comment
FROM reposts r
JOIN users u ON r.user_id = u.id
WHERE r.original_post_id = ?
`

type GetPostRepostsRow struct {
	ID                string         `json:"id"`
	Username          string         `json:"username"`
	DisplayName       sql.NullString `json:"display_name"`
	IsQuoteRepost     sql.NullBool   `json:"is_quote_repost"`
	AdditionalComment sql.NullString `json:"additional_comment"`
}

func (q *Queries) GetPostReposts(ctx context.Context, originalPostID sql.NullString) ([]GetPostRepostsRow, error) {
	rows, err := q.db.QueryContext(ctx, getPostReposts, originalPostID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPostRepostsRow
	for rows.Next() {
		var i GetPostRepostsRow
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.DisplayName,
			&i.IsQuoteRepost,
			&i.AdditionalComment,
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