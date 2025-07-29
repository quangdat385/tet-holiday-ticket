package model

import "time"

type OrderInput struct {
	OrderNumber string    `json:"order_number"`
	OrderAmount float32   `json:"order_amount"`
	TerminalID  int64     `json:"terminal_id"`
	OrderDate   time.Time `json:"order_date"`
}
type OrderEventInput struct {
	OrderNumber string    `json:"order_number"`
	OrderAmount float32   `json:"order_amount"`
	TerminalID  int64     `json:"terminal_id"`
	OrderDate   time.Time `json:"order_date"`
	OrderNotes  string    `json:"order_notes"`
	OrderItem   OrderItem `json:"order_item"`
}
type OrderItem struct {
	ItemID    int64   `json:"item_id"`
	ItemName  string  `json:"item_name"`
	ItemPrice float32 `json:"item_price"`
	ItemCount int32   `json:"item_count"`
}

type OrderOutPut struct {
	ID          int32     `json:"id"`
	OrderNumber string    `json:"order_number"`
	OrderAmount float32   `json:"order_amount"`
	TerminalID  int64     `json:"terminal_id"`
	OrderDate   time.Time `json:"order_date"`
	OrderNotes  string    `json:"order_notes"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type OrderEvent struct {
	Type  string          `json:"type"`
	Order OrderEventInput `json:"order"`
}

const (
	OrderEventTypeCreateOrder string = "create-order"
	OrderEventTypeReOrder     string = "re-order"
	OrderEventCreatePayment   string = "payment-order"
	OrderEventTypeCancelOrder string = "cancel-order"
	OrderEventTypeRefundOrder string = "refund-order"
)
