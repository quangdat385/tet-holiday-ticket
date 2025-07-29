package vo

type SeatReservationIDRequest struct {
	SeatReservationID int `uri:"seat_reservation_id" binding:"required"`
}
type CreateSeatReservationRequest struct {
	TrainID       int    `json:"train_id" binding:"required"`
	SeatID        int    `json:"seat_id" binding:"required"`
	OrderNumber   string `json:"order_number" binding:"required"`
	FromStationID int    `json:"from_station_id" binding:"required"`
	ToStationID   int    `json:"to_station_id" binding:"required"`
}
type UpdateSeatReservationRequest struct {
	SeatReservationID int    `json:"seat_reservation_id" binding:"required"`
	TrainID           int    `json:"train_id" binding:"optional"`
	SeatID            int    `json:"seat_id" binding:"optional"`
	OrderNumber       string `json:"order_number" binding:"optional"`
	FromStationID     int    `json:"from_station_id" binding:"optional"`
	ToStationID       int    `json:"to_station_id" binding:"optional"`
}
type SeatReservationListRequest struct {
	TrainID       int    `form:"train_id" binding:"-"`
	OrderNumber   string `form:"order_number" binding:"-"`
	FromStationID int    `form:"from_station_id" binding:"-"`
	ToStationID   int    `form:"to_station_id" binding:"-"`
	Limit         int    `form:"limit" binding:"required"`
	Page          int    `form:"page" binding:"required"`
}
