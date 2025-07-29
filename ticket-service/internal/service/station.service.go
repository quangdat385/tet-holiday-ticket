package service

import (
	"context"

	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
)

type (
	IStationService interface {
		GetStationByID(context context.Context, id int64) (out model.StationOutput, err error)
		GetAllStation(context context.Context, in model.StationListInput) (out []model.StationOutput, err error)
		CreateStation(context context.Context, station model.StationInput) (out model.StationOutput, err error)
		UpdateStation(context context.Context, id int64, station model.StationInput) (out model.StationOutput, err error)
		UpdateStationStatus(context context.Context, id int64, status int32) (out model.StationOutput, err error)
		DeleteStation(context context.Context, id int64) (out bool, err error)
	}
)

var (
	localStationService IStationService
)

func StationService() IStationService {
	if localStationService == nil {
		panic("implement localStationService not found for interface IStationService")
	}
	return localStationService
}
func InitStationService(i IStationService) {
	localStationService = i
}
