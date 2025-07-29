-- name: InsertTicketItem :execresult
INSERT INTO pre_go_ticket_item_99999 (
        name,
        description,
        ticket_id,
        train_id,
        seat_class,
        stock_initial,
        stock_available,
        departure_time,
        is_stock_prepared,
        price_original,
        price_flash,
        sale_start_time,
        sale_end_time,
        status,
        activity_id,
        updated_at,
        created_at
    )
VALUES (
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
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
-- name: UpdateTicketItem :execresult
UPDATE pre_go_ticket_item_99999
SET name = ?,
    ticket_id = ?,
    train_id = ?,
    description = ?,
    departure_time = ?,
    seat_class = ?,
    stock_initial = ?,
    stock_available = ?,
    is_stock_prepared = ?,
    price_original = ?,
    price_flash = ?,
    sale_start_time = ?,
    sale_end_time = ?,
    status = ?,
    activity_id = ?,
    updated_at = NOW()
WHERE id = ?;
-- name: DeleteTicketItem :execresult
DELETE FROM pre_go_ticket_item_99999
WHERE id = ?;
-- name: GetTicketItemById :one
SELECT id,
    name,
    ticket_id,
    train_id,
    seat_class,
    description,
    departure_time,
    stock_initial,
    stock_available,
    is_stock_prepared,
    price_original,
    price_flash,
    sale_start_time,
    sale_end_time,
    status,
    activity_id,
    updated_at,
    created_at
FROM pre_go_ticket_item_99999
WHERE id = ?;
-- name: UpdateTicketItemStock :execresult
UPDATE pre_go_ticket_item_99999
SET stock_available = stock_available - ?,
    updated_at = NOW()
WHERE id = ?
    AND stock_available = ?;