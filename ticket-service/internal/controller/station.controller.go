package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/service"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/vo"
	"github.com/quangdat385/holiday-ticket/ticket-service/response"
)

var StationController = new(cStationController)

type cStationController struct {
}

// @Summary      Create station
// @Description  Create station
// @Tags         station
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        payload  body      vo.CreateStationRequest  true  "Create station request"
// @Success      200  {object}  response.ResponseData{data=model.StationOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /station/create [post]
func (c *cStationController) CreateStation(ctx *gin.Context) {
	// Parse request body
	var req vo.CreateStationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	// Call service to create station
	station, err := service.StationService().CreateStation(ctx, model.StationInput{
		Name:   req.Name,
		Code:   req.Code,
		Status: int32(req.Status),
	})
	if err != nil {
		response.ErrorResponse(ctx, response.CreateErrorCodeStatus, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.SuccessCodeStatus, station)
}

// @Summary      Update station
// @Description  Update station
// @Tags         station
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// Param        payload  body      vo.UpdateStationRequest  true  "Update station request"
// @Success      200  {object}  response.ResponseData{data=model.StationOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /station/update [patch]
func (c *cStationController) UpdateStation(ctx *gin.Context) {
	// Parse request body
	var req vo.UpdateStationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	// Call service to update station
	station, err := service.StationService().UpdateStation(ctx, int64(req.StationID), model.StationInput{
		Name:   req.Name,
		Code:   req.Code,
		Status: int32(req.Status),
	})
	if err != nil {
		response.ErrorResponse(ctx, response.UpdateErrorCodeStatus, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.SuccessCodeStatus, station)
}

// @Summary      Update station status
// @Description  Update station status
// @Tags         station
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        station_id  path      int  true  "Station ID"
// @Param        payload  body      vo.UpdateStationStatusRequest  true  "Update station status request"
// @Success      200  {object}  response.ResponseData{data=model.StationOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /station/update-status/{station_id} [patch]
func (c *cStationController) UpdateStationStatus(ctx *gin.Context) {
	var stationIDRequest vo.StationIDRequest
	// Parse request body
	if err := ctx.ShouldBindUri(&stationIDRequest); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	var req vo.UpdateStationStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	// Call service to update station status
	station, err := service.StationService().UpdateStationStatus(ctx, int64(stationIDRequest.StationID), int32(req.Status))
	if err != nil {
		response.ErrorResponse(ctx, response.UpdateErrorCodeStatus, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.SuccessCodeStatus, station)
}

// @Summary      Delete station
// @Description  Delete station
// @Tags         station
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        id  path      int  true  "Station ID"
// @Success      200  {object}  response.ResponseData{data=bool}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /station/delete/{id} [delete]
func (c *cStationController) DeleteStation(ctx *gin.Context) {
	var stationIDRequest vo.StationIDRequest
	// Parse request body
	if err := ctx.ShouldBindUri(&stationIDRequest); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	// Call service to delete station
	out, err := service.StationService().DeleteStation(ctx, int64(stationIDRequest.StationID))
	if err != nil {
		response.ErrorResponse(ctx, response.DeleteErrorCodeStatus, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.SuccessCodeStatus, out)
}

// @Summary      Get station by id
// @Description  Get station by id
// @Tags         station
// @Accept       json
// @Produce      json
// @Param        id  path      int  true  "Station ID"
// @Success      200  {object}  response.ResponseData{data=model.StationOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /station/get-by-id/{id} [get]
func (c *cStationController) GetStationByID(ctx *gin.Context) {
	var stationIDRequest vo.StationIDRequest
	// Parse request body
	if err := ctx.ShouldBindUri(&stationIDRequest); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	// Call service to get station by id
	station, err := service.StationService().GetStationByID(ctx, int64(stationIDRequest.StationID))
	if err != nil {
		response.ErrorResponse(ctx, response.NotFoundErrorCodeStatus, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.SuccessCodeStatus, station)
}

// @Summary      Get all station
// @Description  Get all station
// @Tags         station
// @Accept       json
// @Produce      json
// @Param        status  query      int  true  "Status"
// @Param        limit   query      int  true  "Limit"
// @Param        page    query      int  true  "Page"
// @Success      200  {object}  response.ResponseData{data=[]model.StationOutput}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /station/get-all [get]
func (c *cStationController) GetAllStation(ctx *gin.Context) {
	var stationStatusRequest vo.StationListRequest
	// Parse request body
	if err := ctx.ShouldBindQuery(&stationStatusRequest); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	// Call service to get all station
	stations, err := service.StationService().GetAllStation(ctx, model.StationListInput{
		Page:   stationStatusRequest.Page,
		Limit:  stationStatusRequest.Limit,
		Status: int32(stationStatusRequest.Status),
	})
	if err != nil {
		response.ErrorResponse(ctx, response.NotFoundErrorCodeStatus, err.Error())
		return
	}

	response.SuccessResponse(ctx, response.SuccessCodeStatus, stations)
}
