package usecase

import (
	"github.com/maicongiehl/techtur-api/internal/entity"
	"github.com/maicongiehl/techtur-api/internal/infra/database"
)


type DeleteTicketUseCase struct {
	TicketRepository database.TicketRepository
}

func NewDeleteTicketUseCase(
	TicketRepository database.TicketRepository,
) *DeleteTicketUseCase {
	return &DeleteTicketUseCase{
		TicketRepository: TicketRepository,
	}
}

func (c *DeleteTicketUseCase) Execute(id int) error {
	entity := entity.Ticket{
		ID:								id,
	}

	err := c.TicketRepository.Delete(&entity)
	if err != nil {
		return err
	}

	return nil
}
