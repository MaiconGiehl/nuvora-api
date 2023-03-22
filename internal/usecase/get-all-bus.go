package usecase

import (
	"github.com/MaiconGiehl/API/internal/dto"
	"github.com/MaiconGiehl/API/internal/infra/database"
)


type GetAllBusUseCase struct {
	BusRepository database.BusRepository
}

func NewGetAllBusUseCase(
	BusRepository database.BusRepository,
) *GetAllBusUseCase {
	return &GetAllBusUseCase{
		BusRepository: BusRepository,
	}
}

func (c *GetAllBusUseCase) Execute() (*[]dto.BusOutputDTO, error) {
	output, err := c.BusRepository.GetAll()
	if err != nil {
		return output, err
	}

	return output, nil
}
