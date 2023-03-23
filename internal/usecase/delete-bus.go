package usecase

import (
	"github.com/MaiconGiehl/API/internal/entity"
	"github.com/MaiconGiehl/API/internal/infra/database"
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
