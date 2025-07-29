package service

import (
	"context"

	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
)

type (
	ITrainService interface {
		GetTrainByID(context context.Context, id int64) (out model.TrainOutput, err error)
		UpdateTrain(context context.Context, id int64, train model.TrainInput) (out model.TrainOutput, err error)
		UpdateTrainStatus(context context.Context, id int64, status int32) (out model.TrainOutput, err error)
		DeleteTrain(context context.Context, id int64) (err error)
		CreateTrain(context context.Context, train model.CreateTrainInput) (out model.TrainOutput, err error)
	}
)

var (
	localTrainService ITrainService
)

func TrainService() ITrainService {
	if localTrainService == nil {
		panic("implement localTrainService not found for interface ITrainService")
	}
	return localTrainService
}
func InitTrainService(i ITrainService) {
	localTrainService = i
}
