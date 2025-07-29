package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/service"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/vo"
	"github.com/quangdat385/holiday-ticket/ticket-service/response"
)

var TrainController = new(cTrainController)

type cTrainController struct {
}

// @Summary      Create train
// @Description  Create train
// @Tags         train
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        payload  body      vo.CreateTrainRequest  true  "Create train request"
// @Success      200  {object}  response.ResponseData{data=model.TrainOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /train/create [post]
func (c *cTrainController) CreateTrain(ctx *gin.Context) {
	// Parse request body
	var req vo.CreateTrainRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	// Call service to create train
	train, err := service.TrainService().CreateTrain(ctx, model.CreateTrainInput{
		Name:               req.Name,
		Code:               req.Code,
		DepartureStationID: int64(req.DepartureStationID),
		ArrivalStationID:   int64(req.ArrivalStationID),
		DepartureTime:      req.DepartureTime,
		ArrivalTime:        req.ArrivalTime,
		Status:             int32(req.Status),
		Direction:          req.Direction,
		TrainType:          req.TrainType,
	})
	if err != nil {
		response.ErrorResponse(ctx, response.CreateErrorCodeStatus, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.SuccessCodeStatus, train)
}

// @Summary      Update train
// @Description  Update train
// @Tags         train
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        train_id path int true "train id"
// @Param        payload  body      vo.UpdateTrainRequest  true  "Update train request"
// @Success      200  {object}  response.ResponseData{data=model.TrainOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /train/update/{train_id} [patch]
func (c *cTrainController) UpdateTrain(ctx *gin.Context) {
	// Parse request body
	var req vo.UpdateTrainRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	// Parse train ID from URL
	var trainIDReq vo.TrainIDRequest
	if err := ctx.ShouldBindUri(&trainIDReq); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	if trainIDReq.TrainID != req.TrainID {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, "train_id in url and body must be the same")
		return
	}
	// Call service to update train
	train, err := service.TrainService().UpdateTrain(ctx, int64(trainIDReq.TrainID), model.TrainInput{
		Name:          req.Name,
		DepartureTime: req.DepartureTime,
		ArrivalTime:   req.ArrivalTime,
		Status:        int32(req.Status),
		Direction:     req.Direction,
	})
	if err != nil {
		response.ErrorResponse(ctx, response.UpdateErrorCodeStatus, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.SuccessCodeStatus, train)
}

// @Summary      Update train status
// @Description  Update train status
// @Tags         train
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        train_id path int true "train id"
// @Param        payload  body      vo.UpdateTrainStatusRequest  true  "Update train status request"
// @Success      200  {object}  response.ResponseData{data=model.TrainOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /train/update-status/{train_id}/ [patch]
func (c *cTrainController) UpdateTrainStatus(ctx *gin.Context) {
	// Parse train ID from URL
	var trainIDReq vo.TrainIDRequest
	if err := ctx.ShouldBindUri(&trainIDReq); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	// Parse status from URL
	var statusReq vo.UpdateTrainStatusRequest
	if err := ctx.ShouldBindJSON(&statusReq); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	if trainIDReq.TrainID != statusReq.TrainID {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, "train_id in url and body must be the same")
		return
	}
	// Call service to update train status
	train, err := service.TrainService().UpdateTrainStatus(ctx, int64(trainIDReq.TrainID), int32(statusReq.Status))
	if err != nil {
		response.ErrorResponse(ctx, response.UpdateErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, train)
}

// @Summary      Get train by id
// @Description  Get train by id
// @Tags         train
// @Accept       json
// @Produce      json
// @Param        x-device-id header string true "Device ID"
// @Param        train_id path int true "train id"
// @Success      200  {object}  response.ResponseData{data=model.TrainOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /train/get-by-id/{train_id} [get]
func (c *cTrainController) GetTrainByID(ctx *gin.Context) {
	// Parse train ID from URL
	var trainIDReq vo.TrainIDRequest
	if err := ctx.ShouldBindUri(&trainIDReq); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	// Call service to get train by ID
	train, err := service.TrainService().GetTrainByID(ctx, int64(trainIDReq.TrainID))
	if err != nil {
		response.ErrorResponse(ctx, response.NotFoundErrorCodeStatus, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.SuccessCodeStatus, train)
}

// @Summary      Delete train
// @Description  Delete train
// @Tags         train
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        train_id path int true "train id"
// @Success      200  {object}  response.ErrorResponseData{data=nil}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /train/delete/{train_id} [delete]
func (c *cTrainController) DeleteTrain(ctx *gin.Context) {
	// Parse train ID from URL
	var trainIDReq vo.TrainIDRequest
	if err := ctx.ShouldBindUri(&trainIDReq); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	// Call service to delete train
	err := service.TrainService().DeleteTrain(ctx, int64(trainIDReq.TrainID))
	if err != nil {
		response.ErrorResponse(ctx, response.DeleteErrorCodeStatus, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.SuccessCodeStatus, nil)
}

//
