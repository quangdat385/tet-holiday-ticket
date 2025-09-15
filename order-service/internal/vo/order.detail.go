package vo

import "time"

type OrderDetailIDRequest struct {
	ID int64 `uri:"id" binding:"required"`
}
type CreateOrderDetailRequest struct {
	TicketItemID     int64     `json:"ticket_item_id" binding:"required"`
	OrderNumber      string    `json:"order_number" binding:"required"`
	PassengerName    string    `json:"passenger_name" binding:"required"`
	DepartureStation string    `json:"departure_station" binding:"required"`
	ArrivalStation   string    `json:"arrival_station" binding:"required"`
	DepartureTime    time.Time `json:"departure_time" binding:"required"`
	PassengerID      int64     `json:"passenger_id" binding:"required"`
	SeatClass        string    `json:"seat_class" binding:"required"`
	TicketPrice      float64   `json:"ticket_price" binding:"required"`
	SeatNumber       string    `json:"seat_number" binding:"required"`
}
type UpdateOrderDetailRequest struct {
	OrderDetailID    int64     `json:"order_detail_id" binding:"required"`
	TicketItemID     int64     `json:"ticket_item_id" binding:"-"`
	OrderNumber      string    `json:"order_number" binding:"-"`
	PassengerName    string    `json:"passenger_name" binding:"-"`
	DepartureStation string    `json:"departure_station" binding:"-"`
	ArrivalStation   string    `json:"arrival_station" binding:"-"`
	DepartureTime    time.Time `json:"departure_time" binding:"-"`
	PassengerID      int64     `json:"passenger_id" binding:"-"`
	SeatClass        string    `json:"seat_class" binding:"-"`
	TicketPrice      float64   `json:"ticket_price" binding:"-"`
	SeatNumber       string    `json:"seat_number" binding:"-"`
}
