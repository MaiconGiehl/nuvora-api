package usecase

import (
	"github.com/MaiconGiehl/API/internal/dto"
	"github.com/MaiconGiehl/API/internal/entity"
	"github.com/MaiconGiehl/API/internal/infra/database"
)


type CreateBusUseCase struct {
	BusRepository database.BusRepository
}

func NewCreateBusUseCase(
	BusRepository database.BusRepository,
) *CreateBusUseCase {
	return &CreateBusUseCase{
		BusRepository: BusRepository,
	}
}

func (c *CreateBusUseCase) Execute(input *dto.BusInputDTO) error {
	entity := entity.Bus{
		Number: input.Number,
		MaxPassengers: input.MaxPassengers,
	}

	err := c.BusRepository.Save(&entity)
	if err != nil {
		return err
	}

	return nil
}
