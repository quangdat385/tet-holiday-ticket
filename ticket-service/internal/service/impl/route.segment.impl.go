package impl

import (
	"context"

	"github.com/quangdat385/holiday-ticket/ticket-service/internal/database"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model/mapper"
	"github.com/quangdat385/holiday-ticket/ticket-service/response"
)

type sRouteSegment struct {
	r *database.Queries
}

func NewRouteSegmentImpl(r *database.Queries) *sRouteSegment {
	return &sRouteSegment{
		r: r,
	}
}

func (s *sRouteSegment) GetRouteSegmentsByRouteID(context context.Context, routeID int64) (out model.RouteSegmentOutPut, err error) {
	routeSegment, err := s.r.GetRouteSegmentById(context, routeID)
	if err != nil {
		return out, err
	}
	if routeSegment == (database.GetRouteSegmentByIdRow{}) {
		return out, response.ErrRouteSegmentNotFoundErr
	}
	out = mapper.ToRouteSegmentDTO(routeSegment)
	return out, nil
}
func (s *sRouteSegment) GetRouteSegmentsByTrainID(context context.Context, trainID int64) (out []model.RouteSegmentOutPut, err error) {
	routeSegments, err := s.r.GetRouteSegmentsByTrainId(context, trainID)
	if err != nil {
		return out, err
	}
	if len(routeSegments) == 0 {
		return out, response.ErrRouteSegmentNotFoundErr
	}
	for _, routeSegment := range routeSegments {
		if routeSegment.ID != 0 {
			routeSegmentMapper := database.GetRouteSegmentByIdRow(routeSegment)
			out = append(out, mapper.ToRouteSegmentDTO(routeSegmentMapper))
		}
	}
	return out, nil
}
func (s *sRouteSegment) GetRouteSegmentsByFromStationID(context context.Context, fromStationID int64) (out []model.RouteSegmentOutPut, err error) {
	routeSegments, err := s.r.GetRouteSegmentsByFromStationId(context, fromStationID)
	if err != nil {
		return out, err
	}
	if len(routeSegments) == 0 {
		return out, response.ErrRouteSegmentNotFoundErr
	}
	for _, routeSegment := range routeSegments {
		if routeSegment.ID != 0 {
			routeSegmentMapper := database.GetRouteSegmentByIdRow(routeSegment)
			out = append(out, mapper.ToRouteSegmentDTO(routeSegmentMapper))
		}
	}
	return out, nil
}
func (s *sRouteSegment) GetRouteSegmentsByToStationID(context context.Context, toStationID int64) (out []model.RouteSegmentOutPut, err error) {
	routeSegments, err := s.r.GetRouteSegmentsByToStationId(context, toStationID)
	if err != nil {
		return out, err
	}
	if len(routeSegments) == 0 {
		return out, response.ErrRouteSegmentNotFoundErr
	}
	for _, routeSegment := range routeSegments {
		if routeSegment.ID != 0 {
			routeSegmentMapper := database.GetRouteSegmentByIdRow(routeSegment)
			out = append(out, mapper.ToRouteSegmentDTO(routeSegmentMapper))
		}
	}
	return out, nil
}
func (s *sRouteSegment) CreateRouteSegment(context context.Context, in model.RouteSegmentCreateInPut) (out model.RouteSegmentOutPut, err error) {
	result, err := s.r.InsertRouteSegment(context, database.InsertRouteSegmentParams{
		TrainID:       in.TrainID,
		FromStationID: in.FromStationID,
		ToStationID:   in.ToStationID,
		SegmentOrder:  in.SegmentOrder,
		DistanceKm:    in.DistanceKm,
	})
	if err != nil {
		return out, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return out, err
	}
	routeSegment, err := s.r.GetRouteSegmentById(context, id)
	if err != nil {
		return out, err
	}
	if routeSegment == (database.GetRouteSegmentByIdRow{}) {
		return out, response.ErrCreateRouteSegmentErr
	}
	out = mapper.ToRouteSegmentDTO(routeSegment)
	return out, nil
}
func (s *sRouteSegment) UpdateRouteSegment(context context.Context, in model.RouteSegmentUpdateInPut) (out model.RouteSegmentOutPut, err error) {
	routeSegment, err := s.r.GetRouteSegmentById(context, in.ID)
	if err != nil {
		return out, err
	}
	if routeSegment == (database.GetRouteSegmentByIdRow{}) {
		return out, response.ErrRouteSegmentNotFoundErr
	}
	if in.TrainID != 0 {
		routeSegment.TrainID = in.TrainID
	}
	if in.FromStationID != 0 {
		routeSegment.FromStationID = in.FromStationID
	}
	if in.ToStationID != 0 {
		routeSegment.ToStationID = in.ToStationID
	}
	if in.SegmentOrder != 0 {
		routeSegment.SegmentOrder = in.SegmentOrder
	}
	if in.DistanceKm != 0 {
		routeSegment.DistanceKm = in.DistanceKm
	}
	result, err := s.r.UpdateRouteSegment(context, database.UpdateRouteSegmentParams{
		ID:            in.ID,
		TrainID:       routeSegment.TrainID,
		FromStationID: routeSegment.FromStationID,
		ToStationID:   routeSegment.ToStationID,
		SegmentOrder:  routeSegment.SegmentOrder,
		DistanceKm:    routeSegment.DistanceKm,
	})
	if err != nil {
		return out, err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return out, err
	}
	out = mapper.ToRouteSegmentDTO(routeSegment)
	return out, nil
}
func (s *sRouteSegment) DeleteRouteSegment(context context.Context, id int64) (out bool, err error) {
	routeSegment, err := s.r.GetRouteSegmentById(context, id)
	if err != nil {
		return out, err
	}
	if routeSegment == (database.GetRouteSegmentByIdRow{}) {
		return out, response.ErrRouteSegmentNotFoundErr
	}
	_, err = s.r.DeleteRouteSegment(context, id)
	if err != nil {
		return out, err
	}
	out = true
	return out, nil
}
