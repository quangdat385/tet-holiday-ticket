package impl

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/quangdat385/holiday-ticket/order-service/global"
	"github.com/quangdat385/holiday-ticket/order-service/internal/database"
	"github.com/quangdat385/holiday-ticket/order-service/internal/model"
	"github.com/quangdat385/holiday-ticket/order-service/internal/model/mapper"
	"github.com/quangdat385/holiday-ticket/order-service/internal/vo"
	"github.com/quangdat385/holiday-ticket/order-service/response"
	"github.com/segmentio/kafka-go"
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
	if order.ID == 0 {
		return out, response.ErrNotFoundDataErr
	}
	out = mapper.ToOrderDTO(order)
	return out, nil
}
func (s *sOrderService) GetOrderByOrderNumber(ctx context.Context, orderNumber string) (out model.OrderOutPut, err error) {
	order, err := s.r.GetOrderByOrderNumber(ctx, orderNumber)
	if err != nil {
		return out, err
	}
	if order.ID == 0 {
		return out, response.ErrNotFoundDataErr
	}
	out = mapper.ToOrderDTO(database.GetOrderByIdRow{
		ID:          order.ID,
		StationCode: order.StationCode,
		UserID:      order.UserID,
		OrderNumber: order.OrderNumber,
		OrderAmount: order.OrderAmount,
		TerminalID:  order.TerminalID,
		OrderDate:   order.OrderDate,
		OrderNotes:  order.OrderNotes,
		OrderItem:   order.OrderItem,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
	})
	return out, nil
}
func (s *sOrderService) GetOrdersByUserId(ctx context.Context, userID int64, page int32, Limit int32) (out []model.OrderOutPut, err error) {
	orders, err := s.r.GetOrdersByUserId(ctx, database.GetOrdersByUserIdParams{
		UserID: userID,
		Limit:  Limit,
		Offset: (page - 1) * Limit,
	})
	if err != nil {
		return out, err
	}
	for _, order := range orders {
		out = append(out, mapper.ToOrderDTO(database.GetOrderByIdRow{
			ID:          order.ID,
			StationCode: order.StationCode,
			UserID:      order.UserID,
			OrderNumber: order.OrderNumber,
			OrderAmount: order.OrderAmount,
			TerminalID:  order.TerminalID,
			OrderDate:   order.OrderDate,
			OrderNotes:  order.OrderNotes,
			OrderItem:   order.OrderItem,
			CreatedAt:   order.CreatedAt,
			UpdatedAt:   order.UpdatedAt,
		}))
	}
	return out, nil
}
func (s *sOrderService) CreateOrder(ctx context.Context, in model.OrderInput) (out bool, err error) {
	before := in.OrderDate.Add(-10 * time.Second)
	checkIdempotent, err := s.r.CheckIdempotencyOrder(ctx, database.CheckIdempotencyOrderParams{
		UserID:      in.UserID,
		OrderDate:   before,
		OrderDate_2: in.OrderDate,
	})
	if err != nil {
		return out, err
	}
	if len(checkIdempotent) > 0 {
		return out, response.ErrDuplicateDataErr
	}
	var eventOrder model.OrderEvent
	eventOrder.Type = model.OrderEventTypeCreateOrder
	eventOrder.Order = model.OrderEventInput{
		StationCode: in.StationCode,
		UserID:      in.UserID,
		OrderNumber: in.OrderNumber,
		OrderAmount: in.OrderAmount,
		TerminalID:  in.TerminalID,
		OrderDate:   in.OrderDate,
		OrderNotes:  "Order-->Pending",
		OrderItem:   in.OrderItem,
	}
	orderEventJSON, err := json.Marshal(eventOrder)
	if err != nil {
		return out, err
	}
	key := model.OrderEventTypeCreateOrder
	msg := kafka.Message{
		Key:   []byte(key),
		Value: orderEventJSON,
		Time:  time.Now(),
	}
	err = global.KafkaProducer.WriteMessages(ctx, msg)
	if err != nil {
		return out, err
	}
	out = true
	return out, nil
}
func (s *sOrderService) UpdateOrder(ctx context.Context, in vo.UpdateOrderRequest) (out model.OrderOutPut, err error) {
	orderDB, err := s.r.GetOrderById(ctx, int32(in.OrderId))
	if err != nil {
		return out, err
	}
	if orderDB.ID == 0 {
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
		orderDB.OrderNotes = in.OrderNotes
	}
	if in.StationCode != "" {
		orderDB.StationCode = in.StationCode
	}
	result, err := s.r.UpdateOrder(ctx, database.UpdateOrderParams{
		ID:          orderDB.ID,
		OrderAmount: orderDB.OrderAmount,
		TerminalID:  orderDB.TerminalID,
		OrderDate:   orderDB.OrderDate,
		OrderNotes:  orderDB.OrderNotes,
		StationCode: orderDB.StationCode,
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
func (s *sOrderService) UpdateOrderNotes(ctx context.Context, orderNumber string, orderNotes string) (err error) {
	orderDB, err := s.r.GetOrderByOrderNumber(ctx, orderNumber)
	if err != nil {
		return err
	}
	if orderDB.ID == 0 {
		return response.ErrNotFoundDataErr
	}
	_, err = s.r.UpdateOrderNote(ctx, database.UpdateOrderNoteParams{
		OrderNumber: orderNumber,
		OrderNotes:  orderNotes,
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *sOrderService) DeleteOrder(ctx context.Context, orderID int64) (err error) {
	orderDB, err := s.r.GetOrderById(ctx, int32(orderID))
	if err != nil {
		return err
	}
	if orderDB.ID == 0 {
		return response.ErrNotFoundDataErr
	}
	_, err = s.r.DeleteOrder(ctx, int32(orderID))
	if err != nil {
		return err
	}
	return nil
}
