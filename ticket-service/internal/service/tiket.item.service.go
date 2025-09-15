package service

import (
	"context"

	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
)

type (
	ITicketItem interface {
		GetTicketItemById(ctx context.Context, ticketId int, version int) (out model.TicketItemsOutput, err error)
		CreateTicketItem(ctx context.Context, input model.TicketItemInPut) (out model.TicketItemsOutput, err error)
		UpdateTicketItem(ctx context.Context, input model.UpdateTicketItemInPut) (out model.TicketItemsOutput, err error)
		DeleteTicketItem(ctx context.Context, ticketItemId int) (err error)
		DecreaseStock(ctx context.Context, ticketId int, quantity int) (out int, code int)
		SetStockCache(ctx context.Context, ticketId int, stock int, expiration int) (message string, err error)
	}
)

var (
	localTicketItem ITicketItem
)

func TicketItem() ITicketItem {
	if localTicketItem == nil {
		panic("implement localTicketItem not found for interface ITicketItem")
	}

	return localTicketItem
}

func InitTicketItem(i ITicketItem) {
	localTicketItem = i
}
