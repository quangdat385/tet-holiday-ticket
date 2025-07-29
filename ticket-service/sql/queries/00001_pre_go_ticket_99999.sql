-- name: InsertTicket :execresult
INSERT INTO pre_go_ticket_99999 (
        id,
        name,
        description,
        start_time,
        end_time,
        status,
        updated_at,
        created_at
    )
VALUES (?, ?, ?, ?, ?, ?, NOW(), NOW());
-- name: UpdateTicket :exec
UPDATE pre_go_ticket_99999
SET name = ?,
    description = ?,
    start_time = ?,
    end_time = ?,
    status = ?,
    updated_at = NOW()
WHERE id = ?;
-- name: DeleteTicket :execresult
DELETE FROM pre_go_ticket_99999
WHERE id = ?;
-- name: GetTicketById :one
SELECT id,
    name,
    description,
    start_time,
    end_time,
    status,
    updated_at,
    created_at
FROM pre_go_ticket_99999
WHERE id = ?;
-- name: GetAllTickets :many
SELECT id,
    name,
    description,
    start_time,
    end_time,
    status,
    updated_at,
    created_at
FROM pre_go_ticket_99999
WHERE status = ?
LIMIT ? OFFSET ?;