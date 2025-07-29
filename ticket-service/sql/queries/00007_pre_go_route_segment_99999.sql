-- name: InsertRouteSegment :execresult
INSERT INTO pre_go_route_segment_99999 (
        train_id,
        from_station_id,
        to_station_id,
        segment_order,
        distance_km,
        updated_at,
        created_at
    )
VALUES (?, ?, ?, ?, ?, NOW(), NOW());
-- name: GetRouteSegmentById :one
SELECT id,
    train_id,
    from_station_id,
    to_station_id,
    segment_order,
    distance_km
FROM pre_go_route_segment_99999
WHERE id = ?;
-- name: GetRouteSegmentsByTrainId :many
SELECT id,
    train_id,
    from_station_id,
    to_station_id,
    segment_order,
    distance_km
FROM pre_go_route_segment_99999
WHERE train_id = ?;
-- name: GetRouteSegmentsByFromStationId :many
SELECT id,
    train_id,
    from_station_id,
    to_station_id,
    segment_order,
    distance_km
FROM pre_go_route_segment_99999
WHERE from_station_id = ?;
-- name: GetRouteSegmentsByToStationId :many
SELECT id,
    train_id,
    from_station_id,
    to_station_id,
    segment_order,
    distance_km
FROM pre_go_route_segment_99999
WHERE to_station_id = ?;
-- name: GetRouteSegmentBySegmentOrder :one
SELECT id,
    train_id,
    from_station_id,
    to_station_id,
    segment_order,
    distance_km
FROM pre_go_route_segment_99999
WHERE train_id = ?
    AND segment_order = ?;
-- name: UpdateRouteSegment :execresult
UPDATE pre_go_route_segment_99999
SET train_id = ?,
    from_station_id = ?,
    to_station_id = ?,
    segment_order = ?,
    distance_km = ?,
    updated_at = NOW()
WHERE id = ?;
-- name: DeleteRouteSegment :execresult
DELETE FROM pre_go_route_segment_99999
WHERE id = ?;