package mapper

import (
	"github.com/quangdat385/holiday-ticket/communications-service/internal/database"
	"github.com/quangdat385/holiday-ticket/communications-service/internal/model"
)

func ToInformationDTO(information database.PreGoCommunicationInfo99999) model.InformationOutput {
	return model.InformationOutput{
		ID:        information.ID,
		UserID:    information.UserID,
		Status:    information.Status.Bool,
		CreatedAt: information.CreatedAt.Time,
		UpdatedAt: information.UpdatedAt.Time,
	}
}
