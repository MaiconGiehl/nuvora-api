package usecase

import (
	"github.com/maicongiehl/techtur-api/internal/dto"
	"github.com/maicongiehl/techtur-api/internal/entity"
	"github.com/maicongiehl/techtur-api/internal/enum"
	"github.com/maicongiehl/techtur-api/internal/infra/database"
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
		AccountID:						input.AccountID,
  	Status: 							enum.NotPaid,
  	TravelID: 						input.TravelID,
	}

	err := c.TicketRepository.Save(&entity)
	if err != nil {
		return err
	}

	return nil
}
