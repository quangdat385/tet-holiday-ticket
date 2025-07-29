package model

import "time"

type OrderDetailInput struct {
	TicketItemID     int64     `json:"ticket_item_id"`
	OrderNumber      string    `json:"order_number"`
	PassengerName    string    `json:"passenger_name"`
	DepartureStation string    `json:"departure_station"`
	ArrivalStation   string    `json:"arrival_station"`
	DepartureTime    time.Time `json:"departure_time"`
	PassengerID      int64     `json:"departure_id"`
	SeatClass        string    `json:"seat_class"`
	TicketPrice      float64   `json:"ticket_price"`
	SeatNumber       string    `json:"seat_number"`
}

type OrderDetailOutput struct {
	ID               int64     `json:"id"`
	TicketItemID     int64     `json:"ticket_item_id"`
	OrderNumber      string    `json:"order_number"`
	PassengerName    string    `json:"passenger_name"`
	DepartureStation string    `json:"departure_station"`
	ArrivalStation   string    `json:"arrival_station"`
	DepartureTime    time.Time `json:"departure_time"`
	PassengerID      int64     `json:"departure_id"`
	SeatClass        string    `json:"seat_class"`
	TicketPrice      float64   `json:"ticket_price"`
	SeatNumber       string    `json:"seat_number"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
