package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/service"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/vo"
	"github.com/quangdat385/holiday-ticket/ticket-service/response"
)

var TicketController = new(cTicketController)

type cTicketController struct {
}

// @Summary      Create ticket
// @Description  Create ticket
// @Tags         ticket
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        payload  body      vo.CreateTicketRequest  true  "Create ticket request"
// @Success      200  {object}  response.ResponseData{data=model.TicketOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /ticket/create [post]
func (c *cTicketController) CreateTicket(ctx *gin.Context) {
	var params vo.CreateTicketRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	ticket, err := service.TicketHome().CreateTicket(ctx, model.CreateTicketInput{
		Name:        params.Name,
		Description: params.Description,
		StartTime:   params.StartTime,
		EndTime:     params.EndTime,
		Status:      params.Status,
	})
	if err != nil {
		response.ErrorResponse(ctx, response.CreateErrorCodeStatus, err.Error())
		return
	}
	if ticket.ID == 0 { // Assuming IsEmpty() is a method to check if the struct is empty
		response.ErrorResponse(ctx, response.CreateErrorCodeStatus, "create ticket failed")
		return
	}
	response.SuccessResponse(ctx, response.CreateSuccessCodeStatus, ticket)
}

// @Summary      Get ticket By ID
// @Description  Get ticket By ID
// @Tags         ticket
// @Accept       json
// @Produce      json
// @Param        ticket_id  path      string  true  "ticket_id"
// @Success      200  {object}  response.ResponseData{data=model.TicketOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /ticket/{ticket_id} [get]
func (c *cTicketController) GetTicketById(ctx *gin.Context) {
	var params vo.TicketIdRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	ticket, err := service.TicketHome().GetTicketById(ctx, params.TicketId)
	if err != nil {
		response.ErrorResponse(ctx, response.FoundTicketErrCodeStatus, err.Error())
		return
	}
	if ticket.ID == 0 { // Assuming IsEmpty() is a method to check if the struct is empty
		response.ErrorResponse(ctx, response.FoundTicketErrCodeStatus, "ticket not found")
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, ticket)
}

// @Summary      Update ticket
// @Description  Update ticket
// @Tags         ticket
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        ticket_id  path      string  true  "ticket_id"
// @Param        payload    body      vo.UpdateTicketRequest  true  "Update ticket request"
// @Success      200  {object}  response.ResponseData{data=model.TicketOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /ticket/update/{ticket_id} [patch]
func (c *cTicketController) UpdateTicket(ctx *gin.Context) {
	var params vo.UpdateTicketRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	var ticketIdParams vo.TicketIdRequest
	if err := ctx.ShouldBindUri(&ticketIdParams); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	if ticketIdParams.TicketId != params.TicketId {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, "ticket_id not match")
		return
	}
	ticket, err := service.TicketHome().UpdateTicket(ctx, model.UpdateTicketInput{
		ID:          int64(ticketIdParams.TicketId),
		Name:        params.Name,
		Description: params.Description,
		Status:      params.Status,
	})
	if err != nil {
		response.ErrorResponse(ctx, response.UpdateErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.UpdateSuccessCodeStatus, ticket)
}

// @Summary      Get all tickets
// @Description  Get all tickets
// @Tags         ticket
// @Accept       json
// @Produce      json
// @Param        limit  query     int  true  "Limit"
// @Param        page   query     int  true  "Page"
// @Param        status query     int  true  "Status"
// @Success      200  {object}  response.ResponseData{data=[]model.TicketOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /ticket/get-all-ticket [get]
func (c *cTicketController) GetAllTicket(ctx *gin.Context) {
	var params vo.GetAllTicketsRequest
	if err := ctx.ShouldBindQuery(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	if params.Page <= 0 {
		params.Page = 1
	}
	fmt.Printf("params: %+v\n", params)
	tickets, err := service.TicketHome().GetAllTickets(ctx, model.GetAllTicketsInput{
		Page:   int32(params.Page),
		Limit:  int32(params.Limit),
		Status: int32(params.Status),
	})
	if err != nil {
		response.ErrorResponse(ctx, response.FoundTicketErrCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, tickets)
}

// @Summary      Delete ticket
// @Description  Delete ticket
// @Tags         ticket
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        ticket_id  path      string  true  "ticket_id"
// @Success      200  {object}  response.ResponseData{data=nil}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /ticket/delete/{ticket_id} [delete]
func (c *cTicketController) DeleteTicket(ctx *gin.Context) {
	var params vo.TicketIdRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	err := service.TicketHome().DeleteTicket(ctx, params.TicketId)
	if err != nil {
		response.ErrorResponse(ctx, response.DeleteErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.DeleteSuccessCodeStatus, nil)
}
