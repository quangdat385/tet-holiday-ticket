package mapper

import (
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/database"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
)

func ToSeatReservationDTO(seatReservation database.GetSeatReservationByIdRow) model.SeatReservationOutput {
	return model.SeatReservationOutput{
		ID:            seatReservation.ID,
		SeatID:        seatReservation.SeatID,
		TrainID:       seatReservation.TrainID,
		OrderNumber:   seatReservation.OrderNumber,
		FromStationID: seatReservation.FromStationID,
		ToStationID:   seatReservation.ToStationID,
	}
}
func ToSeatReservationListDTO(seatReservations []database.GetSeatReservationsByOrderNumberRow) []model.SeatReservationOutput {
	var seatReservationList []model.SeatReservationOutput
	for _, seatReservation := range seatReservations {
		if seatReservation.ID != 0 {
			seatReservationMapper := database.GetSeatReservationByIdRow(seatReservation)
			seatReservationList = append(seatReservationList, ToSeatReservationDTO(seatReservationMapper))
		}
	}
	return seatReservationList
}
