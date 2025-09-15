-- name: InsertOrder :execresult
INSERT INTO pre_go_order_052025_99999 (
        order_number,
        user_id,
        station_code,
        order_amount,
        terminal_id,
        order_date,
        order_notes,
        order_item,
        created_at,
        updated_at
    )
VALUES (?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW());
-- name: UpdateOrder :execresult
UPDATE pre_go_order_052025_99999
SET order_amount = ?,
    terminal_id = ?,
    station_code = ?,
    order_date = ?,
    order_notes = ?,
    updated_at = NOW()
WHERE id = ?;
-- name: UpdateOrderNote :execresult
UPDATE pre_go_order_052025_99999
SET order_notes = ?
WHERE order_number = ?;
-- name: DeleteOrder :execresult
DELETE FROM pre_go_order_052025_99999
WHERE id = ?;
-- name: GetOrderById :one
SELECT id,
    order_number,
    order_amount,
    terminal_id,
    station_code,
    user_id,
    order_date,
    order_notes,
    order_item,
    created_at,
    updated_at
FROM pre_go_order_052025_99999
WHERE id = ?;
-- name: GetOrderByOrderNumber :one
SELECT id,
    order_number,
    order_amount,
    terminal_id,
    user_id,
    station_code,
    order_date,
    order_notes,
    order_item,
    created_at,
    updated_at
FROM pre_go_order_052025_99999
WHERE order_number = ?;
-- name: GetOrdersByUserId :many
SELECT id,
    order_number,
    order_amount,
    terminal_id,
    station_code,
    user_id,
    order_date,
    order_notes,
    order_item,
    created_at,
    updated_at
FROM pre_go_order_052025_99999
WHERE user_id = ?
LIMIT ? OFFSET ?;
-- name: CheckIdempotencyOrder :many
SELECT id
FROM pre_go_order_052025_99999
WHERE user_id = ?
    AND order_date >= ?
    AND order_date <= ?;