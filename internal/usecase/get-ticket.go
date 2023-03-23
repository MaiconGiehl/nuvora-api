package usecase

import (
	"github.com/MaiconGiehl/API/internal/dto"
	"github.com/MaiconGiehl/API/internal/entity"
	"github.com/MaiconGiehl/API/internal/infra/database"
)


type GetTicketUseCase struct {
	TicketRepository database.TicketRepository
}

func NewGetTicketUseCase(
	TicketRepository database.TicketRepository,
) *GetTicketUseCase {
	return &GetTicketUseCase{
		TicketRepository: TicketRepository,
	}
}

func (c *GetTicketUseCase) Execute(id int) (*dto.TicketOutputDTO, error) {
	entity := entity.Ticket{
		ID:								id,
	}

	output, err := c.TicketRepository.GetById(&entity)
	if err != nil {
		return output, err
	}

	return output, nil
}
