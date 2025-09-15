package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/order-service/internal/service"
	"github.com/quangdat385/holiday-ticket/order-service/internal/vo"
	"github.com/quangdat385/holiday-ticket/order-service/response"
)

var OrderDetailController = new(cOrderDetailController)

type cOrderDetailController struct{}

// @Summary Get order detail by ID
// @Description Get order detail by ID
// @Tags order-detail
// @Accept  json
// @Produce  json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param id path string true "Order Detail ID"
// @Success 200 {object} response.ResponseData{data=model.OrderDetailOutput}
// @Failure 400 {object} response.ErrorResponseData
// @Failure 404 {object} response.ErrorResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /order/detail/get-by-id/{id} [get]
func (c *cOrderDetailController) GetOrderDetailByID(ctx *gin.Context) {
	var params vo.OrderDetailIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	orderDetail, err := service.OrderDetailService().GetOrderDetailByID(ctx, int32(params.ID))
	if err != nil {
		response.ErrorResponse(ctx, response.NotFoundErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, orderDetail)
}

// @Summary Create order detail
// @Description Create a new order detail
// @Tags order-detail
// @Accept  json
// @Produce  json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param payload body vo.CreateOrderDetailRequest true "Create order detail request"
// @Success 200 {object} response.ResponseData{data=model.OrderDetailOutput}
// @Failure 400 {object} response.ErrorResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /order/detail/create [post]
func (c *cOrderDetailController) CreateOrderDetail(ctx *gin.Context) {
	var params vo.CreateOrderDetailRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	orderDetail, err := service.OrderDetailService().CreateOrderDetail(ctx, params)
	if err != nil {
		response.ErrorResponse(ctx, response.CreateErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.CreateSuccessCodeStatus, orderDetail)
}

// @Summary Update order detail
// @Description Update an existing order detail
// @Tags order-detail
// @Accept  json
// @Produce  json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param id path string true "Order Detail ID"
// @Param payload body vo.UpdateOrderDetailRequest true "Update order detail request"
// @Success 200 {object} response.ResponseData{data=model.OrderDetailOutput}
// @Failure 400 {object} response.ErrorResponseData
// @Failure 404 {object} response.ErrorResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /order/detail/update/{id} [put]
func (c *cOrderDetailController) UpdateOrderDetail(ctx *gin.Context) {
	var params vo.UpdateOrderDetailRequest
	var uriParams vo.OrderDetailIDRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	if err := ctx.ShouldBindUri(&uriParams); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	orderDetail, err := service.OrderDetailService().UpdateOrderDetail(ctx, params)
	if err != nil {
		response.ErrorResponse(ctx, response.UpdateErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.UpdateSuccessCodeStatus, orderDetail)
}

// @Summary Delete order detail
// @Description Delete an order detail by ID
// @Tags order-detail
// @Accept  json
// @Produce  json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param id path string true "Order Detail ID"
// @Success 200 {object} response.ResponseData{data=string}
// @Failure 400 {object} response.ErrorResponseData
// @Failure 404 {object} response.ErrorResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /order/detail/delete/{id} [delete]
func (c *cOrderDetailController) DeleteOrderDetail(ctx *gin.Context) {
	var params vo.OrderDetailIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	err := service.OrderDetailService().DeleteOrderDetail(ctx, int32(params.ID))
	if err != nil {
		response.ErrorResponse(ctx, response.DeleteErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.DeleteSuccessCodeStatus, "Order detail deleted successfully")
}
