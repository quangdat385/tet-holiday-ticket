-- name: InsertTicketSegmentPrice :execresult
INSERT INTO pre_go_ticket_segment_price_99999 (
        ticket_item_id,
        route_segment_id,
        price,
        updated_at,
        created_at
    )
VALUES (?, ?, ?, NOW(), NOW());
-- name: GetTicketSegmentPriceById :one
SELECT id,
    ticket_item_id,
    route_segment_id,
    price
FROM pre_go_ticket_segment_price_99999
WHERE id = ?;
-- name: GetTicketSegmentPricesByRouteSegmentId :one
SELECT id,
    ticket_item_id,
    route_segment_id,
    price
FROM pre_go_ticket_segment_price_99999
WHERE route_segment_id = ?;
-- name: GetTicketSegmentPricesFromSegmentIDToToSegmentID :many
SELECT id,
    ticket_item_id,
    route_segment_id,
    price
FROM pre_go_ticket_segment_price_99999
WHERE route_segment_id >= ?
    AND route_segment_id <= ?;
-- name: GetAllTicketSegmentPrices :many
SELECT id,
    ticket_item_id,
    route_segment_id,
    price
FROM pre_go_ticket_segment_price_99999
LIMIT ? OFFSET ?;
-- name: UpdateTicketSegmentPrice :execresult
UPDATE pre_go_ticket_segment_price_99999
SET ticket_item_id = ?,
    route_segment_id = ?,
    price = ?,
    updated_at = NOW()
WHERE id = ?;
-- name: DeleteTicketSegmentPrice :execresult
DELETE FROM pre_go_ticket_segment_price_99999
WHERE id = ?;