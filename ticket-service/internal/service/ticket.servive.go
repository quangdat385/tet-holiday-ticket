package service

import (
	"context"

	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
)

type (
	ITicketHome interface {
		GetTicketById(ctx context.Context, ticketId int) (out model.TicketOutput, err error)
		GetAllTickets(ctx context.Context, in model.GetAllTicketsInput) (out []model.TicketOutput, err error)
		CreateTicket(ctx context.Context, ticket model.CreateTicketInput) (out model.TicketOutput, err error)
		UpdateTicket(ctx context.Context, ticket model.UpdateTicketInput) (out model.TicketOutput, err error)
		DeleteTicket(ctx context.Context, ticketId int) (err error)
	}
)

var (
	localTicketHome ITicketHome
)

func TicketHome() ITicketHome {
	if localTicketHome == nil {
		panic("implement localTicketHome not found for interface ITicketHome")
	}

	return localTicketHome
}

func InitTicketHome(i ITicketHome) {
	localTicketHome = i
}
