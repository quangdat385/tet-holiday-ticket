package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/service"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/vo"
	"github.com/quangdat385/holiday-ticket/ticket-service/response"
)

var TicketSegmentPriceController = new(cTicketSegmentPriceController)

type cTicketSegmentPriceController struct {
}

// @Summary      Create ticket segment price
// @Description  Create ticket segment price
// @Tags         ticket-segment-price
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        payload  body      vo.CreateTicketSegmentPriceRequest  true  "Create ticket segment price request"
// @Success      200  {object}  response.ResponseData{data=model.TicketSegmentPriceOutPut}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /ticket-segment-price/create [post]
func (c *cTicketSegmentPriceController) CreateTicketSegmentPrice(ctx *gin.Context) {
	var params vo.CreateTicketSegmentPriceRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	ticketSegmentPrice, err := service.TicketSegmentPriceService().CreateTicketSegmentPrice(ctx, model.TicketSegmentPriceCreateInPut{
		TicketItemID:   int64(params.TicketItemID),
		RouteSegmentID: int64(params.RouteSegmentID),
		Price:          params.Price,
	})

	if err != nil {
		response.ErrorResponse(ctx, response.CreateErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.CreateSuccessCodeStatus, ticketSegmentPrice)
}

// @Summary      Update ticket segment price
// @Description  Update ticket segment price
// @Tags         ticket-segment-price
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        ticket_segment_price_id  path      int64  true  "Ticket Segment Price ID"
// @Param        payload  body      vo.UpdateTicketSegmentPriceRequest  true  "Update ticket segment price request"
// @Success      200  {object}  response.ResponseData{data=model.TicketSegmentPriceOutPut}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /ticket-segment-price/update/{ticket_segment_price_id} [patch]
func (c *cTicketSegmentPriceController) UpdateTicketSegmentPrice(ctx *gin.Context) {
	var params vo.UpdateTicketSegmentPriceRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	ticketSegmentPrice, err := service.TicketSegmentPriceService().UpdateTicketSegmentPrice(ctx, model.TicketSegmentPriceUpdateInPut{
		ID:             int64(params.TicketSegmentPriceID),
		TicketItemID:   int64(params.TicketItemID),
		RouteSegmentID: int64(params.RouteSegmentID),
		Price:          params.Price,
	})

	if err != nil {
		response.ErrorResponse(ctx, response.UpdateErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.UpdateSuccessCodeStatus, ticketSegmentPrice)
}

// @Summary      Delete ticket segment price
// @Description  Delete ticket segment price
// @Tags         ticket-segment-price
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        ticket_segment_price_id  path      int64  true  "Ticket Segment Price ID"
// @Success      200  {object}  response.ResponseData{data=bool}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /ticket-segment-price/delete/{ticket_segment_price_id} [delete]
func (c *cTicketSegmentPriceController) DeleteTicketSegmentPrice(ctx *gin.Context) {
	var params vo.TicketSegmentPriceIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	err := service.TicketSegmentPriceService().DeleteTicketSegmentPrice(ctx, int64(params.TicketSegmentPriceID))
	if err != nil {
		response.ErrorResponse(ctx, response.DeleteErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.DeleteSuccessCodeStatus, true)
}

// @Summary      Get ticket segment price by ID
// @Description  Get ticket segment price by ID
// @Tags         ticket-segment-price
// @Accept       json
// @Produce      json
// @Param        ticket_segment_price_id  path      int64  true  "Ticket Segment Price ID"
// @Success      200  {object}  response.ResponseData{data=model.TicketSegmentPriceOutPut}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /ticket-segment-price/get-one/{ticket_segment_price_id} [get]
func (c *cTicketSegmentPriceController) GetTicketSegmentPriceById(ctx *gin.Context) {
	var params vo.TicketSegmentPriceIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	ticketSegmentPrice, err := service.TicketSegmentPriceService().GetTicketSegmentPriceByRouteSegmentId(ctx, int64(params.TicketSegmentPriceID))
	if err != nil {
		response.ErrorResponse(ctx, response.NotFoundErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, ticketSegmentPrice)
}

// @Summary      Get all ticket segment price
// @Description  Get all ticket segment price
// @Tags         ticket-segment-price
// @Accept       json
// @Produce      json
// @Param        limit  query     int  true  "Limit"
// @Param        page   query     int  true  "Page"
// @Success      200  {object}  response.ResponseData{data=[]model.TicketSegmentPriceOutPut}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /ticket-segment-price/get-all [get]
func (c *cTicketSegmentPriceController) GetAllTicketSegmentPrice(ctx *gin.Context) {
	var params vo.GetListTicketSegmentPriceRequest
	if err := ctx.ShouldBindQuery(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	ticketSegmentPrice, err := service.TicketSegmentPriceService().GetAllTicketSegmentPrice(ctx, model.GetAllTicketSegmentPriceInPut{
		Page:  int64(params.Page),
		Limit: int64(params.Limit),
	})
	if err != nil {
		response.ErrorResponse(ctx, response.NotFoundErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, ticketSegmentPrice)
}

// @Summary      Get ticket segment price by order segment ID
// @Description  Get ticket segment price by order segment ID
// @Tags         ticket-segment-price
// @Accept       json
// @Produce      json
// @Param  route_segment_id  path      int64  true  "Route Segment ID"
// @Success      200  {object}  response.ResponseData{data=model.TicketSegmentPriceOutPut}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /ticket-segment-price/get-by-route-segment-id/{route_segment_id} [get]
func (c *cTicketSegmentPriceController) GetTicketSegmentPriceByRouteSegmentId(ctx *gin.Context) {
	var params vo.TicketSegmentPriceIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	ticketSegmentPrice, err := service.TicketSegmentPriceService().GetTicketSegmentPriceByRouteSegmentId(ctx, int64(params.TicketSegmentPriceID))
	if err != nil {
		response.ErrorResponse(ctx, response.NotFoundErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, ticketSegmentPrice)
}

// @Summary      Get all ticket segment price by from route segment ID to route segment ID
// @Description  Get all ticket segment price by from route segment ID to route segment ID
// @Tags         ticket-segment-price
// @Accept       json
// @Produce     json
// @Param       from_segment_id query int64  true  "FromSegmentID"
// @Param       to_segment_id   query int64  true  "ToSegmentID"
// @Success      200  {object}  response.ResponseData{data=[]model.TicketSegmentPriceOutPut}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /ticket-segment-price/get-by-from-to-route-segment-id [get]
func (c *cTicketSegmentPriceController) GetAllTicketSegmentPriceByFromToRouteSegmentId(ctx *gin.Context) {
	var params vo.GetTicketSegmentPriceFromRouteSegmentIDToToSegmentIDRequest
	if err := ctx.ShouldBindQuery(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	ticketSegmentPrice, err := service.TicketSegmentPriceService().GetAllTicketSegmentPriceFromSegmentIDToToSegmentID(ctx, model.TicketSegmentPriceListInPut{
		FromSegmentID: int64(params.FromSegmentID),
		ToSegmentID:   int64(params.ToSegmentID),
	})

	if err != nil {
		response.ErrorResponse(ctx, response.NotFoundErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, ticketSegmentPrice)
}
