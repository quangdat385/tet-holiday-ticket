-- name: InsertCommunicationMedia :execresult
INSERT INTO pre_go_communication_media_99999 (
        message_id,
        conversation_id,
        user_id,
        file_name,
        original_name,
        mime_type,
        size,
        url,
        created_at,
        updated_at
    )
VALUES (?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW());
-- name: GetCommunicationMediaById :one
SELECT *
FROM pre_go_communication_media_99999
WHERE id = ?;
-- name: GetCommunicationMediaByMessageId :many
SELECT *
FROM pre_go_communication_media_99999
WHERE message_id = ?;
-- name: GetCommunicationMediaByConversationId :many
SELECT *
FROM pre_go_communication_media_99999
WHERE conversation_id = ?
LIMIT ? OFFSET ?;
-- name: DeleteCommunicationMedia :execresult
DELETE FROM pre_go_communication_media_99999
WHERE id = ?;
