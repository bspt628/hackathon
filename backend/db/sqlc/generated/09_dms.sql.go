// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: 09_dms.sql

package sqlc

import (
	"context"
	"database/sql"
)

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

const sendDM = `-- name: SendDM :exec
INSERT INTO dms (id, sender_id, receiver_id, content, created_at)
VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP)
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
