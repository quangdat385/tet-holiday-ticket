package service

import (
	"context"

	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
)

type (
	IRouteSegmentService interface {
		GetRouteSegmentsByRouteID(context context.Context, routeID int64) (out model.RouteSegmentOutPut, err error)
		GetRouteSegmentsByTrainID(context context.Context, trainID int64) (out []model.RouteSegmentOutPut, err error)
		GetRouteSegmentsByFromStationID(context context.Context, fromStationID int64) (out []model.RouteSegmentOutPut, err error)
		GetRouteSegmentsByToStationID(context context.Context, toStationID int64) (out []model.RouteSegmentOutPut, err error)
		CreateRouteSegment(context context.Context, in model.RouteSegmentCreateInPut) (out model.RouteSegmentOutPut, err error)
		UpdateRouteSegment(context context.Context, in model.RouteSegmentUpdateInPut) (out model.RouteSegmentOutPut, err error)
		DeleteRouteSegment(context context.Context, id int64) (out bool, err error)
	}
)

var (
	localRouteSegmentService IRouteSegmentService
)

func RouteSegmentService() IRouteSegmentService {
	if localRouteSegmentService == nil {
		panic("implement localRouteSegmentService not found for interface IRouteSegmentService")
	}
	return localRouteSegmentService
}
func InitRouteSegmentService(i IRouteSegmentService) {
	localRouteSegmentService = i
}
