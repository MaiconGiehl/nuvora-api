package usecase

import (
	"github.com/maicongiehl/nuvera-api/internal/dto"
	"github.com/maicongiehl/nuvera-api/internal/entity"
	"github.com/maicongiehl/nuvera-api/internal/infra/database"
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
	entity := entity.Bus{
		Id:							id,
		Number: 				input.Number,
		MaxPassengers: 	input.MaxPassengers,
	}

	err := c.BusRepository.Update(&entity)
	if err != nil {
		return err
	}

	return nil
}
