package model

import "time"

type TrainOutput struct {
	ID                 int64     `json:"id"`
	Name               string    `json:"name"`
	Code               string    `json:"code"`
	DepartureStationID int64     `json:"departure_station_id"`
	ArrivalStationID   int64     `json:"arrival_station_id"`
	DepartureTime      time.Time `json:"departure_time"`
	ArrivalTime        time.Time `json:"arrival_time"`
	TrainType          string    `json:"train_type"`
	Status             int32     `json:"status"`
	Direction          string    `json:"direction"`
}
type TrainInput struct {
	Name          string    `json:"name"`
	DepartureTime time.Time `json:"departure_time"`
	ArrivalTime   time.Time `json:"arrival_time"`
	Status        int32     `json:"status"`
	Direction     string    `json:"direction"`
	TrainType     string    `json:"train_type"`
}
type CreateTrainInput struct {
	Name               string    `json:"name"`
	Code               string    `json:"code"`
	DepartureStationID int64     `json:"departure_station_id"`
	ArrivalStationID   int64     `json:"arrival_station_id"`
	DepartureTime      time.Time `json:"departure_time"`
	ArrivalTime        time.Time `json:"arrival_time"`
	Status             int32     `json:"status"`
	Direction          string    `json:"direction"`
	TrainType          string    `json:"train_type"`
}
