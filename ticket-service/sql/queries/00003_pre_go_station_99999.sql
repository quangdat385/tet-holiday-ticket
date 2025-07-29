-- name: InsertStation :execresult
INSERT INTO pre_go_station_99999 (
        name,
        code,
        updated_at,
        created_at
    )
VALUES (?, ?, NOW(), NOW());
-- name: GetStationById :one
SELECT id,
    name,
    code,
    status
FROM pre_go_station_99999
WHERE id = ?;
-- name: GetStationByCode :one  
SELECT id,
    name,
    code,
    status
FROM pre_go_station_99999
WHERE code = ?;
-- name: UpdateStation :execresult
UPDATE pre_go_station_99999
SET name = ?,
    code = ?,
    status = ?,
    updated_at = NOW()
WHERE id = ?;
-- name: DeleteStation :execresult
DELETE FROM pre_go_station_99999
WHERE id = ?;
-- name: GetStationList :many
SELECT id,
    name,
    code,
    status
FROM pre_go_station_99999
WHERE status = ?
LIMIT ? OFFSET ?;