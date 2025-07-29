package mapper

import (
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/database"
	"github.com/quangdat385/holiday-ticket/ticket-service/internal/model"
)

func ToSeatDTO(seat database.GetSeatByIdRow) model.SeatOutput {
	return model.SeatOutput{
		ID:         seat.ID,
		TrainID:    int64(seat.TrainID),
		SeatNumber: seat.SeatNumber,
		SeatClass:  string(seat.SeatClass),
		Status:     seat.Status,
	}
}
func ToSeatsDTO(seats []database.GetSeatsByTrainIdRow) []model.SeatOutput {
	var seatDTOs []model.SeatOutput
	for _, seat := range seats {
		if seat.ID != 0 {
			seatMapper := database.GetSeatByIdRow(seat)
			seatDTOs = append(seatDTOs, ToSeatDTO(seatMapper))
		}
	}
	return seatDTOs
}
