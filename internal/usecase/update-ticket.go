package usecase

import (
	"github.com/MaiconGiehl/API/internal/dto"
	"github.com/MaiconGiehl/API/internal/entity"
	"github.com/MaiconGiehl/API/internal/infra/database"
)


type UpdateTicketUseCase struct {
	TicketRepository database.TicketRepository
}

func NewUpdateTicketUseCase(
	TicketRepository database.TicketRepository,
) *UpdateTicketUseCase {
	return &UpdateTicketUseCase{
		TicketRepository: TicketRepository,
	}
}

func (c *UpdateTicketUseCase) Execute(id int, input *dto.TicketInputDTO) error {
	entity := entity.Ticket{
		ID:							id,
		ClienteID: 			input.ClienteID,
		Price: 					input.Price,
		Status: 				input.Status,
		DepartureCity: 	input.DepartureCity,
		DepartureTime: 	input.DepartureTime,
		DestinyCity: 		input.DestinyCity,
		DestinyTime: 		input.DestinyTime,
		BusID: 					input.BusID,
	}

	err := c.TicketRepository.Update(&entity)
	if err != nil {
		return err
	}

	return nil
}
