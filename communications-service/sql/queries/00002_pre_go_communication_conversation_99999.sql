-- name: InsertCommunicationConversation :execresult
INSERT INTO pre_go_communication_conversation_99999 (
        title,
        user_ids,
        description,
        type,
        background,
        emoji,
        is_deleted,
        created_at,
        updated_at
    )
VALUES (
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        NOW(),
        NOW()
    );
-- name: GetCommunicationConversationById :one
SELECT id,
    title,
    user_ids,
    description,
    type,
    background,
    emoji,
    is_deleted,
    created_at,
    updated_at
FROM pre_go_communication_conversation_99999
WHERE id = ?;
-- name: GetCommunicationConversationsByUserId :many
SELECT id,
    title,
    user_ids,
    description,
    type,
    background,
    emoji,
    is_deleted,
    created_at,
    updated_at
FROM pre_go_communication_conversation_99999
WHERE JSON_CONTAINS(user_ids, JSON_QUOTE(?), '$')
LIMIT ? OFFSET ?;
-- name: UpdateCommunicationConversation :execresult
UPDATE pre_go_communication_conversation_99999
SET title = ?,
    user_ids = JSON_ARRAY_APPEND(user_ids, '$', ?),
    description = ?,
    type = ?,
    background = ?,
    emoji = ?,
    is_deleted = ?,
    updated_at = NOW()
WHERE id = ?;
-- name: AddUserToCommunicationConversation :execresult
UPDATE pre_go_communication_conversation_99999
SET user_ids = JSON_ARRAY_APPEND(user_ids, '$', ?)
WHERE id = ?;
-- name: RemoveUserFromCommunicationConversation :execresult
UPDATE pre_go_communication_conversation_99999
SET user_ids = JSON_REMOVE(
        user_ids,
        JSON_UNQUOTE(JSON_SEARCH(user_ids, 'one', ?))
    )
WHERE id = ?;
-- name: DeleteCommunicationConversation :execresult
DELETE FROM pre_go_communication_conversation_99999
WHERE id = ?;