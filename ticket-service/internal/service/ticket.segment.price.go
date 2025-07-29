package service

import (
	"context"

	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
)

type (
	ITicketSegmentPriceService interface {
		GetTicketSegmentPriceById(context context.Context, id int64) (out model.TicketSegmentPriceOutPut, err error)
		GetTicketSegmentPriceByRouteSegmentId(context context.Context, routeSegmentId int64) (out model.TicketSegmentPriceOutPut, err error)
		GetAllTicketSegmentPrice(context context.Context, in model.GetAllTicketSegmentPriceInPut) (out []model.TicketSegmentPriceOutPut, err error)
		GetAllTicketSegmentPriceFromSegmentIDToToSegmentID(context context.Context, in model.TicketSegmentPriceListInPut) (out []model.TicketSegmentPriceOutPut, err error)
		CreateTicketSegmentPrice(context context.Context, ticketSegmentPrice model.TicketSegmentPriceCreateInPut) (out model.TicketSegmentPriceOutPut, err error)
		UpdateTicketSegmentPrice(context context.Context, in model.TicketSegmentPriceUpdateInPut) (out model.TicketSegmentPriceOutPut, err error)
		DeleteTicketSegmentPrice(context context.Context, id int64) (err error)
	}
)

var (
	localTicketSegmentPriceService ITicketSegmentPriceService
)

func TicketSegmentPriceService() ITicketSegmentPriceService {
	if localTicketSegmentPriceService == nil {
		panic("implement localTicketSegmentPriceService not found for interface ITicketSegmentPriceService")
	}
	return localTicketSegmentPriceService
}
func InitTicketSegmentPriceService(i ITicketSegmentPriceService) {
	localTicketSegmentPriceService = i
}
