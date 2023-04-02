package usecase

import (
	"github.com/maicongiehl/nuvera-api/internal/entity"
	"github.com/maicongiehl/nuvera-api/internal/infra/database"
)


type DeleteBusUseCase struct {
	BusRepository database.BusRepository
}

func NewDeleteBusUseCase(
	BusRepository database.BusRepository,
) *DeleteBusUseCase {
	return &DeleteBusUseCase{
		BusRepository: BusRepository,
	}
}

func (c *DeleteBusUseCase) Execute(id int) error {
	entity := entity.Bus{
		Id:								id,
	}

	err := c.BusRepository.Delete(&entity)
	if err != nil {
		return err
	}

	return nil
}
