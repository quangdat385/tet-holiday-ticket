package controller

import (
	"fmt"
	"strconv"

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
// @Success      200  {object}  response.ResponseData{data=model.TicketItemsOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /ticket/ticket-item/{ticket_id} [get]
func (c *cTicketItemController) GetTicketItem(ctx *gin.Context) {
	var params vo.TicketItemRequest
	version := ctx.Query("version")

	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	var versionInt int
	fmt.Println("version: ", version)
	if version != "" {
		var err error
		versionInt, err = strconv.Atoi(version)
		if err != nil {
			response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
			return
		}
	} else {
		versionInt = 0
	}

	// call implementation
	ticketItem, err := service.TicketItem().GetTicketItemById(ctx, params.TicketId, versionInt)
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
// @Param        body body vo.TickertItemCreateRequest true "ticket item create request"
// @Success      200  {object}  response.ResponseData{data=model.TicketItemsOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /ticket/ticket-item/create [post]
func (c *cTicketItemController) CreateTicketItem(ctx *gin.Context) {
	var params vo.TickertItemCreateRequest

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
// @Router       /ticket/ticket-item/update/{ticket_item_id} [put]
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
// @Router       /ticket/ticket-item/delete/{ticket_item_id} [delete]
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
