package model

import (
	"time"
)

type TicketOutput struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	StartTime time.Time `json:"start_time"`

	EndTime time.Time `json:"end_time"`

	Status int32 `json:"status"`
}
type CreateTicketInput struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Status      int32     `json:"status"`
}
type CreateTicketOutput struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Status      int32     `json:"status"`
}
type UpdateTicketInput struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int32  `json:"status"`
}
type GetAllTicketsInput struct {
	Page   int32 `json:"page"`
	Limit  int32 `json:"limit"`
	Status int32 `json:"status"`
}
