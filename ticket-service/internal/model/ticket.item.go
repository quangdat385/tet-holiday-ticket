package model

import (
	"time"
)

type TicketItemInPut struct {
	TicketId        int       `json:"ticket_id"`         // Sửa tag và thêm field
	TicketName      string    `json:"ticket_name"`       // Sửa tag
	TrainID         int       `json:"train_id"`          // Sửa tag
	Description     string    `json:"description"`       // Sửa tag
	DepartureTime   time.Time `json:"departure_time"`    // Sửa tag
	SeatClass       string    `json:"seat_class"`        // Sửa tag
	StockAvailable  int       `json:"stock_available"`   // Sửa tag
	StockInitial    int       `json:"stock_initial"`     // Sửa tag
	IsStockPrepared bool      `json:"is_stock_prepared"` // Sửa tag
	PriceOriginal   float32   `json:"price_original"`    // Sửa tag
	PriceFlash      float32   `json:"price_flash"`       // Sửa tag
	SaleStartTime   time.Time `json:"sale_start_time"`   // Sửa tag
	SaleEndTime     time.Time `json:"sale_end_time"`     // Sửa tag
	Status          int       `json:"stutus"`            // Sửa tag
	ActivityId      int64     `json:"activity_id"`       // Sửa tag
}
type UpdateTicketItemInPut struct {
	TicketItemId    int       `json:"ticket_item_id"`              // Sửa tag
	TicketId        int       `json:"ticket_id,omitempty"`         // Sửa tag và thêm field
	TicketName      string    `json:"ticket_name,omitempty"`       // Sửa tag
	TrainID         int       `json:"train_id,omitempty"`          // Sửa tag
	Description     string    `json:"description,omitempty"`       // Sửa tag
	DepartureTime   time.Time `json:"departure_time,omitempty"`    // Sửa tag
	SeatClass       string    `json:"seat_class,omitempty"`        // Sửa tag
	StockAvailable  int       `json:"stock_available,omitempty"`   // Sửa tag
	StockInitial    int       `json:"stock_initial,omitempty"`     // Sửa tag
	IsStockPrepared bool      `json:"is_stock_prepared,omitempty"` // Sửa tag
	PriceOriginal   float32   `json:"price_original,omitempty"`    // Sửa tag
	PriceFlash      float32   `json:"price_flash,omitempty"`       // Sửa tag
	SaleStartTime   time.Time `json:"sale_start_time,omitempty"`   // Sửa tag
	SaleEndTime     time.Time `json:"sale_end_time,omitempty"`     // Sửa tag
	Status          int       `json:"stutus,omitempty"`            // Sửa tag
	ActivityId      int64     `json:"activity_id,omitempty"`       // Sửa tag
}

// VO: Get ticketItems returns
type TicketItemsOutput struct {
	TicketItemId    int       `json:"id"`                // Sửa tag
	TicketId        int       `json:"ticket_id"`         // Sửa tag và thêm field
	TrainID         int       `json:"train_id"`          // Sửa tag
	TicketName      string    `json:"ticket_name"`       // Sửa tag
	Description     string    `json:"description"`       // Sửa tag
	DepartureTime   time.Time `json:"departure_time"`    // Sửa tag
	SeatClass       string    `json:"seat_class"`        // Sửa tag
	StockAvailable  int       `json:"stock_available"`   // Sửa tag
	StockInitial    int       `json:"stock_initial"`     // Sửa tag
	IsStockPrepared bool      `json:"is_stock_prepared"` // Sửa tag
	PriceOriginal   float32   `json:"price_original"`    // Sửa tag
	PriceFlash      float32   `json:"price_flash"`       // Sửa tag
	SaleStartTime   time.Time `json:"sale_start_time"`   // Sửa tag
	SaleEndTime     time.Time `json:"sale_end_time"`     // Sửa tag
	Status          int       `json:"stutus"`            // Sửa tag
	ActivityId      int64     `json:"activity_id"`       // Sửa tag
	Version         int       `json:"version,omitempty"` // Sửa tag
}

type SetStockCacheRequest struct {
	TicketId int `json:"ticket_id"`
	Stock    int `json:"stock"`
	Duration int `json:"duration"`
}

func (t TicketItemsOutput) IsEmpty() bool {
	return t.TicketId == 0 && t.TicketName == "" && t.StockAvailable == 0 && t.StockInitial == 0
}
