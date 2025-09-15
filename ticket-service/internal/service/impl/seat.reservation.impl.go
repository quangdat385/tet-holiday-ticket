package impl

import (
	"context"

	"github.com/quangdat385/holiday-ticket/ticket-service/internal/database"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model/mapper"
	"github.com/quangdat385/holiday-ticket/ticket-service/response"
)

type sSeatReservation struct {
	r *database.Queries
}

func NewSeatReservation(r *database.Queries) *sSeatReservation {
	return &sSeatReservation{
		r: r,
	}
}
func (s *sSeatReservation) GetSeatReservationById(context context.Context, id int64) (out model.SeatReservationOutput, err error) {
	seatReservation, err := s.r.GetSeatReservationById(context, id)
	if err != nil {
		return out, err
	}
	if seatReservation.ID == 0 {
		return out, response.ErrSeatReserverNotFoundErr
	}
	out = mapper.ToSeatReservationDTO(seatReservation)
	return out, nil
}
func (s *sSeatReservation) GetSeatReservationByOrderNumber(context context.Context, in model.SeatReservationListInput) (out []model.SeatReservationOutput, err error) {
	offset := (in.Page - 1) * in.Limit
	seatReservations, err := s.r.GetSeatReservationsByOrderNumber(context, database.GetSeatReservationsByOrderNumberParams{
		OrderNumber: in.OrderNumber,
		Limit:       int32(in.Limit),
		Offset:      int32(offset),
	})

	if err != nil {
		return out, err
	}
	if len(seatReservations) == 0 {
		return out, response.ErrSeatReserverNotFoundErr
	}
	out = mapper.ToSeatReservationListDTO(seatReservations)
	return out, nil
}
func (s *sSeatReservation) GetSeatReservationByTrainId(context context.Context, in model.SeatReservationListInput) (out []model.SeatReservationOutput, err error) {
	seatReservations, err := s.r.GetSeatReservationsByTrainId(context, database.GetSeatReservationsByTrainIdParams{
		TrainID: in.TrainID,
		Limit:   int32(in.Limit),
		Offset:  int32((in.Page - 1) * in.Limit),
	})
	if err != nil {
		return out, err
	}
	if len(seatReservations) != 0 {
		return out, response.ErrSeatReserverNotFoundErr
	}
	for _, seatReservation := range seatReservations {
		if seatReservation.ID != 0 {
			seatReservationMapper := database.GetSeatReservationByIdRow(seatReservation)
			out = append(out, mapper.ToSeatReservationDTO(seatReservationMapper))
		}
	}
	return out, nil
}
func (s *sSeatReservation) CreateSeatReservation(context context.Context, in model.SeatReservationCreateInput) (out model.SeatReservationOutput, err error) {
	seatDB, err := s.r.GetSeatById(context, in.SeatID)
	if err != nil {
		return out, err
	}
	if seatDB.ID == 0 {
		return out, response.ErrSeatNotFoundErr
	}
	if seatDB.Status == 1 {
		return out, response.ErrSeatAlreadyReservedErr
	}
	result, err := s.r.InsertSeatReservation(context, database.InsertSeatReservationParams{
		SeatID:        in.SeatID,
		TrainID:       in.TrainID,
		OrderNumber:   in.OrderNumber,
		FromStationID: in.FromStationID,
		ToStationID:   in.ToStationID,
	})
	if err != nil {
		return out, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return out, err
	}
	seatReservation, err := s.r.GetSeatReservationById(context, id)
	if err != nil {
		return out, err
	}
	out = mapper.ToSeatReservationDTO(seatReservation)
	if out.ID == 0 {
		return out, response.ErrCreateSeatReservedErr
	}
	return out, nil
}
func (s *sSeatReservation) UpdateSeatReservation(context context.Context, in model.SeatReservationUpdateInput) (out model.SeatReservationOutput, err error) {
	seatReservation, err := s.r.GetSeatReservationById(context, in.ID)
	if err != nil {
		return out, err
	}
	if seatReservation.ID == 0 {
		return out, response.ErrNotFoundDataErr
	}
	if in.OrderNumber != "" {
		seatReservation.OrderNumber = in.OrderNumber
	}
	if in.FromStationID != 0 {
		seatReservation.FromStationID = in.FromStationID
	}
	if in.ToStationID != 0 {
		seatReservation.ToStationID = in.ToStationID
	}
	if in.TrainID != 0 {
		seatReservation.TrainID = in.TrainID
	}
	ok, err := s.r.UpdateSeatReservation(context, database.UpdateSeatReservationParams{
		ID:            seatReservation.ID,
		SeatID:        seatReservation.SeatID,
		TrainID:       seatReservation.TrainID,
		FromStationID: seatReservation.FromStationID,
		ToStationID:   seatReservation.ToStationID,
		OrderNumber:   seatReservation.OrderNumber,
	})
	if err != nil {
		return out, err
	}
	_, err = ok.LastInsertId()
	if err != nil {
		return out, err
	}
	out = mapper.ToSeatReservationDTO(seatReservation)
	return out, nil
}
func (s *sSeatReservation) DeleteSeatReservation(context context.Context, id int64) (out bool, err error) {
	seatReservation, err := s.r.GetSeatReservationById(context, id)
	if err != nil {
		return out, err
	}
	if seatReservation.ID == 0 {
		return out, response.ErrSeatReserverNotFoundErr
	}
	_, err = s.r.DeleteSeatReservation(context, id)
	if err != nil {
		return out, err
	}
	out = true
	return out, nil
}
