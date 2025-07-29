package impl

import (
	"context"
	"database/sql"

	"github.com/quangdat385/holiday-ticket/ticket-service/internal/database"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model/mapper"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/utils"
	"github.com/quangdat385/holiday-ticket/ticket-service/response"
)

type sTicket struct {
	r *database.Queries
}

func NewTicketImpl(r *database.Queries) *sTicket {
	return &sTicket{
		r: r,
	}
}

func (s *sTicket) GetTicketById(ctx context.Context, ticketId int) (out model.TicketOutput, err error) {
	result, err := s.r.GetTicketById(ctx, int64(ticketId))

	if err != nil {

		return out, err
	}
	if result.ID == 0 {
		return out, nil
	}
	out = mapper.ToTicketDTO(result)
	return out, nil
}
func (s *sTicket) GetAllTickets(ctx context.Context, in model.GetAllTicketsInput) (out []model.TicketOutput, err error) {
	result, err := s.r.GetAllTickets(ctx, database.GetAllTicketsParams{
		Limit:  in.Limit,
		Offset: (in.Page - 1) * in.Limit,
		Status: 1,
	})
	if err != nil {
		return out, err
	}
	if len(result) == 0 {
		return out, nil
	}
	for _, ticket := range result {
		if ticket.ID == 0 {
			continue
		}
		out = append(out, mapper.ToTicketDTO(ticket))
	}
	return out, nil
}
func (s *sTicket) CreateTicket(ctx context.Context, ticket model.CreateTicketInput) (out model.TicketOutput, err error) {
	result, err := s.r.InsertTicket(ctx, database.InsertTicketParams{
		Name:        ticket.Name,
		Description: sql.NullString{String: ticket.Description, Valid: true},
		StartTime:   ticket.StartTime,
		EndTime:     ticket.EndTime,
		Status:      ticket.Status,
	})
	if err != nil {
		return out, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil || lastInsertId == 0 {
		return out, nil
	}
	ticketResult, err := s.r.GetTicketById(ctx, lastInsertId)
	if err != nil {
		return out, err
	}
	out = mapper.ToTicketDTO(ticketResult)
	return out, nil
}
func (s *sTicket) UpdateTicket(ctx context.Context, ticket model.UpdateTicketInput) (out model.TicketOutput, err error) {
	result, err := s.r.GetTicketById(ctx, ticket.ID)
	if err != nil {
		return out, err
	}
	if result.ID == 0 {
		return out, response.ErrTicketNotFoundErr
	}
	if utils.CheckNil(ticket.Name) {
		result.Name = ticket.Name
	}
	if utils.CheckNil(ticket.Description) {
		result.Description = sql.NullString{String: ticket.Description, Valid: true}
	}
	if utils.CheckNil(ticket.Status) {
		result.Status = ticket.Status
	}
	err = s.r.UpdateTicket(ctx, database.UpdateTicketParams{
		ID:          ticket.ID,
		Name:        result.Name,
		Description: result.Description,
		StartTime:   result.StartTime,
		EndTime:     result.EndTime,
		Status:      result.Status,
	})
	if err != nil {
		return out, err
	}
	out = mapper.ToTicketDTO(result)
	return out, nil
}
func (s *sTicket) DeleteTicket(ctx context.Context, ticketId int) (err error) {
	result, err := s.r.GetTicketById(ctx, int64(ticketId))
	if err != nil {
		return err
	}
	if result.ID == 0 {
		return response.ErrTicketNotFoundErr
	}
	_, err = s.r.DeleteTicket(ctx, int64(ticketId))
	if err != nil {
		return err
	}
	return nil
}
