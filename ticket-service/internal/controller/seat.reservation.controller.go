package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/service"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/vo"
	"github.com/quangdat385/holiday-ticket/ticket-service/response"
)

var SeatReservationController = new(cSeatReservationController)

type cSeatReservationController struct {
}

// @Summary Reserve a seat
// @Description Reserve a seat
// @Tags SeatReservation
// @Accept json
// @Produce json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param payload body vo.CreateSeatReservationRequest true "Create seat reservation request"
// @Success 200 {object} response.ResponseData{data=model.SeatReservationOutput}
// @Failure 400 {object} response.ErrorResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /seat-reservation/create [post]
func (c *cSeatReservationController) ReserveSeat(ctx *gin.Context) {
	// Bind the request body to the SeatReservation struct
	var seatReservation vo.CreateSeatReservationRequest
	if err := ctx.ShouldBindJSON(&seatReservation); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	// Call the service to reserve the seat
	result, err := service.SeatReservationService().CreateSeatReservation(ctx, model.SeatReservationCreateInput{
		SeatID:        int64(seatReservation.SeatID),
		TrainID:       int64(seatReservation.TrainID),
		OrderNumber:   seatReservation.OrderNumber,
		FromStationID: int64(seatReservation.FromStationID),
		ToStationID:   int64(seatReservation.ToStationID),
	})
	if err != nil {
		response.ErrorResponse(ctx, response.CreateErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.CreateSuccessCodeStatus, result)
}

// @Summary Update a seat reservation
// @Description Update a seat reservation
// @Tags SeatReservation
// @Accept json
// @Produce json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param seat_reservation_id path int true "Seat Reservation ID"
// @Param payload body vo.UpdateSeatReservationRequest true "Update seat reservation request"
// @Success 200 {object} response.ResponseData{data=model.SeatReservationOutput}
// @Failure 400 {object} response.ErrorResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /seat-reservation/update/{seat_reservation_id} [patch]
func (c *cSeatReservationController) UpdateSeatReservation(ctx *gin.Context) {
	// Get the seat reservation ID from the URL parameter
	var seatReservationIDRequest vo.SeatReservationIDRequest
	if err := ctx.ShouldBindUri(&seatReservationIDRequest); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	// Bind the request body to the SeatReservation struct
	var seatReservation vo.UpdateSeatReservationRequest
	if err := ctx.ShouldBindJSON(&seatReservation); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	if seatReservationIDRequest.SeatReservationID != seatReservation.SeatReservationID {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, "seat_reservation_id not match")
		return
	}
	// Call the service to update the seat reservation
	result, err := service.SeatReservationService().UpdateSeatReservation(ctx, model.SeatReservationUpdateInput{
		ID:            int64(seatReservationIDRequest.SeatReservationID),
		SeatID:        int64(seatReservation.SeatID),
		TrainID:       int64(seatReservation.TrainID),
		OrderNumber:   seatReservation.OrderNumber,
		FromStationID: int64(seatReservation.FromStationID),
		ToStationID:   int64(seatReservation.ToStationID),
	})
	if err != nil {
		response.ErrorResponse(ctx, response.UpdateErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.UpdateSuccessCodeStatus, result)
}

// @Summary Cancel a seat reservation
// @Description Cancel a seat reservation
// @Tags SeatReservation
// @Accept json
// @Produce json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param seat_reservation_id path int true "Seat Reservation ID"
// @Success 200 {object} response.ResponseData{data=model.SeatReservationOutput}
// @Failure 400 {object} response.ErrorResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /seat-reservation/cancel/{seat_reservation_id} [delete]
func (c *cSeatReservationController) CancelSeatReservation(ctx *gin.Context) {
	// Get the seat reservation ID from the URL parameter
	var seatReservationIDRequest vo.SeatReservationIDRequest
	if err := ctx.ShouldBindUri(&seatReservationIDRequest); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	// Call the service to cancel the seat reservation
	result, err := service.SeatReservationService().DeleteSeatReservation(ctx, int64(seatReservationIDRequest.SeatReservationID))
	if err != nil {
		response.ErrorResponse(ctx, response.NotFoundErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.DeleteSuccessCodeStatus, result)
}

// @Summary Get seat reservation by seat reservation id
// @Description Get seat reservation by seat reservation id
// @Tags SeatReservation
// @Accept json
// @Produce json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param seat_reservation_id path int true "Seat Reservation ID"
// @Success 200 {object} response.ResponseData{data=model.SeatReservationOutput}
// @Failure 400 {object} response.ErrorResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /seat-reservation/get-one/{seat_reservation_id} [get]
func (c *cSeatReservationController) GetSeatReservationById(ctx *gin.Context) {
	// Get the seat reservation ID from the URL parameter
	var seatReservationIDRequest vo.SeatReservationIDRequest
	if err := ctx.ShouldBindUri(&seatReservationIDRequest); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	// Call the service to get the seat reservation by ID
	result, err := service.SeatReservationService().GetSeatReservationById(ctx, int64(seatReservationIDRequest.SeatReservationID))
	if err != nil {
		response.ErrorResponse(ctx, response.NotFoundErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, result)
}

// @Summary Get all seat reservations by order number
// @Description Get all seat reservations
// @Tags SeatReservation
// @Accept json
// @Produce json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param order_number query string true "Order Number"
// @Param from_station_id query int false "From Station ID"
// @Param to_station_id query int false "To Station ID"
// @Param train_id query int false "Train ID"
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Success      200  {object}  response.ResponseData{data=[]model.SeatReservationOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /seat-reservation/get-all-by-order-number [get]
func (c *cSeatReservationController) GetAllSeatReservationsByOrderNumber(ctx *gin.Context) {
	var seatReservationListRequest vo.SeatReservationListRequest
	if err := ctx.ShouldBindQuery(&seatReservationListRequest); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	// Call the service to get all seat reservations
	result, err := service.SeatReservationService().GetSeatReservationByOrderNumber(ctx, model.SeatReservationListInput{
		OrderNumber: seatReservationListRequest.OrderNumber,
		Page:        int64(seatReservationListRequest.Page),
		Limit:       int64(seatReservationListRequest.Limit),
	})
	if err != nil {
		response.ErrorResponse(ctx, response.NotFoundErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, result)
}

// @Summary Get all seat reservations by train id
// @Description Get all seat reservations by train id
// @Tags SeatReservation
// @Accept json
// @Produce json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param order_number query string false "Order Number"
// @Param from_station_id query int false "From Station ID"
// @Param to_station_id query int false "To Station ID"
// @Param train_id query int true "Train ID"
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Success      200  {object}  response.ResponseData{data=[]model.SeatReservationOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /seat-reservation/get-all-by-train-id [get]
func (c *cSeatReservationController) GetAllSeatReservationsByTrainId(ctx *gin.Context) {
	var seatReservationListRequest vo.SeatReservationListRequest
	if err := ctx.ShouldBindQuery(&seatReservationListRequest); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	// Call the service to get all seat reservations by train id
	result, err := service.SeatReservationService().GetSeatReservationByTrainId(ctx, model.SeatReservationListInput{
		TrainID: int64(seatReservationListRequest.TrainID),
		Page:    int64(seatReservationListRequest.Page),
		Limit:   int64(seatReservationListRequest.Limit),
	})
	if err != nil {
		response.ErrorResponse(ctx, response.NotFoundErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, result)
}
