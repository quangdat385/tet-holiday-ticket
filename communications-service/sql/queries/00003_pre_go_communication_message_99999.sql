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
-- name: GetCommunicationMessagesByListOfConversationIds :many
SELECT m.*,
    MAX(m.created_at) AS last_message_time
FROM pre_go_communication_conversation_99999 c
    LEFT JOIN pre_go_communication_message_99999 m ON c.id = m.conversation_id
WHERE c.id = ?
ORDER BY last_message_time DESC
LIMIT ? OFFSET ?;
-- name: GetCommunicationMessagesByUserId :many
SELECT *
FROM pre_go_communication_message_99999
WHERE user_id = ?
LIMIT ? OFFSET ?;
-- name: UpdateCommunicationMessage :execresult
INSERT INTO pre_go_communication_message_read_99999 (message_id, user_id, read_at)
VALUES (?, ?, NOW());
-- name: DeleteCommunicationMessage :execresult
DELETE FROM pre_go_communication_message_99999
WHERE id = ?;