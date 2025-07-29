package service

import (
	"context"

	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
)

type (
	// IStationService interface
	ISeatReservationServive interface {
		GetSeatReservationById(context context.Context, id int64) (out model.SeatReservationOutput, err error)
		GetSeatReservationByOrderNumber(context context.Context, in model.SeatReservationListInput) (out []model.SeatReservationOutput, err error)
		GetSeatReservationByTrainId(context context.Context, in model.SeatReservationListInput) (out []model.SeatReservationOutput, err error)
		CreateSeatReservation(context context.Context, in model.SeatReservationCreateInput) (out model.SeatReservationOutput, err error)
		UpdateSeatReservation(context context.Context, in model.SeatReservationUpdateInput) (out model.SeatReservationOutput, err error)
		DeleteSeatReservation(context context.Context, id int64) (out bool, err error)
	}
)

var (
	localSeatReservationService ISeatReservationServive
)

func SeatReservationService() ISeatReservationServive {
	if localSeatReservationService == nil {
		panic("implement localSeatReservationService not found for interface ISeatReservationServive")
	}
	return localSeatReservationService
}
func InitSeatReservationService(i ISeatReservationServive) {
	localSeatReservationService = i
}
