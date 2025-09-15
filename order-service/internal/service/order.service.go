package service

import (
	"context"

	"github.com/quangdat385/holiday-ticket/order-service/internal/model"
	"github.com/quangdat385/holiday-ticket/order-service/internal/vo"
)

type (
	IOrderService interface {
		GetOrderByID(ctx context.Context, orderID int64) (out model.OrderOutPut, err error)
		CreateOrder(ctx context.Context, in model.OrderInput) (out bool, err error)
		UpdateOrder(ctx context.Context, in vo.UpdateOrderRequest) (out model.OrderOutPut, err error)
		UpdateOrderNotes(ctx context.Context, orderNumber string, orderNotes string) (err error)
		DeleteOrder(ctx context.Context, orderID int64) (err error)
		GetOrderByOrderNumber(ctx context.Context, orderNumber string) (out model.OrderOutPut, err error)
		GetOrdersByUserId(ctx context.Context, userID int64, page int32, pageSize int32) (out []model.OrderOutPut, err error)
	}
)

var (
	localOrderService IOrderService
)

func OrderService() IOrderService {
	if localOrderService == nil {
		panic("implement localOrderService not found for interface IOrderService")
	}
	return localOrderService
}
func InitOrderService(i IOrderService) {
	localOrderService = i
}
