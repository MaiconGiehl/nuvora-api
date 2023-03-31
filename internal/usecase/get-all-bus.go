package usecase

import (
	"github.com/maicongiehl/techtur-api/internal/dto"
	"github.com/maicongiehl/techtur-api/internal/infra/database"
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
