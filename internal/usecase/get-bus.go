package usecase

import (
	"github.com/maicongiehl/techtur-api/internal/dto"
	"github.com/maicongiehl/techtur-api/internal/entity"
	"github.com/maicongiehl/techtur-api/internal/infra/database"
)


type GetBusUseCase struct {
	BusRepository database.BusRepository
}

func NewGetBusUseCase(
	BusRepository database.BusRepository,
) *GetBusUseCase {
	return &GetBusUseCase{
		BusRepository: BusRepository,
	}
}

func (c *GetBusUseCase) Execute(id int) (*dto.BusOutputDTO, error) {
	entity := entity.Bus{
		Id:								id,
	}

	output, err := c.BusRepository.GetById(&entity)
	if err != nil {
		return output, err
	}

	return output, nil
}
