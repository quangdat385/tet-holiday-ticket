package impl

import (
	"context"

	"github.com/quangdat385/holiday-ticket/ticket-service/internal/database"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model/mapper"
	"github.com/quangdat385/holiday-ticket/ticket-service/response"
)

type sSeat struct {
	r *database.Queries
}

func NewSeatImpl(r *database.Queries) *sSeat {
	return &sSeat{
		r: r,
	}
}
func (s *sSeat) GetSeatByID(contect context.Context, id int64) (out model.SeatOutput, err error) {
	seat, err := s.r.GetSeatById(contect, id)
	if err != nil {
		return out, err
	}
	out = mapper.ToSeatDTO(seat)
	return out, nil
}
func (s *sSeat) GetSeatsByTrainID(contect context.Context, in model.SeatListInput) (out []model.SeatOutput, err error) {
	seats, err := s.r.GetSeatsByTrainId(contect, database.GetSeatsByTrainIdParams{
		TrainID: in.TrainID,
		Limit:   int32(in.Limit),
		Offset:  int32((in.Page - 1) * in.Limit),
	})
	if err != nil {
		return out, err
	}
	if len(seats) == 0 {
		return out, nil
	}
	out = mapper.ToSeatsDTO(seats)
	if len(out) == 0 {
		return nil, response.ErrSeatNotFoundErr
	}
	return out, nil
}
func (s *sSeat) UpdateSeat(contect context.Context, in model.UpdateSeatInput) (out model.SeatOutput, err error) {
	seatDb, err := s.r.GetSeatById(contect, in.ID)
	if err != nil {
		return out, err
	}
	if seatDb.ID == 0 {
		return out, response.ErrSeatNotFoundErr
	}
	seat, err := s.r.UpdateSeat(contect, database.UpdateSeatParams{
		ID:         in.ID,
		Status:     in.Status,
		SeatClass:  database.PreGoSeat99999SeatClass(in.SeatClass),
		SeatNumber: in.SeatNumber,
		TrainID:    seatDb.TrainID,
	})
	if err != nil {
		return out, err
	}
	_, err = seat.LastInsertId()
	if err != nil {
		return out, err
	}
	seatDb.SeatClass = database.PreGoSeat99999SeatClass(in.SeatClass)
	seatDb.SeatNumber = in.SeatNumber
	seatDb.Status = in.Status
	out = mapper.ToSeatDTO(seatDb)
	return out, nil
}
func (s *sSeat) CreateSeat(contect context.Context, in model.CreateSeatInput) (out model.SeatOutput, err error) {
	seat, err := s.r.InsertSeat(contect, database.InsertSeatParams{
		TrainID:    in.TrainID,
		SeatNumber: in.SeatNumber,
		SeatClass:  database.PreGoSeat99999SeatClass(in.SeatClass),
		Status:     in.Status,
	})
	if err != nil {
		return out, err
	}
	id, err := seat.LastInsertId()
	if err != nil {
		return out, err
	}
	seatDb, err := s.r.GetSeatById(contect, id)
	if err != nil {
		return out, err
	}
	out = mapper.ToSeatDTO(seatDb)
	return out, nil
}
func (s *sSeat) DeleteSeat(contect context.Context, id int64) (out bool, err error) {
	seatDb, err := s.r.GetSeatById(contect, id)
	if err != nil {
		return out, err
	}
	if seatDb.ID == 0 {
		return out, response.ErrSeatNotFoundErr
	}
	_, err = s.r.DeleteSeat(contect, id)
	if err != nil {
		return out, err
	}
	out = true
	return out, nil
}
