-- name: InsertOrderDetail :execresult
INSERT INTO pre_go_ticket_order_detail_052025_99999 (
        ticket_item_id,
        order_number,
        passenger_name,
        departure_station,
        arrival_station,
        departure_time,
        passenger_id,
        seat_class,
        ticket_price,
        seat_number,
        created_at,
        updated_at
    )
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW());
-- name: UpdateOrderDetail :execresult
UPDATE pre_go_ticket_order_detail_052025_99999
SET ticket_item_id = ?,
    order_number = ?,
    passenger_name = ?,
    departure_station = ?,
    arrival_station = ?,
    departure_time = ?,
    passenger_id = ?,
    seat_class = ?,
    ticket_price = ?,
    seat_number = ?,
    updated_at = NOW()
WHERE id = ?;
-- name: DeleteOrderDetail :execresult
DELETE FROM pre_go_ticket_order_detail_052025_99999
WHERE id = ?;
-- name: GetOrderDetailById :one
SELECT id,
    ticket_item_id,
    order_number,
    passenger_name,
    departure_station,
    arrival_station,
    departure_time,
    passenger_id,
    seat_class,
    ticket_price,
    seat_number,
    created_at,
    updated_at
FROM pre_go_ticket_order_detail_052025_99999
WHERE id = ?;