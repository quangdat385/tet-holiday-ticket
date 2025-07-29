-- name: InsertCommunicationMessage :execresult
INSERT INTO pre_go_communication_message_99999 (
        conversation_id,
        user_id,
        status,
        message,
        type,
        created_at,
        updated_at
    )
VALUES (?, ?, ?, ?, ?, NOW(), NOW());
-- name: GetCommunicationMessageById :one
SELECT *
FROM pre_go_communication_message_99999
WHERE id = ?;
-- name: GetCommunicationMessagesByConversationId :many 
SELECT *
FROM pre_go_communication_message_99999
WHERE conversation_id = ?
LIMIT ? OFFSET ?;
-- name: GetCommunicationMessagesByUserId :many
SELECT *
FROM pre_go_communication_message_99999
WHERE user_id = ?
LIMIT ? OFFSET ?;
-- name: UpdateCommunicationMessage :execresult
UPDATE pre_go_communication_message_99999
SET status = ?,
    updated_at = NOW()
WHERE id = ?;
-- name: DeleteCommunicationMessage :execresult
DELETE FROM pre_go_communication_message_99999
WHERE id = ?;