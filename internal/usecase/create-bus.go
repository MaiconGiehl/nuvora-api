package usecase

import (
	"github.com/maicongiehl/nuvera-api/internal/dto"
	"github.com/maicongiehl/nuvera-api/internal/entity"
	"github.com/maicongiehl/nuvera-api/internal/infra/database"
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
		CompanyID: input.CompanyID,
	}

	err := c.BusRepository.Save(&entity)
	if err != nil {
		return err
	}

	return nil
}
