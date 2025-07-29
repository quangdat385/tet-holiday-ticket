package vo

import "time"

type TrainIDRequest struct {
	TrainID int `uri:"train_id" binding:"required"`
}
type CreateTrainRequest struct {
	Code               string    `json:"code" binding:"required"`
	Name               string    `json:"name" binding:"required"`
	DepartureStationID int       `json:"departure_station_id" binding:"required"`
	ArrivalStationID   int       `json:"arrival_station_id" binding:"required"`
	DepartureTime      time.Time `json:"departure_time" binding:"required"`
	ArrivalTime        time.Time `json:"arrival_time" binding:"required"`
	Status             int       `json:"status" binding:"-"`
	Direction          string    `json:"direction" binding:"required"`
	TrainType          string    `json:"train_type" binding:"required"`
}
type UpdateTrainStatusRequest struct {
	TrainID int `json:"train_id" binding:"required"`
	Status  int `json:"status" binding:"required"`
}
type UpdateTrainRequest struct {
	TrainID            int       `json:"train_id" binding:"required"`
	Code               string    `json:"code" binding:"-"`
	Name               string    `json:"name" binding:"-"`
	DepartureStationID int       `json:"departure_station_id" binding:"-"`
	ArrivalStationID   int       `json:"arrival_station_id" binding:"-"`
	DepartureTime      time.Time `json:"departure_time" binding:"-"`
	ArrivalTime        time.Time `json:"arrival_time" binding:"-"`
	Status             int       `json:"status" binding:"-"`
	Direction          string    `json:"direction" binding:"-"`
	TrainType          string    `json:"train_type" binding:"-"`
}
