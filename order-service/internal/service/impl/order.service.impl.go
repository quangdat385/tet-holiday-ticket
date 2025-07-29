package impl

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/quangdat385/holiday-ticket/order-service/internal/database"
	"github.com/quangdat385/holiday-ticket/order-service/internal/model"
	"github.com/quangdat385/holiday-ticket/order-service/internal/model/mapper"
	"github.com/quangdat385/holiday-ticket/order-service/internal/vo"
	"github.com/quangdat385/holiday-ticket/order-service/response"
)

type sOrderService struct {
	r *database.Queries
}

func NewOrderServiceImpl(r *database.Queries) *sOrderService {
	return &sOrderService{
		r: r,
	}
}
func (s *sOrderService) GetOrderByID(ctx context.Context, orderID int64) (out model.OrderOutPut, err error) {
	order, err := s.r.GetOrderById(ctx, int32(orderID))
	if err != nil {
		return out, err
	}
	if order == (database.GetOrderByIdRow{}) {
		return out, response.ErrNotFoundDataErr
	}
	out = mapper.ToOrderDTO(order)
	return out, nil
}

func (s *sOrderService) CreateOrder(ctx context.Context, in vo.CreateOrderRequest) (out model.OrderOutPut, err error) {
	result, err := s.r.InsertOrder(ctx, database.InsertOrderParams{
		OrderNumber: in.OrderNUmber,
		OrderAmount: strconv.FormatFloat(float64(in.OrderAmount), 'f', -1, 32),
		TerminalID:  in.TerminalID,
		OrderDate:   in.OrderDate,
		OrderNotes:  sql.NullString{String: "Order-->Pending", Valid: true},
	})
	if err != nil {
		return out, err
	}
	lastOrderId, err := result.LastInsertId()
	if err != nil {
		return out, err
	}
	order, err := s.r.GetOrderById(ctx, int32(lastOrderId))
	if err != nil {
		return out, err
	}
	out = mapper.ToOrderDTO(order)
	return out, nil
}
func (s *sOrderService) UpdateOrder(ctx context.Context, in vo.UpdateOrderRequest) (out model.OrderOutPut, err error) {
	orderDB, err := s.r.GetOrderById(ctx, int32(in.OrderId))
	if err != nil {
		return out, err
	}
	if orderDB == (database.GetOrderByIdRow{}) {
		return out, response.ErrNotFoundDataErr
	}
	if in.OrderAmount != 0 {
		orderDB.OrderAmount = strconv.FormatFloat(float64(in.OrderAmount), 'f', -1, 32)
	}
	if in.TerminalID != 0 {
		orderDB.TerminalID = in.TerminalID
	}
	if !in.OrderDate.IsZero() {
		orderDB.OrderDate = in.OrderDate
	}
	if in.OrderNotes != "" {
		orderDB.OrderNotes = sql.NullString{String: in.OrderNotes, Valid: true}
	}
	result, err := s.r.UpdateOrder(ctx, database.UpdateOrderParams{
		ID:          orderDB.ID,
		OrderAmount: orderDB.OrderAmount,
		TerminalID:  orderDB.TerminalID,
		OrderDate:   orderDB.OrderDate,
		OrderNotes:  orderDB.OrderNotes,
	})
	if err != nil {
		return out, err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return out, err
	}
	out = mapper.ToOrderDTO(orderDB)
	return out, nil
}
func (s *sOrderService) DeleteOrder(ctx context.Context, orderID int64) (err error) {
	orderDB, err := s.r.GetOrderById(ctx, int32(orderID))
	if err != nil {
		return err
	}
	if orderDB == (database.GetOrderByIdRow{}) {
		return response.ErrNotFoundDataErr
	}
	_, err = s.r.DeleteOrder(ctx, int32(orderID))
	if err != nil {
		return err
	}
	return nil
}
