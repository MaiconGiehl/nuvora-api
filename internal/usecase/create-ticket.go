package usecase

import (
	"github.com/MaiconGiehl/API/internal/dto"
	"github.com/MaiconGiehl/API/internal/entity"
	"github.com/MaiconGiehl/API/internal/infra/database"
)


type CreateTicketUseCase struct {
	TicketRepository database.TicketRepository
}

func NewCreateTicketUseCase(
	TicketRepository database.TicketRepository,
) *CreateTicketUseCase {
	return &CreateTicketUseCase{
		TicketRepository: TicketRepository,
	}
}

func (c *CreateTicketUseCase) Execute(input *dto.TicketInputDTO) error {
	entity := entity.Ticket{
		ClienteID: 			input.ClienteID,
		Price: 					input.Price,
		Status: 				input.Status,
		DepartureCity: 	input.DepartureCity,
		DepartureTime: 	input.DepartureTime,
		DestinyCity: 		input.DestinyCity,
		DestinyTime: 		input.DestinyTime,
		BusID: 					input.BusID,
	}

	err := c.TicketRepository.Save(&entity)
	if err != nil {
		return err
	}

	return nil
}
