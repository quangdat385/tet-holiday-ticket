package service

import (
	"context"

	"github.com/quangdat385/holiday-ticket/order-service/internal/model"
	"github.com/quangdat385/holiday-ticket/order-service/internal/vo"
)

type (
	IOrderDetailService interface {
		GetOrderDetailByID(ctx context.Context, orderDetailId int32) (out model.OrderDetailOutput, err error)
		CreateOrderDetail(ctx context.Context, in vo.CreateOrderDetailRequest) (out model.OrderDetailOutput, err error)
		UpdateOrderDetail(ctx context.Context, in vo.UpdateOrderDetailRequest) (out model.OrderDetailOutput, err error)
		DeleteOrderDetail(ctx context.Context, orderDetailId int32) (err error)
	}
)

var (
	localOrderDetailService IOrderDetailService
)

func OrderDetailService() IOrderDetailService {
	if localOrderDetailService == nil {
		panic("implement localOrderDetailService not found for interface IOrderDetailService")
	}
	return localOrderDetailService
}
func InitOrderDetailService(i IOrderDetailService) {
	localOrderDetailService = i
}
