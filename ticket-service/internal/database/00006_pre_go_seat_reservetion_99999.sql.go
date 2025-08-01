// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: 00006_pre_go_seat_reservetion_99999.sql

package database

import (
	"context"
	"database/sql"
)

const deleteSeatReservation = `-- name: DeleteSeatReservation :execresult
DELETE FROM pre_go_seat_reservation_99999
WHERE id = ?
`

func (q *Queries) DeleteSeatReservation(ctx context.Context, id int64) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteSeatReservation, id)
}

const deleteSeatReservationsByOrderNumber = `-- name: DeleteSeatReservationsByOrderNumber :execresult
DELETE FROM pre_go_seat_reservation_99999
WHERE order_number = ?
`

func (q *Queries) DeleteSeatReservationsByOrderNumber(ctx context.Context, orderNumber string) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteSeatReservationsByOrderNumber, orderNumber)
}

const deleteSeatReservationsByTrainId = `-- name: DeleteSeatReservationsByTrainId :execresult
DELETE FROM pre_go_seat_reservation_99999
WHERE train_id = ?
`

func (q *Queries) DeleteSeatReservationsByTrainId(ctx context.Context, trainID int64) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteSeatReservationsByTrainId, trainID)
}

const getSeatReservationById = `-- name: GetSeatReservationById :one
SELECT id,
    train_id,
    seat_id,
    order_number,
    from_station_id,
    to_station_id
FROM pre_go_seat_reservation_99999
WHERE id = ?
`

type GetSeatReservationByIdRow struct {
	ID            int64
	TrainID       int64
	SeatID        int64
	OrderNumber   string
	FromStationID int64
	ToStationID   int64
}

func (q *Queries) GetSeatReservationById(ctx context.Context, id int64) (GetSeatReservationByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getSeatReservationById, id)
	var i GetSeatReservationByIdRow
	err := row.Scan(
		&i.ID,
		&i.TrainID,
		&i.SeatID,
		&i.OrderNumber,
		&i.FromStationID,
		&i.ToStationID,
	)
	return i, err
}

const getSeatReservationsByOrderNumber = `-- name: GetSeatReservationsByOrderNumber :many
SELECT id,
    train_id,
    seat_id,
    order_number,
    from_station_id,
    to_station_id
FROM pre_go_seat_reservation_99999
WHERE order_number = ?
LIMIT ? OFFSET ?
`

type GetSeatReservationsByOrderNumberParams struct {
	OrderNumber string
	Limit       int32
	Offset      int32
}

type GetSeatReservationsByOrderNumberRow struct {
	ID            int64
	TrainID       int64
	SeatID        int64
	OrderNumber   string
	FromStationID int64
	ToStationID   int64
}

func (q *Queries) GetSeatReservationsByOrderNumber(ctx context.Context, arg GetSeatReservationsByOrderNumberParams) ([]GetSeatReservationsByOrderNumberRow, error) {
	rows, err := q.db.QueryContext(ctx, getSeatReservationsByOrderNumber, arg.OrderNumber, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetSeatReservationsByOrderNumberRow
	for rows.Next() {
		var i GetSeatReservationsByOrderNumberRow
		if err := rows.Scan(
			&i.ID,
			&i.TrainID,
			&i.SeatID,
			&i.OrderNumber,
			&i.FromStationID,
			&i.ToStationID,
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

const getSeatReservationsByTrainId = `-- name: GetSeatReservationsByTrainId :many
SELECT id,
    train_id,
    seat_id,
    order_number,
    from_station_id,
    to_station_id
FROM pre_go_seat_reservation_99999
WHERE train_id = ?
LIMIT ? OFFSET ?
`

type GetSeatReservationsByTrainIdParams struct {
	TrainID int64
	Limit   int32
	Offset  int32
}

type GetSeatReservationsByTrainIdRow struct {
	ID            int64
	TrainID       int64
	SeatID        int64
	OrderNumber   string
	FromStationID int64
	ToStationID   int64
}

func (q *Queries) GetSeatReservationsByTrainId(ctx context.Context, arg GetSeatReservationsByTrainIdParams) ([]GetSeatReservationsByTrainIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getSeatReservationsByTrainId, arg.TrainID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetSeatReservationsByTrainIdRow
	for rows.Next() {
		var i GetSeatReservationsByTrainIdRow
		if err := rows.Scan(
			&i.ID,
			&i.TrainID,
			&i.SeatID,
			&i.OrderNumber,
			&i.FromStationID,
			&i.ToStationID,
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

const insertSeatReservation = `-- name: InsertSeatReservation :execresult
INSERT INTO pre_go_seat_reservation_99999 (
        train_id,
        seat_id,
        order_number,
        from_station_id,
        to_station_id,
        updated_at,
        created_at
    )
VALUES (?, ?, ?, ?, ?, NOW(), NOW())
`

type InsertSeatReservationParams struct {
	TrainID       int64
	SeatID        int64
	OrderNumber   string
	FromStationID int64
	ToStationID   int64
}

func (q *Queries) InsertSeatReservation(ctx context.Context, arg InsertSeatReservationParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, insertSeatReservation,
		arg.TrainID,
		arg.SeatID,
		arg.OrderNumber,
		arg.FromStationID,
		arg.ToStationID,
	)
}

const updateSeatReservation = `-- name: UpdateSeatReservation :execresult
UPDATE pre_go_seat_reservation_99999
SET train_id = ?,
    seat_id = ?,
    order_number = ?,
    from_station_id = ?,
    to_station_id = ?,
    updated_at = NOW()
WHERE id = ?
`

type UpdateSeatReservationParams struct {
	TrainID       int64
	SeatID        int64
	OrderNumber   string
	FromStationID int64
	ToStationID   int64
	ID            int64
}

func (q *Queries) UpdateSeatReservation(ctx context.Context, arg UpdateSeatReservationParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateSeatReservation,
		arg.TrainID,
		arg.SeatID,
		arg.OrderNumber,
		arg.FromStationID,
		arg.ToStationID,
		arg.ID,
	)
}
