package model

var SEAT_ENUM = []string{"ECONOMY", "BUSINESS", "FIRST"}

type CreateSeatInput struct {
	TrainID    int64  `json:"train_id"`
	SeatNumber string `json:"seat_number"`
	SeatClass  string `json:"seat_class"`
	Status     int32  `json:"status"`
}

type SeatOutput struct {
	ID         int64  `json:"id"`
	TrainID    int64  `json:"train_id"`
	SeatNumber string `json:"seat_number"`
	SeatClass  string `json:"seat_class"`
	Status     int32  `json:"status"`
}
type UpdateSeatInput struct {
	ID         int64  `json:"id"`
	SeatNumber string `json:"seat_number"`
	SeatClass  string `json:"seat_class"`
	Status     int32  `json:"status"`
}
type SeatListInput struct {
	TrainID int64 `json:"train_id"`
	Page    int64 `json:"page"`
	Limit   int64 `json:"limit"`
}
