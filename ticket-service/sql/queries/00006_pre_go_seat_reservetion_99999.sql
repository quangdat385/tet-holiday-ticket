-- name: InsertSeatReservation :execresult
INSERT INTO pre_go_seat_reservation_99999 (
        train_id,
        seat_id,
        order_number,
        from_station_id,
        to_station_id,
        updated_at,
        created_at
    )
VALUES (?, ?, ?, ?, ?, NOW(), NOW());
-- name: GetSeatReservationById :one
SELECT id,
    train_id,
    seat_id,
    order_number,
    from_station_id,
    to_station_id
FROM pre_go_seat_reservation_99999
WHERE id = ?;
-- name: GetSeatReservationsByTrainId :many
SELECT id,
    train_id,
    seat_id,
    order_number,
    from_station_id,
    to_station_id
FROM pre_go_seat_reservation_99999
WHERE train_id = ?
LIMIT ? OFFSET ?;
-- name: GetSeatReservationsByOrderNumber :many
SELECT id,
    train_id,
    seat_id,
    order_number,
    from_station_id,
    to_station_id
FROM pre_go_seat_reservation_99999
WHERE order_number = ?
LIMIT ? OFFSET ?;
-- name: UpdateSeatReservation :execresult
UPDATE pre_go_seat_reservation_99999
SET train_id = ?,
    seat_id = ?,
    order_number = ?,
    from_station_id = ?,
    to_station_id = ?,
    updated_at = NOW()
WHERE id = ?;
-- name: DeleteSeatReservation :execresult
DELETE FROM pre_go_seat_reservation_99999
WHERE id = ?;
-- name: DeleteSeatReservationsByTrainId :execresult
DELETE FROM pre_go_seat_reservation_99999
WHERE train_id = ?;
-- name: DeleteSeatReservationsByOrderNumber :execresult
DELETE FROM pre_go_seat_reservation_99999
WHERE order_number = ?;