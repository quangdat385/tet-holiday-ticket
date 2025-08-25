package mapper

import (
	"encoding/json"
	"strconv"

	"github.com/quangdat385/holiday-ticket/order-service/internal/database"
	"github.com/quangdat385/holiday-ticket/order-service/internal/model"
)

func ToOrderDTO(in database.GetOrderByIdRow) (out model.OrderOutPut) {
	out.ID = in.ID
	out.OrderNumber = in.OrderNumber
	orderAmount, err := strconv.ParseFloat(in.OrderAmount, 32)
	if err != nil {
		out.OrderAmount = 0
	} else {
		out.OrderAmount = float32(orderAmount)
	}
	var orderItem model.OrderItem
	err = json.Unmarshal(in.OrderItem, &orderItem)
	if err != nil {
		orderItem = model.OrderItem{}
	}
	out.StationCode = in.StationCode
	out.UserID = in.UserID
	out.TerminalID = in.TerminalID
	out.OrderDate = in.OrderDate
	out.OrderNotes = in.OrderNotes
	out.OrderItem = orderItem
	out.UpdatedAt = in.UpdatedAt

	return out
}
