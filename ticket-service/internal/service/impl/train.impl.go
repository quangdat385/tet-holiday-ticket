package impl

import (
	"context"

	"github.com/quangdat385/holiday-ticket/ticket-service/internal/database"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model/mapper"
)

type sTrain struct {
	r *database.Queries
}

func NewTrainImpl(r *database.Queries) *sTrain {
	return &sTrain{
		r: r,
	}
}
func (s *sTrain) GetTrainByID(context context.Context, id int64) (out model.TrainOutput, err error) {
	train, err := s.r.GetTrainById(context, id)
	if err != nil {
		return out, err
	}
	out = mapper.ToTrainDTO(train)
	return out, nil
}
func (s *sTrain) UpdateTrain(context context.Context, id int64, train model.TrainInput) (out model.TrainOutput, err error) {
	trainDB, err := s.r.GetTrainById(context, id)
	if err != nil {
		return out, err
	}
	if train.Name != "" {
		trainDB.Name = train.Name
	}
	if !train.DepartureTime.IsZero() {
		trainDB.DepartureTime = train.DepartureTime
	}
	if !train.ArrivalTime.IsZero() {
		trainDB.ArrivalTime = train.ArrivalTime
	}
	if train.Status != 0 {
		trainDB.Status = train.Status
	}
	if train.Direction != "" {
		trainDB.Direction = database.PreGoTrain99999Direction(train.Direction)
	}
	if train.TrainType != "" {
		trainDB.TrainType = database.PreGoTrain99999TrainType(train.TrainType)
	}
	result, err := s.r.UpdateTrain(context, database.UpdateTrainParams{
		ID:                 trainDB.ID,
		Name:               trainDB.Name,
		Code:               trainDB.Code,
		DepartureStationID: trainDB.DepartureStationID,
		ArrivalStationID:   trainDB.ArrivalStationID,
		DepartureTime:      trainDB.DepartureTime,
		ArrivalTime:        trainDB.ArrivalTime,
		Status:             trainDB.Status,
		Direction:          trainDB.Direction,
		TrainType:          trainDB.TrainType,
	})
	if err != nil {
		return out, err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return out, err
	}
	trainDB.DepartureTime = train.DepartureTime
	trainDB.ArrivalTime = train.ArrivalTime
	trainDB.Status = train.Status
	trainDB.Direction = database.PreGoTrain99999Direction(train.Direction)
	out = mapper.ToTrainDTO(trainDB)
	return out, nil
}
func (s *sTrain) UpdateTrainStatus(context context.Context, id int64, status int32) (out model.TrainOutput, err error) {
	trainDB, err := s.r.GetTrainById(context, id)
	if err != nil {
		return out, err
	}
	result, err := s.r.UpdateTrainStatus(context, database.UpdateTrainStatusParams{
		ID:     trainDB.ID,
		Status: status,
	})
	if err != nil {
		return out, err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return out, err
	}
	trainDB.Status = status
	out = mapper.ToTrainDTO(trainDB)
	return out, nil
}
func (s *sTrain) DeleteTrain(context context.Context, id int64) (err error) {
	_, err = s.r.DeleteTrain(context, id)
	if err != nil {
		return err
	}
	return nil
}
func (s *sTrain) CreateTrain(context context.Context, train model.CreateTrainInput) (out model.TrainOutput, err error) {
	result, err := s.r.InsertTrain(context, database.InsertTrainParams{
		Name:               train.Name,
		Code:               train.Code,
		DepartureStationID: train.DepartureStationID,
		ArrivalStationID:   train.ArrivalStationID,
		DepartureTime:      train.DepartureTime,
		ArrivalTime:        train.ArrivalTime,
		Status:             train.Status,
		Direction:          database.PreGoTrain99999Direction(train.Direction),
		TrainType:          database.PreGoTrain99999TrainType(train.TrainType),
	})
	if err != nil {
		return out, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return out, err
	}
	trainDB, err := s.r.GetTrainById(context, id)
	if err != nil {
		return out, err
	}
	out = mapper.ToTrainDTO(trainDB)
	return out, nil
}
