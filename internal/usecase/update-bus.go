package usecase

import (
	"github.com/MaiconGiehl/API/internal/dto"
	"github.com/MaiconGiehl/API/internal/entity"
	"github.com/MaiconGiehl/API/internal/infra/database"
)


type UpdateBusUseCase struct {
	BusRepository database.BusRepository
}

func NewUpdateBusUseCase(
	BusRepository database.BusRepository,
) *UpdateBusUseCase {
	return &UpdateBusUseCase{
		BusRepository: BusRepository,
	}
}

func (c *UpdateBusUseCase) Execute(id int, input *dto.BusInputDTO) error {
	bus := entity.Bus{
		Id:							id,
		Number: 				input.Number,
		MaxPassengers: 	input.MaxPassengers,
	}

	err := c.BusRepository.Update(&bus)
	if err != nil {
		return err
	}

	return nil
}
