-- name: SaveResetToken :exec
INSERT INTO password_reset_tokens (email, token, expiry, created_at)
VALUES (?, ?, ?, CURRENT_TIMESTAMP);


-- name: ValidateResetToken :one
SELECT email FROM password_reset_tokens
WHERE token = ? AND expiry > NOW();

-- name: UpdatePassword :exec
UPDATE users
SET 
    password_hash = ?, 
    last_password_change = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE email = ?;

-- name: DeleteResetToken :exec
DELETE FROM password_reset_tokens
WHERE token = ?;
