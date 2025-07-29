package mapper

import (
	"strconv"

	"github.com/quangdat385/holiday-ticket/ticket-service/internal/database"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
)

func ToTicketSegmentPriceDTO(ticketSegmentPriceRow database.GetTicketSegmentPriceByIdRow) model.TicketSegmentPriceOutPut {
	price, err := strconv.ParseFloat(ticketSegmentPriceRow.Price, 64)
	if err != nil {
		price = 0 // or handle error as needed
	}
	return model.TicketSegmentPriceOutPut{
		ID:             ticketSegmentPriceRow.ID,
		TicketItemID:   ticketSegmentPriceRow.TicketItemID,
		RouteSegmentID: ticketSegmentPriceRow.RouteSegmentID,
		Price:          price,
	}
}
