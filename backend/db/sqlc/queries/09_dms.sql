
-- name: GetDMConversation :many
SELECT *
FROM dms
WHERE (sender_id = ? AND receiver_id = ?)
   OR (sender_id = ? AND receiver_id = ?)
ORDER BY createdAt ASC;

-- name: SendDM :exec
INSERT INTO dms (id, sender_id, receiver_id, content, created_at)
VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP);