package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/service"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/vo"
	"github.com/quangdat385/holiday-ticket/ticket-service/response"
)

var TicketItemController = new(cTicketItemController)

type cTicketItemController struct {
}

// @Summary      Get ticket item by ticket id
// @Description  Get ticket item by ticket id
// @Tags         ticket item
// @Accept       json
// @Produce      json
// @Param        ticket_id path int true "ticket id"
// @Param        version query int true "version"
// @Success      200  {object}  response.ResponseData{data=model.TicketItemsOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /ticket-item/get-by-id/{ticket_id} [get]
func (c *cTicketItemController) GetTicketItem(ctx *gin.Context) {
	var params vo.TicketItemRequest

	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	var query vo.TicketItemQueryRequest
	if err := ctx.ShouldBindQuery(&query); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	// call implementation
	ticketItem, err := service.TicketItem().GetTicketItemById(ctx, params.TicketId, query.Version)
	if err != nil {
		response.ErrorResponse(ctx, response.FoundTicketErrCodeStatus, err.Error())
		return
	}
	if ticketItem.IsEmpty() { // Assuming IsEmpty() is a method to check if the struct is empty
		response.ErrorResponse(ctx, response.FoundTicketErrCodeStatus, "ticket not found")
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, ticketItem)
}

// @Summary      Create ticket item
// @Description  Create ticket item
// @Tags         ticket item
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        body body vo.TicketItemCreateRequest true "ticket item create request"
// @Success      200  {object}  response.ResponseData{data=model.TicketItemsOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /ticket-item/create [post]
func (c *cTicketItemController) CreateTicketItem(ctx *gin.Context) {
	var params vo.TicketItemCreateRequest

	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	// call implementation
	ticketItem, err := service.TicketItem().CreateTicketItem(ctx, model.TicketItemInPut{
		TicketName:      params.TicketName,
		Description:     params.Description,
		TicketId:        params.TicketId,
		TrainID:         params.TrainID,
		SeatClass:       params.SeatClass,
		StockInitial:    params.StockInitial,
		StockAvailable:  params.StockAvailable,
		DepartureTime:   params.DepartureTime,
		IsStockPrepared: params.IsStockPrepared,
		PriceOriginal:   params.PriceOriginal,
		PriceFlash:      params.PriceFlash,
		SaleStartTime:   params.SaleStartTime,
		SaleEndTime:     params.SaleEndTime,
		Status:          params.Status,
		ActivityId:      params.ActivityId,
	})
	if err != nil {
		response.ErrorResponse(ctx, response.CreateSuccessCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, ticketItem)
}

// @Summary      Update ticket item
// @Description  Update ticket item
// @Tags         ticket item
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        ticket_item_id path int true "ticket item id"
// @Param        body body vo.UpdateTicketItemRequest true "ticket item update request"
// @Success      200  {object}  response.ResponseData{data=model.TicketItemsOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /ticket-item/update/{ticket_item_id} [put]
func (c *cTicketItemController) UpdateTicketItem(ctx *gin.Context) {
	var params vo.UpdateTicketItemRequest
	var ID vo.TicketItemRequest

	if err := ctx.ShouldBindUri(&ID); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	if ID.TicketId != params.TicketItemId {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, "ticket item id in uri and body do not match")
		return
	}

	// call implementation
	ticketItem, err := service.TicketItem().UpdateTicketItem(ctx, model.UpdateTicketItemInPut{
		TicketItemId:    params.TicketItemId,
		TicketId:        params.TicketId,
		TicketName:      params.TicketName,
		Description:     params.Description,
		TrainID:         params.TrainID,
		SeatClass:       params.SeatClass,
		StockInitial:    params.StockInitial,
		StockAvailable:  params.StockAvailable,
		DepartureTime:   params.DepartureTime,
		IsStockPrepared: params.IsStockPrepared,
		PriceOriginal:   params.PriceOriginal,
		PriceFlash:      params.PriceFlash,
		SaleStartTime:   params.SaleStartTime,
		SaleEndTime:     params.SaleEndTime,
		Status:          params.Status,
		ActivityId:      params.ActivityId,
	})
	if err != nil {
		response.ErrorResponse(ctx, response.UpdateSuccessCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, ticketItem)
}

// @Summary      delete ticket item by id
// @Description  delete ticket item by id
// @Tags         ticket item
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        ticket_item_id path int true "ticket item id"
// @Success      200  {object}  response.ResponseData{data=model.TicketItemsOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /ticket-item/delete/{ticket_item_id} [delete]
func (c *cTicketItemController) DeleteTicketItem(ctx *gin.Context) {
	var params vo.TicketItemRequest

	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	// call implementation
	err := service.TicketItem().DeleteTicketItem(ctx, params.TicketId)
	if err != nil {
		response.ErrorResponse(ctx, response.DeleteSuccessCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, "ticket item deleted successfully")
}

// @Summary set ticket stock cache
// @Description set ticket stock cache
// @Tags ticket item
// @Accept json
// @Produce json
// @Param x-client-id header string true "Client ID"
// @Param x-device-id header string true "Device ID"
// @Param body body vo.SetStockCacheRequest true "set stock cache request"
// @Success 200 {object} response.ResponseData{data=string}
// @Failure 400 {object} response.ErrorResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /ticket-item/set-stock-cache [put]
func (c *cTicketItemController) SetStockCache(ctx *gin.Context) {
	var params vo.SetStockCacheRequest

	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	// call implementation
	_, err := service.TicketItem().SetStockCache(ctx, params.TicketItemId, params.Stock, params.Expiration)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCorsCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, "stock cache set successfully")
}

// @Summary decrease ticket cache
// @Description decrease ticket cache
// @Tags ticket item
// @Accept json
// @Produce json
// @Param x-client-id header string true "Client ID"
// @Param x-device-id header string true "Device ID"
// @Param body body vo.DecreaseStockRequest true "decrease stock cache request"
// @Success 200 {object} response.ResponseData{data=string}
// @Failure 400 {object} response.ErrorResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /ticket-item/decrease-stock-cache [put]
func (c *cTicketItemController) DecreaseStockCache(ctx *gin.Context) {
	var params vo.DecreaseStockRequest

	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	// call implementation
	_, num := service.TicketItem().DecreaseStock(ctx, params.TicketItemID, params.Stock)
	if num != 3 {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, num)
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, "stock cache decreased successfully")
}
