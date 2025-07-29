package model

type TicketSegmentPriceOutPut struct {
	ID             int64   `json:"id"`
	TicketItemID   int64   `json:"ticket_item_id"`
	RouteSegmentID int64   `json:"route_segment_id"`
	Price          float64 `json:"price"`
}
type TicketSegmentPriceCreateInPut struct {
	TicketItemID   int64   `json:"ticket_item_id"`
	RouteSegmentID int64   `json:"route_segment_id"`
	Price          float64 `json:"price"`
}
type TicketSegmentPriceUpdateInPut struct {
	ID             int64   `json:"id"`
	TicketItemID   int64   `json:"ticket_item_id"`
	RouteSegmentID int64   `json:"route_segment_id"`
	Price          float64 `json:"price"`
}
type TicketSegmentPriceListInPut struct {
	FromSegmentID int64 `json:"from_segment_id"`
	ToSegmentID   int64 `json:"to_segment_id"`
}
type GetAllTicketSegmentPriceInPut struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}
