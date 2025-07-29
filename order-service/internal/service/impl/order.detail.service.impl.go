package impl

import (
	"context"
	"fmt"

	"github.com/quangdat385/holiday-ticket/order-service/internal/database"
	"github.com/quangdat385/holiday-ticket/order-service/internal/model"
	"github.com/quangdat385/holiday-ticket/order-service/internal/model/mapper"
	"github.com/quangdat385/holiday-ticket/order-service/internal/vo"
	"github.com/quangdat385/holiday-ticket/order-service/response"
)

type sOrderDetailService struct {
	r *database.Queries
}

func NewOrderDetailServiceImpl(r *database.Queries) *sOrderDetailService {
	return &sOrderDetailService{
		r: r,
	}
}

func (s *sOrderDetailService) GetOrderDetailByID(ctx context.Context, orderDetailId int32) (out model.OrderDetailOutput, err error) {
	orderDetail, err := s.r.GetOrderDetailById(ctx, orderDetailId)
	if err != nil {
		return out, err
	}
	if (orderDetail == database.GetOrderDetailByIdRow{}) {
		return out, response.ErrNotFoundDataErr
	}
	out = mapper.ToOrderDetailDTO(orderDetail)
	return out, nil
}
func (s *sOrderDetailService) CreateOrderDetail(ctx context.Context, in vo.CreateOrderDetailRequest) (out model.OrderDetailOutput, err error) {
	result, err := s.r.InsertOrderDetail(ctx, database.InsertOrderDetailParams{
		TicketItemID:     in.TicketItemID,
		OrderNumber:      in.OrderNumber,
		PassengerName:    in.PassengerName,
		DepartureStation: in.DepartureStation,
		ArrivalStation:   in.ArrivalStation,
		DepartureTime:    in.DepartureTime,
		TicketPrice:      fmt.Sprintf("%f", in.TicketPrice),
		SeatClass:        database.PreGoTicketOrderDetail05202599999SeatClass(in.SeatClass),
		SeatNumber:       in.SeatNumber,
	})
	if err != nil {
		return out, err
	}
	orderDetailId, err := result.LastInsertId()
	if err != nil {
		return out, err
	}
	orderDetail, err := s.r.GetOrderDetailById(ctx, int32(orderDetailId))
	if err != nil {
		return out, err
	}
	if (orderDetail == database.GetOrderDetailByIdRow{}) {
		return out, response.ErrNotFoundDataErr
	}
	out = mapper.ToOrderDetailDTO(orderDetail)
	return out, nil
}
func (s *sOrderDetailService) UpdateOrderDetail(ctx context.Context, in vo.UpdateOrderDetailRequest) (out model.OrderDetailOutput, err error) {
	orderDetailDB, err := s.r.GetOrderDetailById(ctx, int32(in.OrderDetailID))
	if err != nil {
		return out, err
	}
	if (orderDetailDB == database.GetOrderDetailByIdRow{}) {
		return out, response.ErrNotFoundDataErr
	}
	if in.TicketItemID != 0 {
		orderDetailDB.TicketItemID = in.TicketItemID
	}
	if in.OrderNumber != "" {
		orderDetailDB.OrderNumber = in.OrderNumber
	}
	if in.PassengerName != "" {
		orderDetailDB.PassengerName = in.PassengerName
	}
	if in.DepartureStation != "" {
		orderDetailDB.DepartureStation = in.DepartureStation
	}
	if in.ArrivalStation != "" {
		orderDetailDB.ArrivalStation = in.ArrivalStation
	}
	if !in.DepartureTime.IsZero() {
		orderDetailDB.DepartureTime = in.DepartureTime
	}
	if in.TicketPrice != 0 {
		orderDetailDB.TicketPrice = fmt.Sprintf("%f", in.TicketPrice)
	}
	if in.SeatClass != "" {
		orderDetailDB.SeatClass = database.PreGoTicketOrderDetail05202599999SeatClass(in.SeatClass)
	}
	if in.SeatNumber != "" {
		orderDetailDB.SeatNumber = in.SeatNumber
	}

	result, err := s.r.UpdateOrderDetail(ctx, database.UpdateOrderDetailParams{
		ID:               orderDetailDB.ID,
		TicketItemID:     orderDetailDB.TicketItemID,
		OrderNumber:      orderDetailDB.OrderNumber,
		PassengerName:    orderDetailDB.PassengerName,
		DepartureStation: orderDetailDB.DepartureStation,
		ArrivalStation:   orderDetailDB.ArrivalStation,
		DepartureTime:    orderDetailDB.DepartureTime,
		TicketPrice:      orderDetailDB.TicketPrice,
		SeatClass:        orderDetailDB.SeatClass,
		SeatNumber:       in.SeatNumber,
	})
	if err != nil {
		return out, err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return out, err
	}
	out = mapper.ToOrderDetailDTO(orderDetailDB)
	return out, nil
}
func (s *sOrderDetailService) DeleteOrderDetail(ctx context.Context, orderDetailId int32) (err error) {
	orderDetail, err := s.r.GetOrderDetailById(ctx, orderDetailId)
	if err != nil {
		return err
	}
	if (orderDetail == database.GetOrderDetailByIdRow{}) {
		return response.ErrNotFoundDataErr
	}
	_, err = s.r.DeleteOrderDetail(ctx, orderDetailId)
	if err != nil {
		return err
	}
	return nil
}
