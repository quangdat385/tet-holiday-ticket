package model

type SeatReservationCreateInput struct {
	SeatID        int64  `json:"seat_id"`
	TrainID       int64  `json:"train_id"`
	OrderNumber   string `json:"order_number"`
	FromStationID int64  `json:"from_station_id"`
	ToStationID   int64  `json:"to_station_id"`
}
type SeatReservationUpdateInput struct {
	ID            int64  `json:"id"`
	SeatID        int64  `json:"seat_id"`
	TrainID       int64  `json:"train_id"`
	OrderNumber   string `json:"order_number"`
	FromStationID int64  `json:"from_station_id"`
	ToStationID   int64  `json:"to_station_id"`
}

type SeatReservationOutput struct {
	ID            int64  `json:"id"`
	SeatID        int64  `json:"seat_id"`
	TrainID       int64  `json:"train_id"`
	OrderNumber   string `json:"order_number"`
	FromStationID int64  `json:"from_station_id"`
	ToStationID   int64  `json:"to_station_id"`
}
type SeatReservationListInput struct {
	OrderNumber string `json:"order_number"`
	TrainID     int64  `json:"train_id"`
	Page        int64  `json:"page"`
	Limit       int64  `json:"limit"`
}
