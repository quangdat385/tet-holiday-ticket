-- name: InsertCommunicationInfo :execresult
INSERT INTO pre_go_communication_info_99999 (
        user_id,
        status,
        value,
        type,
        created_at,
        updated_at
    )
VALUES (?, ?, ?, ?, NOW(), NOW());
-- name: GetCommunicationInfoByUserID :one
SELECT id,
    user_id,
    status,
    value,
    type,
    created_at,
    updated_at
FROM pre_go_communication_info_99999
WHERE id = ?;
-- name: GetCommunicationInfoByID :one
SELECT id,
    user_id,
    status,
    value,
    type,
    created_at,
    updated_at
FROM pre_go_communication_info_99999
WHERE id = ?;
-- name: UpdateCommunicationInfo :execresult
UPDATE pre_go_communication_info_99999
SET status = ?,
    value = ?,
    type = ?,
    updated_at = NOW()
WHERE id = ?;
-- name: UpdateCommunicationInfoByUserId :execresult
UPDATE pre_go_communication_info_99999
SET status = ?,
    value = ?,
    type = ?,
    updated_at = NOW()
WHERE user_id = ?;
-- name: DeleteCommunicationInfo :execresult
DELETE FROM pre_go_communication_info_99999
WHERE id = ?;