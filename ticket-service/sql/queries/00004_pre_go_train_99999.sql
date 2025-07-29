-- name: InsertTrain :execresult
INSERT INTO pre_go_train_99999 (
        name,
        code,
        departure_station_id,
        arrival_station_id,
        departure_time,
        arrival_time,
        status,
        direction,
        train_type,
        created_at,
        updated_at
    )
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW());
-- name: GetTrainById :one
SELECT id,
    name,
    code,
    departure_station_id,
    arrival_station_id,
    departure_time,
    arrival_time,
    status,
    direction,
    train_type
FROM pre_go_train_99999
WHERE id = ?;
-- name: UpdateTrain :execresult
UPDATE pre_go_train_99999
SET name = ?,
    code = ?,
    departure_station_id = ?,
    arrival_station_id = ?,
    departure_time = ?,
    arrival_time = ?,
    status = ?,
    direction = ?,
    train_type = ?,
    updated_at = NOW()
WHERE id = ?;
-- name: DeleteTrain :execresult
DELETE FROM pre_go_train_99999
WHERE id = ?;
-- name: GetAllTrain :many
SELECT id,
    name,
    code,
    departure_station_id,
    arrival_station_id,
    departure_time,
    arrival_time,
    status,
    direction train_type
FROM pre_go_train_99999
WHERE status = ?;
-- name: UpdateTrainStatus :execresult
UPDATE pre_go_train_99999
SET status = ?,
    updated_at = NOW()
WHERE id = ?;
-- name: GetTrainsByDepartureTimeRange :many
SELECT id,
    name,
    code,
    departure_station_id,
    arrival_station_id,
    departure_time,
    arrival_time,
    status,
    direction,
    train_type
FROM pre_go_train_99999
WHERE departure_time BETWEEN ? AND ?;