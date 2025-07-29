package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/service"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/vo"
	"github.com/quangdat385/holiday-ticket/ticket-service/response"
)

var RouteSegmenController = new(cRouteSegmentController)

type cRouteSegmentController struct {
}

// @Summary      Create route segment
// @Description  Create route segment
// @Tags         route-segment
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        payload  body      vo.CreateSegmentRequest  true  "Create route segment request"
// @Success      200  {object}  response.ResponseData{data=model.RouteSegmentOutPut}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /route-segment/create [post]
func (c *cRouteSegmentController) CreateRouteSegment(ctx *gin.Context) {
	var params vo.CreateSegmentRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	routeSegment, err := service.RouteSegmentService().CreateRouteSegment(ctx, model.RouteSegmentCreateInPut{
		TrainID:       int64(params.TrainID),
		FromStationID: int64(params.FromStationID),
		ToStationID:   int64(params.ToStationID),
		SegmentOrder:  int32(params.SegmentOrder),
		DistanceKm:    int32(params.DistanceKm),
	})
	if err != nil {
		response.ErrorResponse(ctx, response.CreateErrorCodeStatus, err.Error())
		return
	}
	if routeSegment.ID == 0 { // Assuming IsEmpty() is a method to check if the struct is empty
		response.ErrorResponse(ctx, response.CreateErrorCodeStatus, "create route segment failed")
		return
	}
	response.SuccessResponse(ctx, response.CreateSuccessCodeStatus, routeSegment)
}

// @Summary      Get route segment
// @Description  Get route segment
// @Tags         route-segment
// @Accept       json
// @Produce      json
// @Param        segment_id   path      int  true  "Route Segment ID"
// @Success      200  {object}  response.ResponseData{data=model.RouteSegmentOutPut}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /route-segment/get-one/{segment_id} [get]
func (c *cRouteSegmentController) GetRouteSegment(ctx *gin.Context) {
	var params vo.SegmentIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	routeSegment, err := service.RouteSegmentService().GetRouteSegmentsByRouteID(ctx, int64(params.SegmentID))
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	if routeSegment.ID == 0 { // Assuming IsEmpty() is a method to check if the struct is empty
		response.ErrorResponse(ctx, response.NotFoundErrorCodeStatus, "route segment not found")
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, routeSegment)
}

// @Summary      Get route segments by train id
// @Description  Get route segments by train id
// @Tags         route-segment
// @Accept       json
// @Produce      json
// @Param      train_id   path      int  true  "Train ID"
// @Success      200  {object}  response.ResponseData{data=[]model.RouteSegmentOutPut}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /route-segment/get-by-train-id/{train_id} [get]
func (c *cRouteSegmentController) GetRouteSegmentsByTrainID(ctx *gin.Context) {
	var params vo.TrainIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	routeSegments, err := service.RouteSegmentService().GetRouteSegmentsByTrainID(ctx, int64(params.TrainID))
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, routeSegments)
}

// @Summary      Get route segments by from station id
// @Description  Get route segments by from station id
// @Tags         route-segment
// @Accept       json
// @Produce      json
// @Param        from_station_id   path      int  true  "From Station ID"
// @Success      200  {object}  response.ResponseData{data=[]model.RouteSegmentOutPut}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /route-segment/get-by-from-station-id/{from_station_id} [get]
func (c *cRouteSegmentController) GetRouteSegmentsByFromStationID(ctx *gin.Context) {
	var params vo.SegmentFromStationIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	routeSegments, err := service.RouteSegmentService().GetRouteSegmentsByFromStationID(ctx, int64(params.FromStationID))
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, routeSegments)
}

// @Summary      Get route segments by to station id
// @Description  Get route segments by to station id
// @Tags         route-segment
// @Accept       json
// @Produce      json
// @Param        to_station_id   path      int  true  "To Station ID"
// @Success      200  {object}  response.ResponseData{data=[]model.RouteSegmentOutPut}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /route-segment/get-by-to-station-id/{to_station_id} [get]
func (c *cRouteSegmentController) GetRouteSegmentsByToStationID(ctx *gin.Context) {
	var params vo.SegmentToStationIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	routeSegments, err := service.RouteSegmentService().GetRouteSegmentsByToStationID(ctx, int64(params.ToStationID))
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, routeSegments)
}

// @Summary      Update route segment
// @Description  Update route segment
// @Tags         route-segment
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        segment_id   path      int  true  "Route Segment ID"
// @Param        payload  body      vo.UpdateSegmentRequest  true  "Update route segment request"
// @Success      200  {object}  response.ResponseData{data=model.RouteSegmentOutPut}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /route-segment/update/{segment_id} [put]
func (c *cRouteSegmentController) UpdateRouteSegment(ctx *gin.Context) {
	var params vo.UpdateSegmentRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	var segmentIDParams vo.SegmentIDRequest
	if err := ctx.ShouldBindUri(&segmentIDParams); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}
	if segmentIDParams.SegmentID != params.SegmentID {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, "segment id not match")
		return
	}

	routeSegment, err := service.RouteSegmentService().UpdateRouteSegment(ctx, model.RouteSegmentUpdateInPut{
		ID:            int64(segmentIDParams.SegmentID),
		TrainID:       int64(params.TrainID),
		FromStationID: int64(params.FromStationID),
		ToStationID:   int64(params.ToStationID),
		SegmentOrder:  int32(params.SegmentOrder),
		DistanceKm:    int32(params.DistanceKm),
	})
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.SuccessCodeStatus, routeSegment)
}

// @Summary      Delete route segment
// @Description  Delete route segment
// @Tags         route-segment
// @Accept       json
// @Produce      json
// @Param        x-client-id header string true "Client ID"
// @Param        x-device-id header string true "Device ID"
// @Param        segment_id   path      int  true  "Route Segment ID"
// @Success      200  {object}  response.ResponseData{data=bool}
// @Failure      400  {object}  response.ErrorResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /route-segment/delete/{segment_id} [delete]
func (c *cRouteSegmentController) DeleteRouteSegment(ctx *gin.Context) {
	var params vo.SegmentIDRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.ErrorResponse(ctx, response.ParamInvalidCodeStatus, err.Error())
		return
	}

	result, err := service.RouteSegmentService().DeleteRouteSegment(ctx, int64(params.SegmentID))
	if err != nil {
		response.ErrorResponse(ctx, response.ErrorCodeStatus, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.DeleteSuccessCodeStatus, result)
}
