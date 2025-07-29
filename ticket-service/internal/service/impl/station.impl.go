package impl

import (
	"context"
	"fmt"

	"github.com/quangdat385/holiday-ticket/ticket-service/internal/database"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model/mapper"
)

type sStation struct {
	r *database.Queries
}

func NewStationImpl(r *database.Queries) *sStation {
	return &sStation{
		r: r,
	}
}
func (s *sStation) GetStationByID(context context.Context, id int64) (out model.StationOutput, err error) {
	station, err := s.r.GetStationById(context, id)
	if err != nil {
		return out, err
	}
	out = mapper.ToStationDTO(station)
	return out, nil
}
func (s *sStation) GetAllStation(context context.Context, in model.StationListInput) (out []model.StationOutput, err error) {
	offset := (in.Page - 1) * in.Limit
	stations, err := s.r.GetStationList(context, database.GetStationListParams{
		Offset: int32(offset),
		Limit:  int32(in.Limit),
		Status: in.Status,
	})
	fmt.Println("GetAllStation", in.Page, in.Limit, in.Status, "offset", offset, "limit", in.Limit)
	if err != nil {
		return out, err
	}
	if len(stations) == 0 {
		return out, nil
	}
	out = mapper.ToStationDTOList(stations)
	return out, nil
}
func (s *sStation) CreateStation(context context.Context, station model.StationInput) (out model.StationOutput, err error) {
	result, err := s.r.InsertStation(context, database.InsertStationParams{
		Name: station.Name,
		Code: station.Code,
	})
	if err != nil {
		return out, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return out, err
	}
	stationResult, err := s.r.GetStationById(context, id)
	if err != nil {
		return out, err
	}
	out = mapper.ToStationDTO(stationResult)
	return out, nil
}
func (s *sStation) UpdateStation(context context.Context, id int64, station model.StationInput) (out model.StationOutput, err error) {
	stationDB, err := s.r.GetStationById(context, id)
	if err != nil {
		return out, err
	}
	result, err := s.r.UpdateStation(context, database.UpdateStationParams{
		ID:     id,
		Name:   station.Name,
		Code:   station.Code,
		Status: stationDB.Status,
	})
	if err != nil {
		return out, err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return out, err
	}
	stationDB.Name = station.Name
	stationDB.Code = station.Code
	out = mapper.ToStationDTO(stationDB)
	return out, nil
}
func (s *sStation) UpdateStationStatus(context context.Context, id int64, status int32) (out model.StationOutput, err error) {
	stationDB, err := s.r.GetStationById(context, id)
	if err != nil {
		return out, err
	}
	result, err := s.r.UpdateStation(context, database.UpdateStationParams{
		ID:     id,
		Name:   stationDB.Name,
		Code:   stationDB.Code,
		Status: status,
	})
	if err != nil {
		return out, err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return out, err
	}
	stationDB.Status = status
	out = mapper.ToStationDTO(stationDB)
	return out, nil
}
func (s *sStation) DeleteStation(context context.Context, id int64) (out bool, err error) {
	result, err := s.r.DeleteStation(context, id)
	if err != nil {
		return false, err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return false, err
	}
	return true, nil
}
