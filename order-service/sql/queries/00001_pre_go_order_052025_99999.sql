-- name: InsertOrder :execresult
INSERT INTO pre_go_order_052025_99999 (
        order_number,
        order_amount,
        terminal_id,
        order_date,
        order_notes,
        created_at,
        updated_at
    )
VALUES (?, ?, ?, ?, ?, NOW(), NOW());
-- name: UpdateOrder :execresult
UPDATE pre_go_order_052025_99999
SET order_amount = ?,
    terminal_id = ?,
    order_date = ?,
    order_notes = ?,
    updated_at = NOW()
WHERE id = ?;
-- name: DeleteOrder :execresult
DELETE FROM pre_go_order_052025_99999
WHERE id = ?;
-- name: GetOrderById :one
SELECT id,
    order_number,
    order_amount,
    terminal_id,
    order_date,
    order_notes,
    created_at,
    updated_at
FROM pre_go_order_052025_99999
WHERE id = ?;