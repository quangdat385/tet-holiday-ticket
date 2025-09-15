package vo

type TicketSegmentPriceIDRequest struct {
	TicketSegmentPriceID int `uri:"ticket_segment_price_id" binding:"required"`
}
type CreateTicketSegmentPriceRequest struct {
	TicketItemID   int     `json:"ticket_item_id" binding:"required"`
	RouteSegmentID int     `json:"route_segment_id" binding:"required"`
	Price          float64 `json:"price" binding:"required"`
}
type UpdateTicketSegmentPriceRequest struct {
	TicketSegmentPriceID int     `json:"ticket_segment_price_id" binding:"required"`
	TicketItemID         int     `json:"ticket_item_id" binding:"optional"`
	RouteSegmentID       int     `json:"route_segment_id" binding:"optional"`
	Price                float64 `json:"price" binding:"optional"`
}
type GetTicketSegmentPriceFromRouteSegmentIDToToSegmentIDRequest struct {
	FromSegmentID int `json:"from_segment_id" binding:"required"`
	ToSegmentID   int `json:"to_segment_id" binding:"required"`
}
type GetListTicketSegmentPriceRequest struct {
	Page  int `form:"page" binding:"required"`
	Limit int `form:"limit" binding:"required"`
}
