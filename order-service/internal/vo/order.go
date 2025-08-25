package vo

import (
	"time"

	"github.com/quangdat385/holiday-ticket/order-service/internal/model"
)

type OrderIdRequest struct {
	OrderId int32 `json:"order_id" binding:"required"`
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
