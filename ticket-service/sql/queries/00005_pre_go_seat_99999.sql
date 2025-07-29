-- name: InsertSeat :execresult
INSERT INTO pre_go_seat_99999 (
        train_id,
        seat_number,
        seat_class,
        status,
        updated_at,
        created_at
    )
VALUES (?, ?, ?, ?, NOW(), NOW());
-- name: GetSeatById :one
SELECT id,
    train_id,
    seat_number,
    seat_class,
    status
FROM pre_go_seat_99999
WHERE id = ?;
-- name: GetSeatsByTrainId :many
SELECT id,
    train_id,
    seat_number,
    seat_class,
    status
FROM pre_go_seat_99999
WHERE train_id = ?
LIMIT ? OFFSET ?;
-- name: UpdateSeat :execresult
UPDATE pre_go_seat_99999
SET train_id = ?,
    seat_number = ?,
    seat_class = ?,
    status = ?,
    updated_at = NOW()
WHERE id = ?;
-- name: DeleteSeat :execresult
DELETE FROM pre_go_seat_99999
WHERE id = ?;