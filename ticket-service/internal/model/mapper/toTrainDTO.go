package mapper

import (
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/database"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
)

func ToTrainDTO(train database.GetTrainByIdRow) model.TrainOutput {
	return model.TrainOutput{
		ID:                 train.ID,
		Name:               train.Name,
		Code:               train.Code,
		DepartureStationID: train.DepartureStationID,
		ArrivalStationID:   train.ArrivalStationID,
		DepartureTime:      train.DepartureTime.UTC(),
		ArrivalTime:        train.ArrivalTime.UTC(),
		Status:             train.Status,
		Direction:          string(train.Direction),
		TrainType:          string(train.TrainType),
	}
}
