package mapper

import (
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/database"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
)

func ToTicketDTO(ticket database.PreGoTicket99999) model.TicketOutput {
	return model.TicketOutput{
		ID:   int64(ticket.ID),
		Name: ticket.Name,
		Description: func() string {
			if ticket.Description.Valid {
				return ticket.Description.String
			}
			return ""
		}(),
		StartTime: ticket.StartTime.UTC(),
		EndTime:   ticket.EndTime.UTC(),
		Status:    int32(ticket.Status),
	}
}
