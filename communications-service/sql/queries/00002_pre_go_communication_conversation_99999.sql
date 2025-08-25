-- name: InsertCommunicationConversation :execresult
INSERT INTO pre_go_communication_conversation_99999 (
        title,
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
        NOW(),
        NOW()
    );
-- name: GetListCommunicationConversations :many
SELECT c.id,
    c.title,
    JSON_ARRAYAGG(
        JSON_OBJECT(
            'user_id',
            u.user_id,
            'nick_name',
            u.nick_name,
            'is_deleted',
            u.is_deleted,
            'last_message',
            u.last_message,
            'updated_at',
            u.updated_at
        )
    ) AS user_ids,
    c.description,
    c.type,
    c.background,
    c.emoji,
    c.is_deleted,
    c.created_at,
    c.updated_at
FROM pre_go_communication_conversation_99999 AS c
    LEFT JOIN pre_go_communication_conversation_users_99999 AS u ON c.id = u.conversation_id
    AND u.is_deleted = false
WHERE c.id IN (
        SELECT conversation_id
        FROM pre_go_communication_conversation_users_99999 AS t
        WHERE t.user_id = ?
            AND t.is_deleted = false
    )
    AND c.is_deleted = false
GROUP BY c.id,
    c.title,
    c.description,
    c.type,
    c.background,
    c.emoji,
    c.is_deleted,
    c.created_at,
    c.updated_at
LIMIT ? OFFSET ?;
-- name: GetListCommunicationConversationsIsDeleted :many
SELECT c.id,
    c.title,
    JSON_ARRAYAGG(
        JSON_OBJECT(
            'user_id',
            u.user_id,
            'nick_name',
            u.nick_name,
            'is_deleted',
            u.is_deleted,
            'last_message',
            u.last_message,
            'updated_at',
            u.updated_at
        )
    ) AS user_ids,
    c.description,
    c.type,
    c.background,
    c.emoji,
    c.is_deleted,
    c.created_at,
    c.updated_at
FROM pre_go_communication_conversation_99999 AS c
    LEFT JOIN pre_go_communication_conversation_users_99999 AS u ON c.id = u.conversation_id
    AND u.is_deleted = false
WHERE c.id IN (
        select conversation_id
        from pre_go_communication_conversation_users_99999 AS t
        WHERE t.user_id = ?
            AND t.is_deleted = false
    )
    AND c.is_deleted = true
GROUP BY c.id,
    c.title,
    c.description,
    c.type,
    c.background,
    c.emoji,
    c.is_deleted,
    c.created_at,
    c.updated_at
LIMIT ? OFFSET ?;
-- name: GetCommunicationConversationById :one
SELECT c.id,
    c.title,
    JSON_ARRAYAGG(
        JSON_OBJECT(
            'user_id',
            u.user_id,
            'nick_name',
            u.nick_name,
            'is_deleted',
            u.is_deleted
        )
    ) AS user_ids,
    c.description,
    c.type,
    c.background,
    c.emoji,
    c.is_deleted,
    c.created_at,
    c.updated_at
FROM pre_go_communication_conversation_99999 AS c
    LEFT JOIN pre_go_communication_conversation_users_99999 AS u ON c.id = u.conversation_id
    AND u.is_deleted = false
WHERE c.id = ?
    AND c.is_deleted = false
GROUP BY c.id,
    c.title,
    c.description,
    c.type,
    c.background,
    c.emoji,
    c.is_deleted,
    c.created_at,
    c.updated_at;
-- name: GetCommunicationConversationByIDIsDeleted :one
SELECT c.id,
    c.title,
    JSON_ARRAYAGG(
        JSON_OBJECT(
            'user_id',
            u.user_id,
            'nick_name',
            u.nick_name,
            'is_deleted',
            u.is_deleted
        )
    ) AS user_ids,
    c.description,
    c.type,
    c.background,
    c.emoji,
    c.is_deleted,
    c.created_at,
    c.updated_at
FROM pre_go_communication_conversation_99999 AS c
    LEFT JOIN pre_go_communication_conversation_users_99999 AS u ON c.id = u.conversation_id
    AND u.is_deleted = false
WHERE c.id = ?
    AND c.is_deleted = true
GROUP BY c.id,
    c.title,
    c.description,
    c.type,
    c.background,
    c.emoji,
    c.is_deleted,
    c.created_at,
    c.updated_at;
-- name: UpdateCommunicationConversation :execresult
UPDATE pre_go_communication_conversation_99999
SET title = ?,
    description = ?,
    type = ?,
    background = ?,
    emoji = ?,
    is_deleted = ?,
    updated_at = NOW()
WHERE id = ?;
-- name: SoftDeleteCommunicationConversation :execresult
UPDATE pre_go_communication_conversation_99999
SET is_deleted = true,
    updated_at = NOW()
WHERE id = ?;
-- name: AddUserToCommunicationConversation :execresult
INSERT INTO pre_go_communication_conversation_users_99999 (
        conversation_id,
        user_id,
        nick_name
    )
VALUES (?, ?, ?);
-- name: RemoveUserFromCommunicationConversation :execresult
UPDATE pre_go_communication_conversation_users_99999
SET is_deleted = true,
    updated_at = NOW()
WHERE conversation_id = ?
    AND user_id = ?;
-- name: UpdateLastMessageInCommunicationConversation :execresult
UPDATE pre_go_communication_conversation_users_99999
SET last_message = ?,
    updated_at = NOW()
WHERE conversation_id = ?
    AND user_id = ?;
-- name: DeleteCommunicationConversation :execresult
DELETE FROM pre_go_communication_conversation_99999
WHERE id = ?;