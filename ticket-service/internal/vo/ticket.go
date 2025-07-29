package vo

import "time"

type TicketIdRequest struct {
	TicketId int `uri:"ticket_id" binding:"required"`
}

type CreateTicketRequest struct {
	Name        string    `json:"name" form:"name" binding:"required"`
	Description string    `json:"description" form:"description" binding:"-"`
	StartTime   time.Time `json:"start_time"  binding:"required"`
	EndTime     time.Time `json:"end_time" form:"end_time" binding:"required"`
	Status      int32     `json:"status" form:"status" binding:"required"`
}
type UpdateTicketRequest struct {
	TicketId    int    `json:"ticket_id" form:"ticket_id" binding:"required"`
	Name        string `json:"name" form:"name" binding:"-"`
	Description string `json:"description" form:"description" binding:"-"`
	Status      int32  `json:"status" form:"status" binding:"-"`
}
type GetAllTicketsRequest struct {
	Page   int `form:"page" binding:"required,min=1"`
	Limit  int `form:"limit" binding:"required,min=50,max=100"`
	Status int `form:"status" binding:"omitempty,oneof=0 1"`
}
