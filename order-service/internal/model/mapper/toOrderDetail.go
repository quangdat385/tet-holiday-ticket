package mapper

import (
	"strconv"

	"github.com/quangdat385/holiday-ticket/order-service/internal/database"
	"github.com/quangdat385/holiday-ticket/order-service/internal/model"
)

func ToOrderDetailDTO(order database.GetOrderDetailByIdRow) (out model.OrderDetailOutput) {
	out.ID = int64(order.ID)
	out.TicketItemID = order.TicketItemID
	out.OrderNumber = order.OrderNumber
	out.PassengerName = order.PassengerName
	out.DepartureStation = order.DepartureStation
	out.ArrivalStation = order.ArrivalStation
	out.DepartureTime = order.DepartureTime
	// Convert TicketPrice from string to float64
	ticketPrice, err := strconv.ParseFloat(order.TicketPrice, 64)
	if err != nil {
		ticketPrice = 0 // or handle error as needed
	}
	out.TicketPrice = ticketPrice
	out.SeatClass = string(order.SeatClass)
	out.SeatNumber = order.SeatNumber
	out.CreatedAt = order.CreatedAt
	out.UpdatedAt = order.UpdatedAt

	return out
}
