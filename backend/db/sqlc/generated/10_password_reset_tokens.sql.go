// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: 10_password_reset_tokens.sql

package sqlc

import (
	"context"
	"time"
)

const deleteResetToken = `-- name: DeleteResetToken :exec
DELETE FROM password_reset_tokens
WHERE token = ?
`

func (q *Queries) DeleteResetToken(ctx context.Context, token string) error {
	_, err := q.db.ExecContext(ctx, deleteResetToken, token)
	return err
}

const saveResetToken = `-- name: SaveResetToken :exec
INSERT INTO password_reset_tokens (email, token, expiry, created_at)
VALUES (?, ?, ?, CURRENT_TIMESTAMP)
`

type SaveResetTokenParams struct {
	Email  string    `json:"email"`
	Token  string    `json:"token"`
	Expiry time.Time `json:"expiry"`
}

func (q *Queries) SaveResetToken(ctx context.Context, arg SaveResetTokenParams) error {
	_, err := q.db.ExecContext(ctx, saveResetToken, arg.Email, arg.Token, arg.Expiry)
	return err
}

const updatePassword = `-- name: UpdatePassword :exec
UPDATE users
SET 
    password_hash = ?, 
    last_password_change = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE email = ?
`

type UpdatePasswordParams struct {
	PasswordHash string `json:"password_hash"`
	Email        string `json:"email"`
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) error {
	_, err := q.db.ExecContext(ctx, updatePassword, arg.PasswordHash, arg.Email)
	return err
}

const validateResetToken = `-- name: ValidateResetToken :one
SELECT email FROM password_reset_tokens
WHERE token = ? AND expiry > NOW()
`

func (q *Queries) ValidateResetToken(ctx context.Context, token string) (string, error) {
	row := q.db.QueryRowContext(ctx, validateResetToken, token)
	var email string
	err := row.Scan(&email)
	return email, err
}
