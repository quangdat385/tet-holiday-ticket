package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/order-service/internal/model"
	"github.com/quangdat385/holiday-ticket/order-service/internal/service"
	"github.com/quangdat385/holiday-ticket/order-service/internal/vo"
	"github.com/quangdat385/holiday-ticket/order-service/response"
)

var OrderController = new(cOrderController)

type cOrderController struct{}

// OrderController handles order-related operations.

// @Summary      Create order
// @Description  Create a new order
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        payload  body      vo.CreateOrderRequest  true  "Create order request"
// @Success      200  {object}  response.ResponseData{data=model.OrderOutPut}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /order/create [post]
func (c *cOrderController) CreateOrder(ctx *gin.Context) {
	var params vo.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	order, err := service.OrderService().CreateOrder(ctx, model.OrderInput{
		OrderNumber: params.OrderNUmber,
		UserID:      params.UserID,
		OrderAmount: params.OrderAmount,
		TerminalID:  params.TerminalID,
		OrderDate:   params.OrderDate,
		OrderItem:   params.OrderItem,
	})
	if err != nil {
		response.ErrorResponse(ctx, response.CreateErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.CreateSuccessCodeStatus, order)
}

// @Summary      Get order by ID
// @Description  Retrieve an order by its ID
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        id path string true "Order ID"
// @Success      200  {object}  response.ResponseData{data=model.OrderOutPut}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      404  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /order/{id} [get]
func (c *cOrderController) GetOrderByID(ctx *gin.Context) {
	var params vo.OrderIdRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	order, err := service.OrderService().GetOrderByID(ctx, int64(params.OrderId))
	if err != nil {
		response.ErrorResponse(ctx, response.NotFoundErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, order)
}

// @Summary      update order
// @Description  Update an existing order
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        payload  body      vo.UpdateOrderRequest  true  "Update order request"
// @Success      200  {object}  response.ResponseData{data=model.OrderOutPut}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      404  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /order/update [put]
func (c *cOrderController) UpdateOrder(ctx *gin.Context) {
	var params vo.UpdateOrderRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	order, err := service.OrderService().UpdateOrder(ctx, params)
	if err != nil {
		response.ErrorResponse(ctx, response.UpdateErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.UpdateSuccessCodeStatus, order)
}

// @Summary      Delete order
// @Description  Delete an order by its ID
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        id path string true "Order ID"
// @Success      200  {object}  response.ResponseData{data=string}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      404  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /order/delete/{id} [delete]
func (c *cOrderController) DeleteOrder(ctx *gin.Context) {
	var params vo.OrderIdRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	err := service.OrderService().DeleteOrder(ctx, int64(params.OrderId))
	if err != nil {
		response.ErrorResponse(ctx, response.NotFoundErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.DeleteSuccessCodeStatus, "Order deleted successfully")
}
