package usecase

import (
	"github.com/MaiconGiehl/API/internal/dto"
	"github.com/MaiconGiehl/API/internal/infra/database"
)


type GetAllTicketUseCase struct {
	TicketRepository database.TicketRepository
}

func NewGetAllTicketUseCase(
	TicketRepository database.TicketRepository,
) *GetAllTicketUseCase {
	return &GetAllTicketUseCase{
		TicketRepository: TicketRepository,
	}
}

func (c *GetAllTicketUseCase) Execute() (*[]dto.TicketOutputDTO, error) {
	output, err := c.TicketRepository.GetAll()
	if err != nil {
		return output, err
	}

	return output, nil
}
