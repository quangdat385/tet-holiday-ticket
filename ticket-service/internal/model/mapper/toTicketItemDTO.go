package mapper

import (
	"strconv"

	"github.com/quangdat385/holiday-ticket/ticket-service/internal/database"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
)

func parsePriceOriginal(priceStr string) float32 {
	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		// Handle the error appropriately, e.g., log it or return a default value
		return 0.0
	}
	return float32(price)
}

func ToTicketItemDTO(ticketItem database.GetTicketItemByIdRow) model.TicketItemsOutput {
	return model.TicketItemsOutput{
		TicketId:        int(ticketItem.ID),
		TicketName:      ticketItem.Name,
		StockInitial:    int(ticketItem.StockInitial),
		StockAvailable:  int(ticketItem.StockAvailable),
		IsStockPrepared: ticketItem.IsStockPrepared,
		PriceOriginal:   parsePriceOriginal(ticketItem.PriceOriginal),
		PriceFlash:      parsePriceOriginal(ticketItem.PriceFlash),
		SaleStartTime:   ticketItem.SaleStartTime.UTC(),
		SaleEndTime:     ticketItem.SaleEndTime.UTC(),
		Status:          int(ticketItem.Status),
		ActivityId:      int64(ticketItem.ActivityID),
	}
}
