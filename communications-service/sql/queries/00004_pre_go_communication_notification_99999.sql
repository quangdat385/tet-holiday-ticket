-- name: InsertNotification :execresult
INSERT INTO pre_go_communication_notification_99999 (
        `from`,
        `to`,
        content,
        created_at,
        updated_at
    )
VALUES (?, ?, ?, NOW(), NOW());
-- name: GetNotificationById :one
SELECT *
FROM pre_go_communication_notification_99999
WHERE id = ?;
-- name: GetNotificationsByUserIDTo :many
SELECT *
FROM pre_go_communication_notification_99999
WHERE JSON_CONTAINS(`to`, JSON_QUOTE(?), '$')
LIMIT ? OFFSET ?;
-- name: GetNotificationsByUserIDFrom :one
SELECT *
FROM pre_go_communication_notification_99999
WHERE `from` = ?;
-- name: GetNotificationWhenToIsNull :many
SELECT *
FROM pre_go_communication_notification_99999
WHERE `to` IS NULL
LIMIT ? OFFSET ?;
-- name: DeleteNotificationById :execresult
DELETE FROM pre_go_communication_notification_99999
WHERE id = ?;