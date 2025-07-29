package mapper

import (
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/database"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
)

func ToStationDTO(station database.GetStationByIdRow) model.StationOutput {
	return model.StationOutput{
		ID:     int64(station.ID),
		Name:   station.Name,
		Code:   station.Code,
		Status: int32(station.Status),
	}
}
func ToStationDTOList(stations []database.GetStationListRow) []model.StationOutput {
	out := make([]model.StationOutput, 0)
	for _, station := range stations {
		if station.ID != 0 {
			stationOutput := model.StationOutput{
				ID:     int64(station.ID),
				Name:   station.Name,
				Code:   station.Code,
				Status: int32(station.Status),
			}
			out = append(out, stationOutput)
		}
	}
	return out
}
