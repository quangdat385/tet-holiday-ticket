package vo

import (
	"time"

	"github.com/quangdat385/holiday-ticket/order-service/internal/model"
)

type OrderIDRequest struct {
	OrderID int32 `uri:"order_id" binding:"required"`
}
type OrderNumberRequest struct {
	OrderNumber string `uri:"order_number" binding:"required"`
}
type UserIDRequest struct {
	UserID int64 `uri:"user_id" binding:"required"`
}
type QueriesOrderRequest struct {
	Page  int32 `form:"page" binding:"-" optional:"default=1"`
	Limit int32 `form:"page_size" binding:"-" optional:"default=50"`
}

type CreateOrderRequest struct {
	OrderNUmber string          `json:"order_number" binding:"required"`
	OrderAmount float32         `json:"order_amount" binding:"required"`
	TerminalID  int64           `json:"terminal_id" binding:"required"`
	UserID      int64           `json:"user_id" binding:"required"`
	StationCode string          `json:"station_code" binding:"required"`
	OrderDate   time.Time       `json:"order_date" binding:"required"`
	OrderNotes  string          `json:"order_notes" binding:"-"`
	OrderItem   model.OrderItem `json:"order_item" binding:"required"`
}
type UpdateOrderRequest struct {
	OrderId     int32     `json:"order_id" binding:"required"`
	StationCode string    `json:"station_code" binding:"-"`
	OrderNUmber string    `json:"order_number" binding:"-"`
	OrderAmount float32   `json:"order_amount" binding:"-"`
	TerminalID  int64     `json:"terminal_id" binding:"-"`
	OrderDate   time.Time `json:"order_date" binding:"-"`
	OrderNotes  string    `json:"order_notes" binding:"-"`
}
