package model

import "time"

type OrderEventInput struct {
	OrderNumber string    `json:"order_number"`
	UserID      int64     `json:"user_id"`
	StationCode string    `json:"station_code"`
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
type OrderEvent struct {
	Type    string          `json:"type"`
	Order   OrderEventInput `json:"order"`
	Payment PaymentInput    `json:"payment"`
}
type PaymentInput struct {
	OrderNumber string `json:"order_number" binding:"required"`
	UserID      int64  `json:"user_id" binding:"required"`
}

const (
	OrderEventTypeCreateOrder    string = "create-order"
	OrderEventTypeConfirmOrder   string = "confirm-order"
	OrderEventTypeReConfirmOrder string = "re-confirm-order"
	OrderEventTypeOrderSuccess   string = "order-success"
	OrderEventTypeReOrder        string = "re-order"
	OrderEventTypeOrderDetail    string = "create-order-detail"
	OrderEventCreatePayment      string = "create-payment"
	OrderEventTypeCancelOrder    string = "cancel-order"
	OrderEventTypeRefundOrder    string = "refund-order"
)
