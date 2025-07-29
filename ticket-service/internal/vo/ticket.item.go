package vo

import "time"

type TicketItemRequest struct {
	TicketId int `uri:"ticket_id" binding:"required"`
}

type TickertItemCreateRequest struct {
	TicketId        int       `json:"ticket_id" binding:"required"`
	TicketName      string    `json:"ticket_name" binding:"required"`
	TrainID         int       `json:"train_id" binding:"required"`
	Description     string    `json:"description"`
	DepartureTime   time.Time `json:"departure_time" binding:"required"`
	SeatClass       string    `json:"seat_class" binding:"required"`
	StockAvailable  int       `json:"stock_available" binding:"required"`
	StockInitial    int       `json:"stock_initial" binding:"required"`
	IsStockPrepared bool      `json:"is_stock_prepared" binding:"required"`
	PriceOriginal   float32   `json:"price_original" binding:"required"`
	PriceFlash      float32   `json:"price_flash" binding:"required"`
	SaleStartTime   time.Time `json:"sale_start_time" binding:"required"`
	SaleEndTime     time.Time `json:"sale_end_time" binding:"required"`
	Status          int       `json:"status" binding:"required"`
	ActivityId      int64     `json:"activity_id" binding:"required"`
}
type UpdateTicketItemRequest struct {
	TicketItemId    int       `uri:"ticket_item_id" binding:"required"`
	TicketId        int       `json:"ticket_id,omitempty" binding:"-"`
	TicketName      string    `json:"ticket_name,omitempty" binding:"-"`
	TrainID         int       `json:"train_id,omitempty" binding:"-"`
	Description     string    `json:"description,omitempty" binding:"-"`
	DepartureTime   time.Time `json:"departure_time,omitempty" binding:"-"`
	SeatClass       string    `json:"seat_class,omitempty" binding:"-"`
	StockAvailable  int       `json:"stock_available,omitempty" binding:"-"`
	StockInitial    int       `json:"stock_initial,omitempty" binding:"-"`
	IsStockPrepared bool      `json:"is_stock_prepared,omitempty" binding:"-"`
	PriceOriginal   float32   `json:"price_original,omitempty" binding:"-"`
	PriceFlash      float32   `json:"price_flash,omitempty" binding:"-"`
	SaleStartTime   time.Time `json:"sale_start_time,omitempty" binding:"-"`
	SaleEndTime     time.Time `json:"sale_end_time,omitempty" binding:"-"`
	Status          int       `json:"status,omitempty" binding:"-"`
	ActivityId      int64     `json:"activity_id,omitempty" binding:"-"`
}
