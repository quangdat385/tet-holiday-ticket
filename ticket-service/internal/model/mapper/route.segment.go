package mapper

import (
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/database"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
)

func ToRouteSegmentDTO(routeSegment database.GetRouteSegmentByIdRow) model.RouteSegmentOutPut {
	return model.RouteSegmentOutPut{
		ID:            routeSegment.ID,
		TrainID:       routeSegment.TrainID,
		FromStationID: routeSegment.FromStationID,
		ToStationID:   routeSegment.ToStationID,
		SegmentOrder:  routeSegment.SegmentOrder,
		DistanceKm:    routeSegment.DistanceKm,
	}
}
