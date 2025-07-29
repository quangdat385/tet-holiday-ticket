package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/service"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/vo"
	"github.com/quangdat385/holiday-ticket/ticket-service/response"
)

var SeatController = new(cSeatController)

type cSeatController struct {
}

// @Summary      Create seat
// @Description  Create seat
// @Tags         seat
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        payload  body      vo.CreateSeatRequest  true  "Create seat request"
// @Success      200  {object}  response.ResponseData{data=model.SeatOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /seat/create [post]
func (c *cSeatController) CreateSeat(ctx *gin.Context) {
	var params vo.CreateSeatRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	seat, err := service.SeatService().CreateSeat(ctx, model.CreateSeatInput{
		TrainID:    int64(params.TrainID),
		SeatNumber: params.SeatNumber,
		SeatClass:  params.SeatClass,
		Status:     int32(params.Status),
	})
	if err != nil {
		response.ErrorResponse(ctx, response.CreateErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.CreateSuccessCodeStatus, seat)
}

// @Summary      Get seat by id
// @Description  Get seat by id
// @Tags         seat
// @Accept       json
// @Produce      json
// @Param        seat_id   path      int  true  "Seat ID"
// @Success      200  {object}  response.ResponseData{data=model.SeatOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /seat/get-one/{seat_id} [get]
func (c *cSeatController) GetSeat(ctx *gin.Context) {
	var params vo.SeatIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	seat, err := service.SeatService().GetSeatByID(ctx, int64(params.SeatID))
	if err != nil {
		response.ErrorResponse(ctx, response.NotFoundErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, seat)
}

// @Summary      Get seats by train
// @Description  Get seats by train
// @Tags         seat
// @Accept       json
// @Produce      json
// @Param        train_id   path      int  true  "Train ID"
// @Param        limit      query     int  true  "Limit"
// @Param        page       query     int  true  "Page"
// @Success      200  {object}  response.ResponseData{data=[]model.SeatOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /seat/get-by-train/{train_id} [get]
func (c *cSeatController) GetSeatsByTrain(ctx *gin.Context) {
	var train vo.TrainIDRequest
	if err := ctx.ShouldBindUri(&train); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	var params vo.SeatListByTrainIDRequest
	if err := ctx.ShouldBindQuery(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	if train.TrainID != params.TrainID {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, "Train ID in path and query must match")
		return
	}
	seats, err := service.SeatService().GetSeatsByTrainID(ctx, model.SeatListInput{
		TrainID: int64(params.TrainID),
		Page:    int64(params.Page),
		Limit:   int64(params.Limit),
	})
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, seats)
}

// @Summary      Update seat
// @Description  Update seat
// @Tags         seat
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        seat_id   path      int  true  "Seat ID"
// @Param        payload  body      vo.UpdateSeatRequest  true  "Update seat request"
// @Success      200  {object}  response.ResponseData{data=model.SeatOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /seat/update/{seat_id} [put]
func (c *cSeatController) UpdateSeat(ctx *gin.Context) {
	var SeatID vo.SeatIDRequest
	if err := ctx.ShouldBindUri(&SeatID); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	var params vo.UpdateSeatRequest

	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	if SeatID.SeatID != params.SeatID {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, "Seat ID in path and body must match")
		return
	}
	seat, err := service.SeatService().UpdateSeat(ctx, model.UpdateSeatInput{
		ID:         int64(params.SeatID),
		SeatNumber: params.SeatNumber,
		SeatClass:  params.SeatClass,
		Status:     int32(params.Status),
	})
	if err != nil {
		response.ErrorResponse(ctx, response.UpdateErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.UpdateSuccessCodeStatus, seat)
}

// @Summary      Delete seat
// @Description  Delete seat
// @Tags         seat
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        seat_id   path      int  true  "Seat ID"
// @Success      200  {object}  response.ResponseData{data=model.SeatOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /seat/delete/{seat_id} [delete]
func (c *cSeatController) DeleteSeat(ctx *gin.Context) {
	var params vo.SeatIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	seat, err := service.SeatService().DeleteSeat(ctx, int64(params.SeatID))
	if err != nil {
		response.ErrorResponse(ctx, response.NotFoundErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, seat)
}
