package service

import (
	"context"

	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
)

type (
	ISeatService interface {
		GetSeatByID(contect context.Context, id int64) (out model.SeatOutput, err error)
		GetSeatsByTrainID(contect context.Context, in model.SeatListInput) (out []model.SeatOutput, err error)
		UpdateSeat(contect context.Context, in model.UpdateSeatInput) (out model.SeatOutput, err error)
		CreateSeat(contect context.Context, in model.CreateSeatInput) (out model.SeatOutput, err error)
		DeleteSeat(contect context.Context, id int64) (out bool, err error)
	}
)

var (
	localSeatService ISeatService
)

func SeatService() ISeatService {
	if localSeatService == nil {
		panic("implement localSeatService not found for interface ISeatService")
	}
	return localSeatService
}
func InitSeatService(i ISeatService) {
	localSeatService = i
}
