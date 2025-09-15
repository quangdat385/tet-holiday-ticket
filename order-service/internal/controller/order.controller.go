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
		StationCode: params.StationCode,
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
// @Param        order_id path string true "Order ID"
// @Success      200  {object}  response.ResponseData{data=model.OrderOutPut}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      404  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /order/get-by-id/{order_id} [get]
func (c *cOrderController) GetOrderByID(ctx *gin.Context) {
	var params vo.OrderIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	order, err := service.OrderService().GetOrderByID(ctx, int64(params.OrderID))
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
// @Param        order_id path string true "Order ID"
// @Param        payload  body      vo.UpdateOrderRequest  true  "Update order request"
// @Success      200  {object}  response.ResponseData{data=model.OrderOutPut}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      404  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /order/update/{order_id} [put]
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
	var params vo.OrderIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	err := service.OrderService().DeleteOrder(ctx, int64(params.OrderID))
	if err != nil {
		response.ErrorResponse(ctx, response.NotFoundErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.DeleteSuccessCodeStatus, "Order deleted successfully")
}

// @Summary      Get orders by order number
// @Description  Retrieve orders by their order number
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        order_number path string true "Order Number"
// @Success      200  {object}  response.ResponseData{data=[]model.OrderOutPut}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      404  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /order/get-by-order-number/{order_number} [get]
func (c *cOrderController) GetOrderByOrderNumber(ctx *gin.Context) {
	var params vo.OrderNumberRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	order, err := service.OrderService().GetOrderByOrderNumber(ctx, params.OrderNumber)
	if err != nil {
		response.ErrorResponse(ctx, response.NotFoundErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, order)
}

// @Summary      Get orders by user ID
// @Description  Retrieve orders by their user ID
// @Tags         order
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        page query int false "Page number" default(1)
// @Param        page_size query int false "Number of items per page" default(50)
// @Param        user_id path string true "User ID"
// @Success      200  {object}  response.ResponseData{data=[]model.OrderOutPut}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      404  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /order/get-by-user-id/{user_id} [get]
func (c *cOrderController) GetOrdersByUserID(ctx *gin.Context) {
	var params vo.UserIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	var queries vo.QueriesOrderRequest
	if err := ctx.ShouldBindQuery(&queries); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	orders, err := service.OrderService().GetOrdersByUserId(ctx, params.UserID, queries.Page, queries.Limit)
	if err != nil {
		response.ErrorResponse(ctx, response.NotFoundErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, orders)
}
