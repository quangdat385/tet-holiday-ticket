// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: 00007_pre_go_route_segment_99999.sql

package database

import (
	"context"
	"database/sql"
)

const deleteRouteSegment = `-- name: DeleteRouteSegment :execresult
DELETE FROM pre_go_route_segment_99999
WHERE id = ?
`

func (q *Queries) DeleteRouteSegment(ctx context.Context, id int64) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteRouteSegment, id)
}

const getRouteSegmentById = `-- name: GetRouteSegmentById :one
SELECT id,
    train_id,
    from_station_id,
    to_station_id,
    segment_order,
    distance_km
FROM pre_go_route_segment_99999
WHERE id = ?
`

type GetRouteSegmentByIdRow struct {
	ID            int64
	TrainID       int64
	FromStationID int64
	ToStationID   int64
	SegmentOrder  int32
	DistanceKm    int32
}

func (q *Queries) GetRouteSegmentById(ctx context.Context, id int64) (GetRouteSegmentByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getRouteSegmentById, id)
	var i GetRouteSegmentByIdRow
	err := row.Scan(
		&i.ID,
		&i.TrainID,
		&i.FromStationID,
		&i.ToStationID,
		&i.SegmentOrder,
		&i.DistanceKm,
	)
	return i, err
}

const getRouteSegmentBySegmentOrder = `-- name: GetRouteSegmentBySegmentOrder :one
SELECT id,
    train_id,
    from_station_id,
    to_station_id,
    segment_order,
    distance_km
FROM pre_go_route_segment_99999
WHERE train_id = ?
    AND segment_order = ?
`

type GetRouteSegmentBySegmentOrderParams struct {
	TrainID      int64
	SegmentOrder int32
}

type GetRouteSegmentBySegmentOrderRow struct {
	ID            int64
	TrainID       int64
	FromStationID int64
	ToStationID   int64
	SegmentOrder  int32
	DistanceKm    int32
}

func (q *Queries) GetRouteSegmentBySegmentOrder(ctx context.Context, arg GetRouteSegmentBySegmentOrderParams) (GetRouteSegmentBySegmentOrderRow, error) {
	row := q.db.QueryRowContext(ctx, getRouteSegmentBySegmentOrder, arg.TrainID, arg.SegmentOrder)
	var i GetRouteSegmentBySegmentOrderRow
	err := row.Scan(
		&i.ID,
		&i.TrainID,
		&i.FromStationID,
		&i.ToStationID,
		&i.SegmentOrder,
		&i.DistanceKm,
	)
	return i, err
}

const getRouteSegmentsByFromStationId = `-- name: GetRouteSegmentsByFromStationId :many
SELECT id,
    train_id,
    from_station_id,
    to_station_id,
    segment_order,
    distance_km
FROM pre_go_route_segment_99999
WHERE from_station_id = ?
`

type GetRouteSegmentsByFromStationIdRow struct {
	ID            int64
	TrainID       int64
	FromStationID int64
	ToStationID   int64
	SegmentOrder  int32
	DistanceKm    int32
}

func (q *Queries) GetRouteSegmentsByFromStationId(ctx context.Context, fromStationID int64) ([]GetRouteSegmentsByFromStationIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getRouteSegmentsByFromStationId, fromStationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRouteSegmentsByFromStationIdRow
	for rows.Next() {
		var i GetRouteSegmentsByFromStationIdRow
		if err := rows.Scan(
			&i.ID,
			&i.TrainID,
			&i.FromStationID,
			&i.ToStationID,
			&i.SegmentOrder,
			&i.DistanceKm,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRouteSegmentsByToStationId = `-- name: GetRouteSegmentsByToStationId :many
SELECT id,
    train_id,
    from_station_id,
    to_station_id,
    segment_order,
    distance_km
FROM pre_go_route_segment_99999
WHERE to_station_id = ?
`

type GetRouteSegmentsByToStationIdRow struct {
	ID            int64
	TrainID       int64
	FromStationID int64
	ToStationID   int64
	SegmentOrder  int32
	DistanceKm    int32
}

func (q *Queries) GetRouteSegmentsByToStationId(ctx context.Context, toStationID int64) ([]GetRouteSegmentsByToStationIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getRouteSegmentsByToStationId, toStationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRouteSegmentsByToStationIdRow
	for rows.Next() {
		var i GetRouteSegmentsByToStationIdRow
		if err := rows.Scan(
			&i.ID,
			&i.TrainID,
			&i.FromStationID,
			&i.ToStationID,
			&i.SegmentOrder,
			&i.DistanceKm,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getRouteSegmentsByTrainId = `-- name: GetRouteSegmentsByTrainId :many
SELECT id,
    train_id,
    from_station_id,
    to_station_id,
    segment_order,
    distance_km
FROM pre_go_route_segment_99999
WHERE train_id = ?
`

type GetRouteSegmentsByTrainIdRow struct {
	ID            int64
	TrainID       int64
	FromStationID int64
	ToStationID   int64
	SegmentOrder  int32
	DistanceKm    int32
}

func (q *Queries) GetRouteSegmentsByTrainId(ctx context.Context, trainID int64) ([]GetRouteSegmentsByTrainIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getRouteSegmentsByTrainId, trainID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRouteSegmentsByTrainIdRow
	for rows.Next() {
		var i GetRouteSegmentsByTrainIdRow
		if err := rows.Scan(
			&i.ID,
			&i.TrainID,
			&i.FromStationID,
			&i.ToStationID,
			&i.SegmentOrder,
			&i.DistanceKm,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertRouteSegment = `-- name: InsertRouteSegment :execresult
INSERT INTO pre_go_route_segment_99999 (
        train_id,
        from_station_id,
        to_station_id,
        segment_order,
        distance_km,
        updated_at,
        created_at
    )
VALUES (?, ?, ?, ?, ?, NOW(), NOW())
`

type InsertRouteSegmentParams struct {
	TrainID       int64
	FromStationID int64
	ToStationID   int64
	SegmentOrder  int32
	DistanceKm    int32
}

func (q *Queries) InsertRouteSegment(ctx context.Context, arg InsertRouteSegmentParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertRouteSegment,
		arg.TrainID,
		arg.FromStationID,
		arg.ToStationID,
		arg.SegmentOrder,
		arg.DistanceKm,
	)
}

const updateRouteSegment = `-- name: UpdateRouteSegment :execresult
UPDATE pre_go_route_segment_99999
SET train_id = ?,
    from_station_id = ?,
    to_station_id = ?,
    segment_order = ?,
    distance_km = ?,
    updated_at = NOW()
WHERE id = ?
`

type UpdateRouteSegmentParams struct {
	TrainID       int64
	FromStationID int64
	ToStationID   int64
	SegmentOrder  int32
	DistanceKm    int32
	ID            int64
}

func (q *Queries) UpdateRouteSegment(ctx context.Context, arg UpdateRouteSegmentParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateRouteSegment,
		arg.TrainID,
		arg.FromStationID,
		arg.ToStationID,
		arg.SegmentOrder,
		arg.DistanceKm,
		arg.ID,
	)
}
