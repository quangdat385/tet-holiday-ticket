package service

import (
	"context"

	"github.com/quangdat385/holiday-ticket/order-service/internal/model"
	"github.com/quangdat385/holiday-ticket/order-service/internal/vo"
)

type (
	IOrderService interface {
		GetOrderByID(ctx context.Context, orderID int64) (out model.OrderOutPut, err error)
		CreateOrder(ctx context.Context, in vo.CreateOrderRequest) (out model.OrderOutPut, err error)
		UpdateOrder(ctx context.Context, in vo.UpdateOrderRequest) (out model.OrderOutPut, err error)
		DeleteOrder(ctx context.Context, orderID int64) (err error)
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
