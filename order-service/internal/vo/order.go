package vo

import "time"

type OrderIdRequest struct {
	OrderId int32 `json:"order_id" binding:"required"`
}

type CreateOrderRequest struct {
	OrderNUmber string    `json:"order_number" binding:"required"`
	OrderAmount float32   `json:"order_amount" binding:"required"`
	TerminalID  int64     `json:"terminal_id" binding:"required"`
	OrderDate   time.Time `json:"order_date" binding:"required"`
	OrderNotes  string    `json:"order_notes" binding:"-"`
}
type UpdateOrderRequest struct {
	OrderId     int32     `json:"order_id" binding:"required"`
	OrderNUmber string    `json:"order_number" binding:"-"`
	OrderAmount float32   `json:"order_amount" binding:"-"`
	TerminalID  int64     `json:"terminal_id" binding:"-"`
	OrderDate   time.Time `json:"order_date" binding:"-"`
	OrderNotes  string    `json:"order_notes" binding:"-"`
}
